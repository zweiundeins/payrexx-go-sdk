package payrexx_test

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/zweiundeins/payrexx-go-sdk/payrexx"
)

// A hosted checkout, end to end: create a gateway, send the shopper to the link
// it returns, and let the webhook -- not the browser redirect -- decide when the
// order is paid.
func Example() {
	client, err := payrexx.NewClient(payrexx.Config{
		Instance:  os.Getenv("PAYREXX_INSTANCE"),
		APISecret: os.Getenv("PAYREXX_API_SECRET"),
	})
	if err != nil {
		log.Fatal(err)
	}

	// Amounts are in the currency's smallest unit: 5000 is CHF 50.00.
	req := payrexx.NewGatewayCreateRequest(5000, "CHF")
	req.SetPurpose("Abo Basis")
	// referenceId comes back on the webhook; make it something you can look up.
	req.SetReferenceId("order-4711")
	req.SetSuccessRedirectUrl("https://example.ch/danke")
	req.SetFailedRedirectUrl("https://example.ch/fehlgeschlagen")
	req.SetCancelRedirectUrl("https://example.ch/abgebrochen")

	created, _, err := client.GatewayAPI.GatewayCreate(context.Background()).
		GatewayCreateRequest(*req).Execute()
	if err != nil {
		log.Fatal(err)
	}
	if len(created.Data) == 0 {
		log.Fatal("payrexx returned no gateway")
	}
	fmt.Println("send the shopper to:", payrexx.Value(created.Data[0].Link))
}

// Webhooks are the source of truth for payment state -- the shopper can always
// close the tab before the redirect. Verify, then act.
func ExampleParseWebhook() {
	signingKey := os.Getenv("PAYREXX_WEBHOOK_KEY")

	http.HandleFunc("/payrexx/webhook", func(w http.ResponseWriter, r *http.Request) {
		event, err := payrexx.ParseWebhook(r, signingKey)
		if err != nil {
			// Payrexx retries on a non-2xx, so refusing is safe. Do not log the
			// body of a delivery that failed verification as if it were real.
			status := http.StatusBadRequest
			if errors.Is(err, payrexx.ErrInvalidSignature) {
				status = http.StatusUnauthorized
			}
			http.Error(w, "rejected", status)
			return
		}

		tx := event.Transaction
		if tx == nil || !tx.IsLive() {
			w.WriteHeader(http.StatusOK) // nothing to do, but acknowledge it
			return
		}

		// Deliveries repeat: the same event arrives again whenever the handler is
		// slow, errors, or times out after 20 seconds. Key your write on the
		// transaction uuid so the second delivery is a no-op.
		switch payrexx.Value(tx.Status) {
		case payrexx.TransactionConfirmed:
			markPaid(tx.Reference(), payrexx.Value(tx.Uuid), payrexx.Value(tx.Amount))
		case payrexx.TransactionRefunded, payrexx.TransactionChargeback:
			markReversed(tx.Reference(), payrexx.Value(tx.Uuid))
		}
		w.WriteHeader(http.StatusOK)
	})
}

func markPaid(reference, uuid string, amount int32) {}
func markReversed(reference, uuid string)           {}
