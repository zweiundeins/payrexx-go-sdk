package payrexx

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

// The signature vectors below were produced by running Payrexx's own documented
// snippet -- base64_encode(hash_hmac('sha256', http_build_query($params), $secret, true))
// -- against PHP 8.4, so they pin this package to the reference implementation's
// encoding rather than to an assumption about it. Regenerate with:
//
//	php -r '$p=[...]; echo http_build_query($p,"","&"), "\n",
//	        base64_encode(hash_hmac("sha256", http_build_query($p,"","&"), "test-secret", true)), "\n";'
const testSecret = "test-secret"

func TestPHPQueryFromJSONMatchesPHP(t *testing.T) {
	cases := []struct {
		name  string
		body  string
		query string
		sig   string
	}{
		{
			name:  "scalars",
			body:  `{"model":"Page","id":17}`,
			query: "model=Page&id=17",
			sig:   "zp8vWZDr2rL9gGz3TxuYoxR6/rZFt2mJrqkLj5cwgXI=",
		},
		{
			// Space becomes "+", and PHP escapes ~ ! * ( ) where Go's
			// url.QueryEscape would leave "~" untouched.
			name:  "spaces and specials",
			body:  `{"purpose":"Abo Basis","title":"a&b=c+d","sym":"ä~!*()"}`,
			query: "purpose=Abo+Basis&title=a%26b%3Dc%2Bd&sym=%C3%A4%7E%21%2A%28%29",
			sig:   "LubCkW6b2BWdGfMH5+5mrWDSqIiX2KOGnisZDF7fr3I=",
		},
		{
			name:  "nested object",
			body:  `{"contact":{"email":"a@b.ch","name":"Jane Doe"}}`,
			query: "contact%5Bemail%5D=a%40b.ch&contact%5Bname%5D=Jane+Doe",
			sig:   "FtyH7rOrETdLP1b0OPxdwQFva0HAkWi9wa8+ajqSpoM=",
		},
		{
			name:  "list",
			body:  `{"products":["a","b"]}`,
			query: "products%5B0%5D=a&products%5B1%5D=b",
			sig:   "GcCBoqCfN13bzEpwrBvktSp6wnx5435o4aFE5FuXdTs=",
		},
		{
			name:  "list of objects",
			body:  `{"fields":[{"name":"x","value":1},{"name":"y","value":2}]}`,
			query: "fields%5B0%5D%5Bname%5D=x&fields%5B0%5D%5Bvalue%5D=1&fields%5B1%5D%5Bname%5D=y&fields%5B1%5D%5Bvalue%5D=2",
			sig:   "U2K9ROs24FTf7dwAryUqFx/JbaKuMyvekXtKV9oTKW0=",
		},
		{
			// http_build_query renders booleans as 1 and 0 and drops nulls
			// entirely -- "c" is absent from the signed string.
			name:  "booleans and null",
			body:  `{"a":true,"b":false,"c":null,"d":0}`,
			query: "a=1&b=0&d=0",
			sig:   "IxQIzEIateU8B5sj05mbfBBRDXCxE0I5Rp4U059eEbU=",
		},
		{
			// 8.1 has to stay "8.1"; decoding through float64 would render it
			// as 8.100000 and change the signature.
			name:  "numbers keep their literal form",
			body:  `{"amount":1000,"rate":8.1}`,
			query: "amount=1000&rate=8.1",
			sig:   "fBVXSxJ2Y12pEnT3eiljvckcHYQcZ+DIB0BNCKN+AA0=",
		},
		{
			name:  "empty object",
			body:  `{}`,
			query: "",
			sig:   "pBvG2B1kE1dq4JlJleCtiaQW7Jc4lRXDYE9HciEi7us=",
		},
	}

	tr := &authTransport{secret: testSecret}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := phpQueryFromJSON([]byte(tc.body))
			if err != nil {
				t.Fatalf("phpQueryFromJSON: %v", err)
			}
			if got != tc.query {
				t.Errorf("query string\n got %q\nwant %q", got, tc.query)
			}
			if sig := tr.signature(got); sig != tc.sig {
				t.Errorf("signature\n got %q\nwant %q", sig, tc.sig)
			}
		})
	}
}

func TestPHPQueryFromURLQueryMatchesPHP(t *testing.T) {
	// url.Values.Encode sorts and uses Go's escaping; re-rendering has to reach
	// PHP's, including "~" -> %7E, while keeping the transmitted order.
	raw := url.Values{"model": {"Page"}, "id": {"17"}}.Encode()
	if raw != "id=17&model=Page" {
		t.Fatalf("precondition: url.Values.Encode gave %q", raw)
	}
	if got, want := phpQueryFromURLQuery(raw), "id=17&model=Page"; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
	if got, want := phpQueryFromURLQuery("q=a~b+c"), "q=a%7Eb+c"; got != want {
		t.Errorf("tilde: got %q, want %q", got, want)
	}
}

func TestAPIKeySchemeSendsHeaderAndInstance(t *testing.T) {
	var got *http.Request
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got = r
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"status":"success","data":[{"id":42,"status":"active"}]}`)
	}))
	defer srv.Close()

	client, err := NewClient(Config{Instance: "demo", APISecret: testSecret, BaseURL: srv.URL})
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	sub, _, err := client.SubscriptionAPI.SubscriptionRetrieve(context.Background(), 42).Execute()
	if err != nil {
		t.Fatalf("SubscriptionRetrieve: %v", err)
	}

	if key := got.Header.Get("X-API-KEY"); key != testSecret {
		t.Errorf("X-API-KEY = %q, want %q", key, testSecret)
	}
	if inst := got.URL.Query().Get("instance"); inst != "demo" {
		t.Errorf("instance = %q, want %q", inst, "demo")
	}
	if sig := got.URL.Query().Get("ApiSignature"); sig != "" {
		t.Errorf("ApiSignature = %q, want it absent under AuthAPIKey", sig)
	}
	if got.URL.Path != "/Subscription/42/" {
		t.Errorf("path = %q, want /Subscription/42/", got.URL.Path)
	}
	if sub.Data == nil || len(sub.Data) != 1 || sub.Data[0].GetId() != 42 {
		t.Errorf("decoded %+v, want one subscription with id 42", sub.Data)
	}
}

func TestSignatureSchemeSignsQueryAndExcludesInstance(t *testing.T) {
	var got *http.Request
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got = r
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"status":"success","data":[]}`)
	}))
	defer srv.Close()

	client, err := NewClient(Config{
		Instance: "demo", APISecret: testSecret, Auth: AuthSignature, BaseURL: srv.URL,
	})
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	if _, _, err := client.SubscriptionAPI.SubscriptionList(context.Background()).Execute(); err != nil {
		t.Fatalf("SubscriptionList: %v", err)
	}

	if key := got.Header.Get("X-API-KEY"); key != "" {
		t.Errorf("X-API-KEY = %q, want it absent under AuthSignature", key)
	}
	if inst := got.URL.Query().Get("instance"); inst != "demo" {
		t.Errorf("instance = %q, want %q", inst, "demo")
	}

	// The list call sends no parameters of its own, so the signed string is
	// empty -- and `instance`, appended after signing, must not be in it.
	tr := &authTransport{secret: testSecret}
	if sig, want := got.URL.Query().Get("ApiSignature"), tr.signature(""); sig != want {
		t.Errorf("ApiSignature = %q, want %q (signed over the empty parameter set)", sig, want)
	}
	if i, j := strings.Index(got.URL.RawQuery, "ApiSignature"), strings.Index(got.URL.RawQuery, "instance"); i > j {
		t.Errorf("instance was appended before the signature in %q", got.URL.RawQuery)
	}
}

func TestSignatureSchemeSignsJSONBody(t *testing.T) {
	var body []byte
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ = io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"status":"success","data":[{"id":7}]}`)
	}))
	defer srv.Close()

	client, err := NewClient(Config{
		Instance: "demo", APISecret: testSecret, Auth: AuthSignature, BaseURL: srv.URL,
	})
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	req := NewSubscriptionCreateRequest("1", "3", "1000", "CHF", "Abo Basis", "P1M", "P1Y", "P1M")
	if _, _, err := client.SubscriptionAPI.SubscriptionCreate(context.Background()).
		SubscriptionCreateRequest(*req).Execute(); err != nil {
		t.Fatalf("SubscriptionCreate: %v", err)
	}

	var sent map[string]json.RawMessage
	if err := json.Unmarshal(body, &sent); err != nil {
		t.Fatalf("body %s: %v", body, err)
	}
	sig, ok := sent["ApiSignature"]
	if !ok {
		t.Fatalf("body carries no ApiSignature: %s", body)
	}
	// The signature covers the body as it was before ApiSignature was spliced in.
	cut := strings.LastIndex(string(body), `,"ApiSignature":`)
	if cut < 0 {
		t.Fatalf("ApiSignature was not appended last: %s", body)
	}
	original := string(body)[:cut] + "}"
	query, err := phpQueryFromJSON([]byte(original))
	if err != nil {
		t.Fatalf("phpQueryFromJSON: %v", err)
	}
	tr := &authTransport{secret: testSecret}
	want, _ := json.Marshal(tr.signature(query))
	if string(sig) != string(want) {
		t.Errorf("ApiSignature = %s, want %s (over %q)", sig, want, query)
	}
	if strings.Contains(query, "instance") {
		t.Errorf("instance leaked into the signed parameters: %q", query)
	}
}

func TestNewClientRejectsIncompleteConfig(t *testing.T) {
	if _, err := NewClient(Config{APISecret: "x"}); err == nil {
		t.Error("missing Instance was accepted")
	}
	if _, err := NewClient(Config{Instance: "demo"}); err == nil {
		t.Error("missing APISecret was accepted")
	}
}

func TestNewClientLayersUnderCallerTransport(t *testing.T) {
	// A caller's own Transport must still run -- authentication wraps it rather
	// than replacing it.
	inner := &countingTransport{next: http.DefaultTransport}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"status":"success","data":[]}`)
	}))
	defer srv.Close()

	client, err := NewClient(Config{
		Instance:   "demo",
		APISecret:  testSecret,
		BaseURL:    srv.URL,
		HTTPClient: &http.Client{Transport: inner},
	})
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	if _, _, err := client.SubscriptionAPI.SubscriptionList(context.Background()).Execute(); err != nil {
		t.Fatalf("SubscriptionList: %v", err)
	}
	if inner.calls != 1 {
		t.Errorf("caller transport saw %d calls, want 1", inner.calls)
	}
	if inner.lastAuth != testSecret {
		t.Errorf("caller transport saw X-API-KEY %q; authentication should be applied above it", inner.lastAuth)
	}
}

type countingTransport struct {
	next     http.RoundTripper
	calls    int
	lastAuth string
}

func (c *countingTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	c.calls++
	c.lastAuth = r.Header.Get("X-API-KEY")
	return c.next.RoundTrip(r)
}
