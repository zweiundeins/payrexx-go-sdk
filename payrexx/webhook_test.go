package payrexx

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"net/http/httptest"
	"strings"
	"testing"
)

const webhookKey = "webhook-signing-key"

func sign(t *testing.T, body string) string {
	t.Helper()
	mac := hmac.New(sha256.New, []byte(webhookKey))
	mac.Write([]byte(body))
	return hex.EncodeToString(mac.Sum(nil))
}

// A confirmed subscription charge, shaped after the example on
// docs.payrexx.com/developer/guides/webhook/transaction.
const chargeBody = `{
  "transaction": {
    "id": 1234,
    "uuid": "1122aabb",
    "amount": 5000,
    "status": "confirmed",
    "time": "2026-07-20T09:36:07+00:00",
    "mode": "LIVE",
    "type": "E-Commerce",
    "psp": "Native_PSP",
    "pspId": 44,
    "payrexxFee": 145,
    "refundable": true,
    "partiallyRefundable": false,
    "instance": {"name": "demo", "uuid": "aabb1122"},
    "invoice": {
      "number": "IV_123",
      "currency": "CHF",
      "test": 0,
      "referenceId": "sub-42",
      "originalAmount": 5000,
      "refundedAmount": 0,
      "paymentLink": {"hash": "78cabeef23f875c8fec25f2339d2894d", "email": "jane@example.ch"},
      "products": [{"name": "Abo Basis", "price": 5000, "quantity": 1, "vatRate": "8.10"}],
      "custom_fields": [{"type": "email", "name": "E-Mail", "value": "jane@example.ch"}]
    },
    "payment": {"brand": "visa", "wallet": null, "cardNumber": "424242xxxxxx4242", "expiry": "2028-11"},
    "contact": {"id": 7, "email": "jane@example.ch", "firstname": "Jane", "lastname": "Doe"},
    "subscription": {"id": 42, "status": "active", "valid_until": "2026-08-20", "paymentInterval": "P1M"}
  }
}`

func TestParseWebhookDecodesACharge(t *testing.T) {
	r := httptest.NewRequest("POST", "/webhook", strings.NewReader(chargeBody))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set(SignatureHeader, sign(t, chargeBody))

	event, err := ParseWebhook(r, webhookKey)
	if err != nil {
		t.Fatalf("ParseWebhook: %v", err)
	}
	tx := event.Transaction
	if tx == nil {
		t.Fatal("no transaction in the event")
	}
	if got := Value(tx.Status); got != TransactionConfirmed {
		t.Errorf("status = %q, want %q", got, TransactionConfirmed)
	}
	if got := Value(tx.Amount); got != 5000 {
		t.Errorf("amount = %d, want 5000 (smallest currency unit)", got)
	}
	if got := Value(tx.Mode); got != "LIVE" {
		t.Errorf("mode = %q, want LIVE", got)
	}
	// referenceId on the invoice is what maps a delivery back to our own record.
	if got := tx.Reference(); got != "sub-42" {
		t.Errorf("Reference() = %q, want sub-42 (the invoice referenceId)", got)
	}
	if !tx.IsLive() {
		t.Error("IsLive() = false for mode LIVE")
	}
	if got := Value(tx.Subscription.Status); got != SubscriptionActive {
		t.Errorf("subscription.status = %q, want %q", got, SubscriptionActive)
	}
	if got := Value(tx.Subscription.ValidUntil); got != "2026-08-20" {
		t.Errorf("subscription.valid_until = %q, want 2026-08-20", got)
	}
	if got := Value(tx.Payment.CardNumber); got != "424242xxxxxx4242" {
		t.Errorf("payment.cardNumber = %q, want the masked pan", got)
	}
	if tx.Payment.Wallet != nil {
		t.Errorf("payment.wallet = %v, want nil for an explicit JSON null", *tx.Payment.Wallet)
	}
	if got := len(tx.Invoice.Products); got != 1 {
		t.Fatalf("invoice.products has %d entries, want 1", got)
	}
	if got := Value(tx.Invoice.Products[0].VatRate); got != "8.10" {
		t.Errorf("vatRate = %q, want the string 8.10 (it is not a number upstream)", got)
	}
	if string(event.Raw) != chargeBody {
		t.Error("Raw is not the byte-exact body that was verified")
	}
}

func TestParseWebhookRejectsAForgedBody(t *testing.T) {
	r := httptest.NewRequest("POST", "/webhook", strings.NewReader(chargeBody))
	r.Header.Set("Content-Type", "application/json")
	// A signature over a different body: what an attacker replaying an old
	// delivery with an edited amount would present.
	r.Header.Set(SignatureHeader, sign(t, `{"transaction":{"amount":1}}`))

	if _, err := ParseWebhook(r, webhookKey); !errors.Is(err, ErrInvalidSignature) {
		t.Fatalf("err = %v, want ErrInvalidSignature", err)
	}
}

func TestParseWebhookRejectsAMissingSignature(t *testing.T) {
	r := httptest.NewRequest("POST", "/webhook", strings.NewReader(chargeBody))
	r.Header.Set("Content-Type", "application/json")

	if _, err := ParseWebhook(r, webhookKey); !errors.Is(err, ErrInvalidSignature) {
		t.Fatalf("err = %v, want ErrInvalidSignature", err)
	}
}

func TestParseWebhookRejectsAWrongKey(t *testing.T) {
	r := httptest.NewRequest("POST", "/webhook", strings.NewReader(chargeBody))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set(SignatureHeader, sign(t, chargeBody))

	if _, err := ParseWebhook(r, "someone-elses-key"); !errors.Is(err, ErrInvalidSignature) {
		t.Fatalf("err = %v, want ErrInvalidSignature", err)
	}
}

func TestParseWebhookRequiresASigningKey(t *testing.T) {
	// An empty key must not degrade into "verify against the empty-key HMAC",
	// which an attacker can compute.
	r := httptest.NewRequest("POST", "/webhook", strings.NewReader(chargeBody))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set(SignatureHeader, sign(t, chargeBody))

	if _, err := ParseWebhook(r, ""); err == nil {
		t.Fatal("an empty signing key was accepted")
	}
}

func TestVerifyWebhookSignatureAcceptsUppercaseHex(t *testing.T) {
	body := []byte(chargeBody)
	lower := sign(t, chargeBody)
	if !VerifyWebhookSignature(webhookKey, body, lower) {
		t.Error("lowercase hex was rejected")
	}
	if !VerifyWebhookSignature(webhookKey, body, strings.ToUpper(lower)) {
		t.Error("uppercase hex was rejected; casing is not a forgery")
	}
	if VerifyWebhookSignature(webhookKey, body, "not-hex") {
		t.Error("a non-hex signature was accepted")
	}
	if VerifyWebhookSignature(webhookKey, body, "") {
		t.Error("an empty signature was accepted")
	}
}

func TestParseWebhookRejectsFormEncodedDeliveries(t *testing.T) {
	// Payrexx offers a "PHP-Post" content type that this package does not decode;
	// the error has to say so rather than fail as malformed JSON.
	body := "transaction%5Bid%5D=1234"
	r := httptest.NewRequest("POST", "/webhook", strings.NewReader(body))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Set(SignatureHeader, sign(t, body))

	_, err := ParseWebhook(r, webhookKey)
	if err == nil {
		t.Fatal("a form-encoded delivery was accepted")
	}
	if !strings.Contains(err.Error(), "content type") {
		t.Errorf("err = %v, want it to name the content type", err)
	}
}

func TestParseWebhookDecodesASubscriptionOnlyEvent(t *testing.T) {
	body := `{"subscription":{"id":42,"uuid":"aabb1122","status":"cancelled","start":"2026-01-01","end":"2026-07-20"}}`
	r := httptest.NewRequest("POST", "/webhook", strings.NewReader(body))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set(SignatureHeader, sign(t, body))

	event, err := ParseWebhook(r, webhookKey)
	if err != nil {
		t.Fatalf("ParseWebhook: %v", err)
	}
	if event.Transaction != nil {
		t.Error("a transaction was decoded from a subscription-only event")
	}
	if got := Value(event.Subscription.Status); got != SubscriptionCancelled {
		t.Errorf("status = %q, want %q", got, SubscriptionCancelled)
	}
	if got := Value(event.Subscription.End); got != "2026-07-20" {
		t.Errorf("end = %q, want 2026-07-20", got)
	}
}
