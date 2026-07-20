# payrexx-go-sdk

[![CI](https://github.com/zweiundeins/payrexx-go-sdk/actions/workflows/ci.yml/badge.svg)](https://github.com/zweiundeins/payrexx-go-sdk/actions/workflows/ci.yml)
[![Upstream watch](https://github.com/zweiundeins/payrexx-go-sdk/actions/workflows/upstream-watch.yml/badge.svg)](https://github.com/zweiundeins/payrexx-go-sdk/actions/workflows/upstream-watch.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/zweiundeins/payrexx-go-sdk/payrexx.svg)](https://pkg.go.dev/github.com/zweiundeins/payrexx-go-sdk/payrexx)
[![Go Report Card](https://goreportcard.com/badge/github.com/zweiundeins/payrexx-go-sdk)](https://goreportcard.com/report/github.com/zweiundeins/payrexx-go-sdk)

A Go client for the [Payrexx](https://payrexx.com) REST API.

The second badge is the one that matters: it is green when this client was last
confirmed against the live Payrexx API, which is a stronger claim than a passing
unit test. See [Continuous verification](#continuous-verification).

Payrexx publishes an official SDK for PHP and community ones for Node and C#, but
**not Go**, and — unlike most payment providers — no OpenAPI document either. This
module fills both gaps: [`scripts/build_spec.py`](scripts/build_spec.py) assembles
one spec from the per-endpoint fragments Payrexx's documentation site serves, and
the client is generated from it, so it tracks the upstream API. The hand-written
pieces are the two things a generator cannot produce: authentication (Payrexx
declares no security scheme) and webhooks (absent from the API reference
entirely).

- Module: `github.com/zweiundeins/payrexx-go-sdk`
- Import path / package: `github.com/zweiundeins/payrexx-go-sdk/payrexx` (package `payrexx`)
- Generated with [openapi-generator](https://openapi-generator.tech) `7.23.0`
  (pinned in `openapitools.json`) from `openapi.json` (pinned in-repo, and itself
  reproducible via `./gen.sh --refresh-spec`)
- Targets Payrexx API **v1.16**
- No third-party dependencies — pure standard library

## Install

```sh
go get github.com/zweiundeins/payrexx-go-sdk/payrexx
```

## Quick start

```go
import (
    "context"
    "os"

    "github.com/zweiundeins/payrexx-go-sdk/payrexx"
)

client, err := payrexx.NewClient(payrexx.Config{
    Instance:  "example",                          // example.payrexx.com -> "example"
    APISecret: os.Getenv("PAYREXX_API_SECRET"),    // from the admin panel, API settings
})
if err != nil {
    return err
}

// Amounts are always in the currency's smallest unit: 5000 is CHF 50.00.
req := payrexx.NewGatewayCreateRequest(5000, "CHF")
req.SetPurpose("Abo Basis")
req.SetReferenceId("order-4711")   // comes back on the webhook
req.SetSuccessRedirectUrl("https://example.ch/danke")

created, _, err := client.GatewayAPI.GatewayCreate(context.Background()).
    GatewayCreateRequest(*req).Execute()
// created.Data[0].Link is the hosted checkout URL to send the shopper to.
```

Every operation is a request builder terminated by `Execute()`, returning
`(*Model, *http.Response, error)`. API errors come back as
`*payrexx.GenericOpenAPIError`; use `.Body()` for the raw response and `.Model()`
for the parsed `payrexx.Error`. See [`payrexx/example_test.go`](payrexx/example_test.go)
for a full checkout flow.

Note the shape of every response: Payrexx wraps results in
`{"status": "success", "data": [...]}`, so even a single-entity read arrives as a
one-element slice. Check the length before indexing.

You never pass the `instance` parameter that Payrexx requires on every endpoint —
the client appends it. See *Authentication*.

## Services

| Service | What |
|---|---|
| `GatewayAPI` | Hosted checkout pages. `GatewayCreate` returns the `link` you redirect to. |
| `SubscriptionAPI` | Recurring billing: create, read, list, update, cancel. |
| `TransactionAPI` | Read transactions, cancel a waiting one, manage tokenizations. |
| `InvoiceAPI` | `InvoiceCreatePaylink` — a reusable payment link. |
| `BillAPI` | Invoices in the accounting sense (Payrexx's `/Bill/` resource). |
| `PayoutAPI` | Payouts, with or without their underlying transfers. |
| `PaymentProviderAPI` | The PSPs enabled on the instance, and their ids. |
| `EcrAPI` | POS terminals: pairing and card-present payments. |

Payrexx's own naming is worth a warning: the documentation pages titled *Invoice*
describe the `/Bill/` resource (`BillAPI`), while the page titled *Create a
Paylink* is `POST /Invoice/` (`InvoiceAPI.InvoiceCreatePaylink`). The services
here are named after the HTTP resource, which is the unambiguous one.

Rate limit: **600 requests per 5 minutes**, enforced by AWS WAF. Exceeding it
returns `405` and then `403` — neither of which reads like a rate limit, so back
off exponentially on both.

## Authentication

Credentials are the **instance name** (the subdomain you reach your account on)
and its **API secret** (admin panel → API settings). Keep the secret server-side.

Payrexx's OpenAPI fragments declare no security scheme, so authentication is
hand-written in [`payrexx/auth.go`](payrexx/auth.go) and applied as a transparent
`http.RoundTripper` that covers every generated call. Two schemes, selectable via
`Config.Auth`:

| Scheme | `Config.Auth` | How | Notes |
|---|---|---|---|
| **API key** (default) | `AuthAPIKey` | `X-API-KEY: <secret>` | What Payrexx recommends and what payrexx-php sends. |
| **API signature** (legacy) | `AuthSignature` | `ApiSignature` parameter, base64 HMAC-SHA256 | Keeps the secret off the wire. See the caveat below. |

The same transport appends the `instance` query parameter to every request.
`build_spec.py` strips it from the generated operation signatures, because it is
constant for a client's lifetime and has to be excluded from the signed parameter
set — putting it in one place is the only way both stay true.

**On `AuthSignature`:** the signed string is the request's parameters rendered as
PHP's `http_build_query` renders them, which differs from Go's `url.Values` in
several characters. That encoding is pinned by known-answer vectors generated
from Payrexx's own documented PHP snippet (`payrexx/auth_test.go`), and the
scheme is verified end to end against a live account by `TestLiveSignatureScheme`.
It still works, but Payrexx's own PHP SDK abandoned it — prefer `AuthAPIKey`.

To layer your own transport (proxy, tracing, retries) beneath authentication,
pass an `*http.Client` with a `Transport` set:

```go
payrexx.NewClient(payrexx.Config{
    Instance:   "example",
    APISecret:  secret,
    HTTPClient: &http.Client{Transport: myTracingTransport},
})
```

`Config.BaseURL` overrides the API base, including the version segment — use it
to pin an older API version, to reach a platform account on its own domain, or to
point at a test server.

## Webhooks

**Webhooks are the source of truth for payment state, not the browser redirect.**
A shopper who closes the tab after paying never hits your success URL; the
webhook still arrives.

```go
event, err := payrexx.ParseWebhook(r, signingKey)
if err != nil {
    // Payrexx retries on non-2xx, so refusing is safe.
    http.Error(w, "rejected", http.StatusUnauthorized)
    return
}
if tx := event.Transaction; tx != nil && tx.IsLive() &&
    payrexx.Value(tx.Status) == payrexx.TransactionConfirmed {
    markPaid(tx.Reference())   // the referenceId you set on the gateway
}
```

`ParseWebhook` verifies the `X-Webhook-Signature` header — lowercase hex
HMAC-SHA256 over the **raw** body, key used as plain UTF-8, compared in constant
time — and never decodes a body that fails. `event.Raw` is the exact bytes that
were verified.

Three things the API will do to a naive handler:

- **Deliveries repeat.** Payrexx retries up to ten times over several days until
  it gets a success, and a handler that timed out after succeeding looks like a
  failure. Key your writes on the transaction `uuid`.
- **Deliveries time out at 20 seconds.** Acknowledge first, do slow work after.
- **`mode` is `TEST` or `LIVE`.** `WebhookTransaction.IsLive()` is the check;
  test payments reaching production logic is the classic mistake.

Webhook payload types (`WebhookTransaction`, `WebhookSubscription`, …) are
hand-written from Payrexx's webhook object tables rather than reused from the
generated entities — it is a different document describing an overlapping but
distinct payload. Every field is optional; read them with `payrexx.Value(...)`.

Set the webhook's content type to **JSON** in the Payrexx merchant account. The
alternative "Normal (PHP-Post)" form encoding is not decoded here, and
`ParseWebhook` says so rather than failing as malformed JSON.

## Regenerating from upstream

The generator version and the spec are both pinned, so regeneration is
reproducible:

```sh
./gen.sh --refresh-spec   # re-fetch the fragments, rebuild openapi.json, regenerate
./gen.sh                  # regenerate from the committed openapi.json
git diff                  # review added/changed fields before committing
```

`gen.sh` prunes the previous generation, regenerates into `payrexx/`, then
`gofmt`s and builds+vets. Hand-written files (`auth.go`, `webhook.go`, the tests)
are protected by `payrexx/.openapi-generator-ignore` and are never overwritten.

### What build_spec.py changes, and why

Payrexx's fragments do not assemble into a usable spec on their own. Each
correction is a documented, named entry in
[`scripts/build_spec.py`](scripts/build_spec.py) rather than a hand edit, and the
script prints what it did on every run — so an upstream change shows up as a diff
in that output, not as a silent behaviour change.

1. **Field-name typo.** `is_list` on one page against `isList` on five; both
   become Go's `IsList` and stop compiling.
2. **Operation naming and grouping.** Slug-derived ids (`CreateANewSubscription`)
   are renamed resource-first (`SubscriptionCreate`), and tags are added — without
   them every operation lands in a single `DefaultAPI`.
3. **Shared entities.** Each page repeats the entity schema inline under `data`,
   so `Gateway` from a create and `Gateway` from a read would be two unrelated Go
   types. They are hoisted into `components/schemas` and merged. The merge is a
   union: no response schema marks anything required, so a field the endpoint did
   not send is simply nil.
4. **Two missing endpoints.** `DELETE /Subscription/{id}/` (cancel) and
   `DELETE /Gateway/{id}/` are absent from the reference but exercised by
   payrexx-php's own examples. Declared explicitly and marked as such in the spec.
5. **Polymorphic fields.** `EcrPayment.slip` is an object on some pages and an
   array on others; typed free-form so both decode.
6. **Alias spellings.** The Bill pages document `paymentLink` and `payment_link`
   for the same field (and ECR adds `payment-status` next to `payment_status`).
   They collide in Go. The camelCase spelling wins, following payrexx-php's
   `Models\Base::fromArray`, which strips underscores before dispatching to a
   setter — i.e. Payrexx itself treats camelCase as the model. **If an endpoint
   really does answer in snake_case, that field decodes as nil.** Only the
   Bill/Invoice entity is documented both ways.
7. **Live corrections.** Where a real response contradicts the published schema,
   the response wins — a model that cannot decode a real answer is worse than a
   slightly less faithful spec. Each entry names the observation it came from.
   See *Live tests*.
8. **One error type.** Failure bodies are documented three ways, including a
   six-branch `oneOf` whose branches differ only in their examples and which pulls
   in a validation dependency. All of them collapse onto `payrexx.Error`.
9. **`instance` removed** from every operation, and appended by the transport.
10. **JSON-only request bodies.** The fragments offer JSON and form-encoded; both
   generate two sets of request types, and Payrexx recommends JSON.

Two upstream quirks are worth knowing about because they are *coerced*, not just
tidied: the Subscription retrieve and list pages type `data` as a bare string,
contradicting their own examples, and are pointed at the entity the other pages
document; and the `Gateway`/`Transaction`/`Invoice` unions carry fields that only
some endpoints return.

## Testing

```sh
go test ./...
```

Offline and credential-free. Covers the signature encoding against PHP-generated
known-answer vectors, the `instance` handling, transport layering, and webhook
verification including forged, missing, wrong-key and wrong-content-type
deliveries.

### Live tests

The offline suite proves this package is self-consistent. It cannot prove the
assembled spec matches the API Payrexx actually serves, and that spec is built
from documentation already known to be wrong in places. A separate suite calls
the real API to close that gap; it skips unless `PAYREXX_LIVE_TEST` is set, so
`go test ./...` stays offline and credential-free. Running it is described in
[CONTRIBUTING.md](CONTRIBUTING.md#running-the-live-tests).

Three defects were found this way and are fixed in the code above:

- `Gateway.preAuthorization` is documented as an integer and returned as a
  **boolean**. Decoding any real gateway failed outright.
- `Gateway.fields` is documented as an object and returned as `[]` when unset
  (PHP's empty-array serialization), an object otherwise. Also a hard decode
  failure.
- The generator marshals a nil body to the literal `null` and sends it with a
  Content-Type on **GET** requests. Payrexx reads a present body as the parameter
  set, so under `AuthSignature` that `null` displaced the query string and every
  read came back `403 The API secret is not correct`. The transport now drops it.

`Gateway.language` and `Gateway.requestId` are returned but documented nowhere;
they are in the model with that noted.

## Continuous verification

This client is generated from a specification assembled out of documentation that
is demonstrably wrong in several places, and corrected by a script whose every
entry is a claim about how Payrexx really behaves. "It compiles" is therefore not
much of a guarantee on its own, and two things run on top of it.

[`ci.yml`](.github/workflows/ci.yml) runs on every push and pull request and
never contacts Payrexx: gofmt, vet, build and `go test -race` on Go 1.23 and
current, plus a check that regenerating from the committed `openapi.json`
reproduces the committed client byte for byte. Because it is fully offline, a
Payrexx outage cannot turn a pull request red.

[`upstream-watch.yml`](.github/workflows/upstream-watch.yml) runs every Monday
against a real Payrexx account. It rebuilds the specification from
developers.payrexx.com and diffs it against the committed one, then runs the live
tests: a gateway created, read back and deleted, the read endpoints, both
authentication schemes, and the shape of an error response. This is the half that
can break without anyone committing anything, and it is what falsifies a
correction in `build_spec.py` once it stops holding.

### How to tell whether this SDK is currently accurate

**A failing watch opens an issue** titled *Upstream watch failed*, carrying the
specification diff or the failing test output, and comments on that issue instead
of filing duplicates. So:

- [Open issues](https://github.com/zweiundeins/payrexx-go-sdk/issues?q=is%3Aissue+is%3Aopen+%22Upstream+watch+failed%22)
  with that title mean Payrexx has changed something and this SDK has not caught
  up yet. Read the diff before trusting the affected models.
- [Actions history](https://github.com/zweiundeins/payrexx-go-sdk/actions/workflows/upstream-watch.yml)
  shows how recently the client was confirmed against the live API. A green run
  means the models decoded real responses that week.

An issue is used rather than an email because GitHub notifies a scheduled
workflow's failure to a single account, the one that last edited its `cron:`
line. The issue tracker is visible to everyone and outlives whoever set it up.

Failure output never contains a response body. The live tests read a real
account, and the failure they exist to catch is a decode failure, so a verbatim
dump would put customer records into a public log at precisely the wrong moment.
Values are replaced by their types, which is also the more useful artefact when
the disagreement is about a field's type. `TestShapeRedactsCustomerData` runs
offline on every push and fails if that stops being true.

Running the watch, and the account it needs, are covered in
[CONTRIBUTING.md](CONTRIBUTING.md).

## License

The generated client mirrors Payrexx's published API definition; the generated
portions follow openapi-generator's Apache-2.0 output licensing. The hand-written
`auth.go`, `webhook.go` and `scripts/build_spec.py` are part of this repository.
This is not an official Payrexx product.
