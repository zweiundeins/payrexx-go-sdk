/*
Payrexx REST API — webhooks.

This file is hand-written and is NOT produced by the OpenAPI generator (it is
listed in .openapi-generator-ignore so `gen.sh` never overwrites it).

Payrexx's API reference documents no webhooks, so nothing here comes from the
generated spec. The types below follow the object tables published under
docs.payrexx.com/developer/guides/webhook — a different document from the API
reference, describing a payload that overlaps with but is not identical to the
entities the REST endpoints return, which is why these are separate types rather
than a reuse of Transaction and Subscription.

A webhook fires whenever a transaction or subscription changes status, and is
retried up to ten times over several days until the endpoint answers. Two
consequences for the handler: it has to answer within 20 seconds, and it has to
be idempotent, because a slow success is indistinguishable from a failure and
will be delivered again.
*/

package payrexx

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// SignatureHeader carries the webhook signature.
const SignatureHeader = "X-Webhook-Signature"

// ErrInvalidSignature is returned when a webhook body does not match its
// signature header. Reject the delivery; do not act on the payload.
var ErrInvalidSignature = errors.New("payrexx: webhook signature does not match")

// Transaction statuses, from the webhook documentation. A payment is money in
// the account at TransactionConfirmed and at nothing before it.
const (
	TransactionWaiting           = "waiting"
	TransactionConfirmed         = "confirmed"
	TransactionCancelled         = "cancelled"
	TransactionDeclined          = "declined"
	TransactionAuthorized        = "authorized"
	TransactionReserved          = "reserved"
	TransactionRefunded          = "refunded"
	TransactionPartiallyRefunded = "partially-refunded"
	TransactionRefundPending     = "refund_pending"
	TransactionChargeback        = "chargeback"
	TransactionError             = "error"
	TransactionDisputed          = "disputed"
	TransactionExpired           = "expired"
)

// Subscription statuses, from the webhook documentation.
const (
	// SubscriptionActive means further charges are to come.
	SubscriptionActive = "active"
	// SubscriptionOverdue means a charge failed and will be retried; a second
	// failure moves it to SubscriptionFailed.
	SubscriptionOverdue = "overdue"
	// SubscriptionFailed means the subscription could not be created, or a retry
	// after an overdue charge failed too.
	SubscriptionFailed = "failed"
	// SubscriptionInNotice means the shopper cancelled before the end date. The
	// remaining charges still follow.
	SubscriptionInNotice = "in_notice"
	// SubscriptionCancelled means the merchant cancelled with immediate effect.
	// No further charges follow.
	SubscriptionCancelled = "cancelled"
)

// WebhookEvent is the body of a webhook delivery. Exactly one of the pointers is
// populated, matching the object whose status changed.
type WebhookEvent struct {
	Transaction  *WebhookTransaction  `json:"transaction,omitempty"`
	Subscription *WebhookSubscription `json:"subscription,omitempty"`
	// Payout deliveries are documented only as an example payload, so they are
	// left raw rather than given a type this package cannot keep honest.
	Payout json.RawMessage `json:"payout,omitempty"`
	// Raw is the exact body the signature was verified against. Decode from this
	// to reach fields Payrexx has added since, and log this rather than the
	// re-marshalled struct.
	Raw json.RawMessage `json:"-"`
}

// WebhookTransaction is the transaction object as a webhook delivers it. Amounts
// are in the currency's smallest unit: 200 is CHF 2.00.
type WebhookTransaction struct {
	Id          *int32  `json:"id,omitempty"`
	Uuid        *string `json:"uuid,omitempty"`
	Amount      *int32  `json:"amount,omitempty"`
	ReferenceId *string `json:"referenceId,omitempty"`
	// Time is ISO 8601 as of 2025-04-29 ("2025-10-01T09:36:07+00:00"); older
	// deliveries carry "2025-10-01 09:36:07".
	Time   *string `json:"time,omitempty"`
	Status *string `json:"status,omitempty"`
	Lang   *string `json:"lang,omitempty"`
	Psp    *string `json:"psp,omitempty"`
	PspId  *int32  `json:"pspId,omitempty"`
	// Mode is TEST or LIVE. Check it before treating a payment as real money.
	Mode *string `json:"mode,omitempty"`
	// Type is E-Commerce, POS-Terminal or Tap to Pay.
	Type                *string              `json:"type,omitempty"`
	PosSerialNumber     *string              `json:"posSerialNumber,omitempty"`
	PosTerminalName     *string              `json:"posTerminalName,omitempty"`
	Refundable          *bool                `json:"refundable,omitempty"`
	PartiallyRefundable *bool                `json:"partiallyRefundable,omitempty"`
	Instance            *WebhookInstance     `json:"instance,omitempty"`
	Metadata            *WebhookMetadata     `json:"metadata,omitempty"`
	Invoice             *WebhookInvoice      `json:"invoice,omitempty"`
	Payment             *WebhookPayment      `json:"payment,omitempty"`
	Contact             *WebhookContact      `json:"contact,omitempty"`
	Subscription        *WebhookSubscription `json:"subscription,omitempty"`
	PageUuid            *string              `json:"pageUuid,omitempty"`
	// PayrexxFee is the total fee. Platform merchants receive it as "fee"
	// instead, in Fee.
	PayrexxFee *int32 `json:"payrexxFee,omitempty"`
	Fee        *int32 `json:"fee,omitempty"`
	// PreAuthorizationId is the tokenization's own id on a tokenization, or the
	// source tokenization's id on a charge against one. Absent otherwise.
	PreAuthorizationId *int32 `json:"preAuthorizationId,omitempty"`
	// OriginalTransactionId and OriginalTransactionUuid identify the charged
	// transaction, and are present on refunded and partially-refunded statuses.
	OriginalTransactionId   *int32  `json:"originalTransactionId,omitempty"`
	OriginalTransactionUuid *string `json:"originalTransactionUuid,omitempty"`
	PayoutUuid              *string `json:"payoutUuid,omitempty"`
	// The remaining fields are marketplace-only.
	AssociatedSourceTransactionUuids []string `json:"associatedSourceTransactionUuids,omitempty"`
	AssociatedSourceTransactionIds   []int32  `json:"associatedSourceTransactionIds,omitempty"`
	SourceTransactionUuid            *string  `json:"sourceTransactionUuid,omitempty"`
	SourceTransactionId              *string  `json:"sourceTransactionId,omitempty"`
}

// WebhookSubscription is the subscription object as a webhook delivers it. It is
// also embedded in a transaction when that transaction is a subscription charge.
type WebhookSubscription struct {
	Id     *int32  `json:"id,omitempty"`
	Uuid   *string `json:"uuid,omitempty"`
	Status *string `json:"status,omitempty"`
	// Start, End and ValidUntil are dates ("1970-01-01"). ValidUntil is the next
	// pay date.
	Start      *string `json:"start,omitempty"`
	End        *string `json:"end,omitempty"`
	ValidUntil *string `json:"valid_until,omitempty"`
	// PaymentInterval is an ISO 8601 duration as PHP's DateInterval takes it,
	// e.g. "P1M" for monthly.
	PaymentInterval *string         `json:"paymentInterval,omitempty"`
	Invoice         *WebhookInvoice `json:"invoice,omitempty"`
	Contact         *WebhookContact `json:"contact,omitempty"`
}

// WebhookInvoice carries what was bought. ReferenceId and PaymentLink are how a
// delivery is matched back to the checkout that produced it.
type WebhookInvoice struct {
	Number   *string          `json:"number,omitempty"`
	Products []WebhookProduct `json:"products,omitempty"`
	Currency *string          `json:"currency,omitempty"`
	// Test is 1 for a simulated payment, 0 for a real one.
	Test *int32 `json:"test,omitempty"`
	// ReferenceId is the reference your system passed when creating the gateway.
	ReferenceId *string `json:"referenceId,omitempty"`
	Purpose     *string `json:"purpose,omitempty"`
	// PaymentRequestId is the id of the invoice or gateway created through the API.
	PaymentRequestId *int32               `json:"paymentRequestId,omitempty"`
	PaymentLink      *WebhookPaymentLink  `json:"paymentLink,omitempty"`
	ShippingAmount   *int32               `json:"shippingAmount,omitempty"`
	Discount         *WebhookDiscount     `json:"discount,omitempty"`
	OriginalAmount   *int32               `json:"originalAmount,omitempty"`
	RefundedAmount   *int32               `json:"refundedAmount,omitempty"`
	CustomFields     []WebhookCustomField `json:"custom_fields,omitempty"`
	// GoogleAnalyticProducts mirrors Products in Google's item schema.
	GoogleAnalyticProducts json.RawMessage `json:"googleAnalyticProducts,omitempty"`
}

// WebhookPaymentLink identifies the payment link a transaction came through.
type WebhookPaymentLink struct {
	Hash        *string `json:"hash,omitempty"`
	ReferenceId *string `json:"referenceId,omitempty"`
	Email       *string `json:"email,omitempty"`
}

// WebhookProduct is one line of an invoice.
type WebhookProduct struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Price       *int32  `json:"price,omitempty"`
	Quantity    *int32  `json:"quantity,omitempty"`
	Sku         *string `json:"sku,omitempty"`
	VatRate     *string `json:"vatRate,omitempty"`
}

// WebhookDiscount is the discount applied to an invoice.
type WebhookDiscount struct {
	Code       *string `json:"code,omitempty"`
	Amount     *int32  `json:"amount,omitempty"`
	Percentage *string `json:"percentage,omitempty"`
}

// WebhookCustomField is one of the extra fields configured on the payment page.
type WebhookCustomField struct {
	Type  *string `json:"type,omitempty"`
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// WebhookContact is the shopper. Payrexx keys contacts by email address and
// updates them on repeat payments, so the values here are the ones sent with
// this payment, falling back to the stored contact.
type WebhookContact struct {
	Id         *int32  `json:"id,omitempty"`
	Uuid       *string `json:"uuid,omitempty"`
	Title      *string `json:"title,omitempty"`
	Firstname  *string `json:"firstname,omitempty"`
	Lastname   *string `json:"lastname,omitempty"`
	Company    *string `json:"company,omitempty"`
	Street     *string `json:"street,omitempty"`
	Zip        *string `json:"zip,omitempty"`
	Place      *string `json:"place,omitempty"`
	Country    *string `json:"country,omitempty"`
	CountryISO *string `json:"countryISO,omitempty"`
	// Phone and DateOfBirth arrive in whatever format the shopper typed.
	Phone              *string `json:"phone,omitempty"`
	Email              *string `json:"email,omitempty"`
	DateOfBirth        *string `json:"date_of_birth,omitempty"`
	DeliveryTitle      *string `json:"delivery_title,omitempty"`
	DeliveryFirstname  *string `json:"delivery_firstname,omitempty"`
	DeliveryLastname   *string `json:"delivery_lastname,omitempty"`
	DeliveryCompany    *string `json:"delivery_company,omitempty"`
	DeliveryStreet     *string `json:"delivery_street,omitempty"`
	DeliveryZip        *string `json:"delivery_zip,omitempty"`
	DeliveryPlace      *string `json:"delivery_place,omitempty"`
	DeliveryCountry    *string `json:"delivery_country,omitempty"`
	DeliveryCountryISO *string `json:"delivery_countryISO,omitempty"`
}

// WebhookPayment describes how the shopper paid.
type WebhookPayment struct {
	Brand *string `json:"brand,omitempty"`
	// Wallet is apple-pay, google-pay, samsung-pay, visa-click-to-pay or
	// mastercard-click-to-pay, else nil.
	Wallet *string `json:"wallet,omitempty"`
	// CardNumber is already masked by Payrexx ("424242xxxxxx4242"). Present only
	// alongside Expiry.
	CardNumber                   *string         `json:"cardNumber,omitempty"`
	Expiry                       *string         `json:"expiry,omitempty"`
	CardholderName               *string         `json:"cardholderName,omitempty"`
	CardCountry                  *string         `json:"cardCountry,omitempty"`
	PurchaseOnInvoiceInformation json.RawMessage `json:"purchaseOnInvoiceInformation,omitempty"`
	// The remaining fields are marketplace-only.
	OriginInstanceUuid      *string `json:"originInstanceUuid,omitempty"`
	OriginInstanceId        *string `json:"originInstanceId,omitempty"`
	OriginInstanceName      *string `json:"originInstanceName,omitempty"`
	DestinationInstanceUuid *string `json:"destinationInstanceUuid,omitempty"`
	DestinationInstanceId   *string `json:"destinationInstanceId,omitempty"`
	DestinationInstanceName *string `json:"destinationInstanceName,omitempty"`
}

// WebhookInstance names the merchant account the transaction belongs to.
type WebhookInstance struct {
	Name *string `json:"name,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
}

// WebhookMetadata carries provider-specific extras. Every field is optional.
type WebhookMetadata struct {
	PaypalBillingAgreementId *string `json:"paypalBillingAgreementId,omitempty"`
	// DeclineCode and Message describe a failure.
	DeclineCode *string `json:"decline_code,omitempty"`
	Message     *string `json:"message,omitempty"`
	Brand       *string `json:"brand,omitempty"`
}

// Value dereferences an optional field, returning the zero value when it is
// absent. Every field of a webhook payload is optional -- Payrexx omits what
// does not apply to a given transaction -- so this is the safe way to read one:
//
//	if payrexx.Value(tx.Status) == payrexx.TransactionConfirmed { ... }
func Value[T any](p *T) T {
	if p == nil {
		var zero T
		return zero
	}
	return *p
}

// IsLive reports whether the transaction moved real money. Payrexx marks
// simulated payments with mode TEST, and a test transaction reaching production
// business logic is the classic integration mistake.
func (t *WebhookTransaction) IsLive() bool {
	return t != nil && Value(t.Mode) == "LIVE"
}

// Reference returns the referenceId the invoice carries — the value passed when
// the gateway was created, and the documented way to match a delivery back to
// your own record. It falls back to the transaction's own referenceId.
func (t *WebhookTransaction) Reference() string {
	if t == nil {
		return ""
	}
	if t.Invoice != nil && Value(t.Invoice.ReferenceId) != "" {
		return Value(t.Invoice.ReferenceId)
	}
	return Value(t.ReferenceId)
}

// VerifyWebhookSignature reports whether body carries the signature Payrexx
// computed for it. The signature is a lowercase hex HMAC-SHA256 over the raw
// body, keyed with the signing key as plain UTF-8 — the key is not base64 and
// the body must not be re-serialized before hashing.
func VerifyWebhookSignature(signingKey string, body []byte, signature string) bool {
	mac := hmac.New(sha256.New, []byte(signingKey))
	mac.Write(body)
	want := hex.EncodeToString(mac.Sum(nil))
	// Compare over the decoded bytes so a difference in hex casing is not
	// mistaken for a forgery, and in constant time so the comparison leaks
	// nothing about the expected value.
	got, err := hex.DecodeString(strings.TrimSpace(signature))
	if err != nil {
		return false
	}
	expected, _ := hex.DecodeString(want)
	return hmac.Equal(expected, got)
}

// ParseWebhook reads, verifies and decodes a webhook delivery.
//
// It returns ErrInvalidSignature if the body does not match the
// X-Webhook-Signature header, and never decodes an unverified body. Reject the
// delivery on any error; Payrexx retries, so answering non-2xx is safe.
//
// The handler must be idempotent: a delivery that times out — Payrexx allows 20
// seconds — is retried even though it succeeded.
func ParseWebhook(r *http.Request, signingKey string) (*WebhookEvent, error) {
	if signingKey == "" {
		return nil, errors.New("payrexx: webhook signing key is required")
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("payrexx: reading webhook body: %w", err)
	}
	if !VerifyWebhookSignature(signingKey, body, r.Header.Get(SignatureHeader)) {
		return nil, ErrInvalidSignature
	}
	if ct := r.Header.Get("Content-Type"); ct != "" && !strings.Contains(ct, "json") {
		return nil, fmt.Errorf(
			"payrexx: webhook content type %q is not JSON; set the webhook's content type to "+
				"JSON in the Payrexx merchant account (the PHP-Post form encoding is not decoded here)",
			ct)
	}

	var event WebhookEvent
	if err := json.Unmarshal(body, &event); err != nil {
		return nil, fmt.Errorf("payrexx: decoding webhook body: %w", err)
	}
	event.Raw = body
	return &event, nil
}
