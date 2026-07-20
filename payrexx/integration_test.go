package payrexx_test

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/zweiundeins/payrexx-go-sdk/payrexx"
)

/*
Live tests against a real Payrexx account.

These are what the offline suite structurally cannot be: proof that the spec
assembled by scripts/build_spec.py matches the API Payrexx actually serves. Every
correction that script applies is a claim about the live API, and the only way to
falsify one is to call it. Three of the corrections in build_spec.py exist
because these tests failed first.

They are skipped unless PAYREXX_LIVE_TEST is set, so `go test ./...` stays
offline and credential-free. CI runs them from a separate scheduled workflow, not
on pull requests -- a Payrexx outage must not block a merge.

	PAYREXX_LIVE_TEST=1 \
	PAYREXX_INSTANCE=example \
	PAYREXX_API_SECRET=... \
	go test -run TestLive -v ./...

What they do to the account: create gateways and delete them again. A gateway is
a hosted checkout page -- no money moves, nothing is charged, and each test
removes what it created. Nothing else is written.
*/

func liveClient(t *testing.T, auth payrexx.AuthScheme) *payrexx.APIClient {
	t.Helper()
	if os.Getenv("PAYREXX_LIVE_TEST") == "" {
		t.Skip("set PAYREXX_LIVE_TEST=1 to run against a real Payrexx account")
	}
	instance, secret := os.Getenv("PAYREXX_INSTANCE"), os.Getenv("PAYREXX_API_SECRET")
	if instance == "" || secret == "" {
		t.Fatal("PAYREXX_LIVE_TEST is set but PAYREXX_INSTANCE / PAYREXX_API_SECRET are not")
	}
	client, err := payrexx.NewClient(payrexx.Config{
		Instance: instance, APISecret: secret, Auth: auth,
	})
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	return client
}

// TestLiveGatewayRoundTrip is the one that matters: it is the only test that
// sends a request body, so it exercises the generated request models, and the
// only one guaranteed to decode a populated entity on an account with no
// history.
func TestLiveGatewayRoundTrip(t *testing.T) {
	client := liveClient(t, payrexx.AuthAPIKey)
	ctx := context.Background()

	req := payrexx.NewGatewayCreateRequest(100, "CHF")
	req.SetPurpose("payrexx-go-sdk integration test")
	req.SetReferenceId("sdk-live-test")

	created, resp, err := client.GatewayAPI.GatewayCreate(ctx).
		GatewayCreateRequest(*req).Execute()
	if err != nil {
		t.Fatalf("GatewayCreate: %v%s", err, body(err))
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("GatewayCreate: HTTP %d", resp.StatusCode)
	}
	if len(created.Data) != 1 {
		t.Fatalf("GatewayCreate returned %d gateways, want 1", len(created.Data))
	}
	gw := created.Data[0]
	id := payrexx.Value(gw.Id)
	if id == 0 {
		t.Fatal("GatewayCreate returned no id")
	}
	// Whatever else happens, do not leave the gateway behind.
	t.Cleanup(func() {
		if _, _, err := client.GatewayAPI.GatewayDelete(context.Background(), id).Execute(); err != nil {
			t.Errorf("GatewayDelete(%d): %v%s", id, err, body(err))
		}
	})

	if got := payrexx.Value(gw.ReferenceId); got != "sdk-live-test" {
		t.Errorf("referenceId = %q, want it echoed back", got)
	}
	if got := payrexx.Value(gw.Amount); got != 100 {
		t.Errorf("amount = %d, want 100", got)
	}
	// The link is the entire point of a gateway; a create that does not yield one
	// is a create the caller cannot use.
	if link := payrexx.Value(gw.Link); !strings.HasPrefix(link, "https://") {
		t.Errorf("link = %q, want a checkout URL", link)
	}
	if got := payrexx.Value(gw.Status); got != "waiting" {
		t.Errorf("status = %q, want waiting for an unpaid gateway", got)
	}
	// Documented as an integer, observed as a boolean -- the LIVE_OVERRIDES entry
	// in build_spec.py exists because decoding this failed.
	if gw.PreAuthorization == nil {
		t.Error("preAuthorization did not decode; check its type in the spec")
	}
	// [] when unset, an object when set. Typed as free-form for exactly that.
	if gw.Fields == nil {
		t.Error("fields did not decode")
	}

	read, _, err := client.GatewayAPI.GatewayRetrieve(ctx, id).Execute()
	if err != nil {
		t.Fatalf("GatewayRetrieve(%d): %v%s", id, err, body(err))
	}
	if len(read.Data) != 1 {
		t.Fatalf("GatewayRetrieve returned %d gateways, want 1", len(read.Data))
	}
	// Create and read are documented on separate pages with different field sets;
	// build_spec.py merges them into one entity, and this is the assertion that
	// the merge is legitimate.
	if got := payrexx.Value(read.Data[0].Hash); got != payrexx.Value(gw.Hash) {
		t.Errorf("retrieve gave hash %q, create gave %q", got, payrexx.Value(gw.Hash))
	}
}

// TestLiveListsDecode covers the read paths. The assertion is deliberately weak
// on content -- a fresh account has no history -- and strong on decoding: these
// call sites are where a wrong `data` type or a wrong envelope shows up.
func TestLiveListsDecode(t *testing.T) {
	client := liveClient(t, payrexx.AuthAPIKey)
	ctx := context.Background()

	t.Run("subscriptions", func(t *testing.T) {
		// Payrexx documents this response's `data` as a bare string, which its own
		// example contradicts. build_spec.py coerces it to an array of
		// Subscription; if that is wrong, this call fails to decode.
		got, _, err := client.SubscriptionAPI.SubscriptionList(ctx).Execute()
		if err != nil {
			t.Fatalf("SubscriptionList: %v%s", err, body(err))
		}
		if payrexx.Value(got.Status) != "success" {
			t.Errorf("status = %q, want success", payrexx.Value(got.Status))
		}
	})

	t.Run("transactions", func(t *testing.T) {
		got, _, err := client.TransactionAPI.TransactionList(ctx).Execute()
		if err != nil {
			t.Fatalf("TransactionList: %v%s", err, body(err))
		}
		if payrexx.Value(got.Status) != "success" {
			t.Errorf("status = %q, want success", payrexx.Value(got.Status))
		}
	})

	t.Run("bills", func(t *testing.T) {
		// The Bill pages spell the same fields both camelCase and snake_case;
		// build_spec.py keeps camelCase. The meta envelope here is also where the
		// isList/is_list typo lived.
		got, _, err := client.BillAPI.BillList(ctx).Execute()
		if err != nil {
			t.Fatalf("BillList: %v%s", err, body(err))
		}
		if payrexx.Value(got.Status) != "success" {
			t.Errorf("status = %q, want success", payrexx.Value(got.Status))
		}
		if got.Meta != nil && got.Meta.IsList == nil {
			t.Error("meta.isList did not decode; check the FIELD_TYPOS entry")
		}
	})

	t.Run("payouts", func(t *testing.T) {
		if _, _, err := client.PayoutAPI.PayoutList(ctx).Execute(); err != nil {
			t.Fatalf("PayoutList: %v%s", err, body(err))
		}
	})
}

// TestLiveSignatureScheme settles what the offline vectors cannot: the vectors
// prove this package encodes the parameters the way Payrexx's documented PHP
// snippet does, not that Payrexx's servers still accept the result. Payrexx's
// own PHP SDK abandoned this scheme, so it is plausible that they no longer do.
func TestLiveSignatureScheme(t *testing.T) {
	client := liveClient(t, payrexx.AuthSignature)

	_, resp, err := client.TransactionAPI.TransactionList(context.Background()).Execute()
	if err != nil {
		if resp != nil && (resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden) {
			t.Fatalf("AuthSignature was rejected with HTTP %d -- the scheme may be "+
				"retired; use AuthAPIKey and say so in the README%s", resp.StatusCode, body(err))
		}
		t.Fatalf("TransactionList under AuthSignature: %v%s", err, body(err))
	}
}

// TestLiveErrorShape checks that a failure decodes into the single Error type
// build_spec.py collapses Payrexx's three documented error shapes onto.
func TestLiveErrorShape(t *testing.T) {
	client := liveClient(t, payrexx.AuthAPIKey)

	// An id that cannot exist. The interesting part is the error body, not the
	// status code.
	_, resp, err := client.GatewayAPI.GatewayRetrieve(context.Background(), 1).Execute()
	if err == nil {
		t.Fatal("retrieving gateway id 1 succeeded; pick an id that cannot exist")
	}
	var apiErr *payrexx.GenericOpenAPIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("err is %T, want *payrexx.GenericOpenAPIError", err)
	}
	if resp == nil || resp.StatusCode < 400 {
		t.Fatalf("want a 4xx, got %v", resp)
	}
	model, ok := apiErr.Model().(payrexx.Error)
	if !ok {
		t.Fatalf("error model is %T, want payrexx.Error", apiErr.Model())
	}
	// Payrexx answers the classic endpoints with {status, message}; the Bill and
	// ECR families use {error: {...}}. Error carries both, and at least one group
	// has to be populated or the type is telling the caller nothing.
	if payrexx.Value(model.Message) == "" && model.Error == nil {
		t.Errorf("neither message nor error populated: %s", apiErr.Body())
	}
}

// body renders an API error's response body, which is where Payrexx puts the
// reason a call failed.
func body(err error) string {
	var apiErr *payrexx.GenericOpenAPIError
	if errors.As(err, &apiErr) && len(apiErr.Body()) > 0 {
		return "\n  body: " + string(apiErr.Body())
	}
	return ""
}
