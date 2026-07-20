# payrexx-go-sdk

A Go client for the [Payrexx](https://payrexx.com) REST API.

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
assembled spec matches the API Payrexx actually serves — and that spec is built
from documentation already known to be wrong in places. That is what the live
tests are for:

```sh
PAYREXX_LIVE_TEST=1 \
PAYREXX_INSTANCE=example \
PAYREXX_API_SECRET=... \
go test -run TestLive -v ./...
```

They skip without `PAYREXX_LIVE_TEST`, so `go test ./...` stays offline, and CI
runs them from [`upstream-watch.yml`](.github/workflows/upstream-watch.yml)
rather than on pull requests. They create gateways — hosted checkout pages, no
money moves, nothing is charged — and delete them again in a `t.Cleanup`. Nothing
else is written.

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

## CI

The split matters, so it is worth stating: **nothing that runs on a pull request
talks to Payrexx.** A Payrexx outage or a documentation edit can never block a
merge.

[`ci.yml`](.github/workflows/ci.yml) — on every push and PR, entirely offline:
gofmt, vet, build and `go test -race` on Go 1.23 and stable, plus a check that
`./gen.sh` against the committed `openapi.json` reproduces the committed client
byte for byte. That last one is what stops a hand-edit to a generated file from
surviving review and then silently vanishing at the next regeneration.

[`upstream-watch.yml`](.github/workflows/upstream-watch.yml) — Mondays 06:00 UTC
and on demand. Rebuilds the spec from developers.payrexx.com and diffs it against
the committed one, then runs the live tests. Both steps run even if the first
fails, and the job goes red if either does. This is the half that can break
without anyone committing anything: every correction in `build_spec.py` is a
claim about an API whose published schema is already wrong in places.

**On failure it opens an issue** titled *Upstream watch failed*, with the spec
diff or the failing test output in the body, and comments on the existing issue
rather than filing a duplicate. That indirection is deliberate — GitHub's own
notification for a scheduled workflow reaches exactly one person ("Notifications
for scheduled workflows are sent to the user who initially created the workflow",
reassigned to whoever last edits the `cron:` line), at whatever address that
account uses, only if their Actions notification preference is on. An issue is
visible to everyone with repository access and survives people leaving.

Two repository secrets are required, or the live step fails loudly rather than
silently skipping: `PAYREXX_INSTANCE` and `PAYREXX_API_SECRET`.

### Which account to point them at

**Use a Payrexx account that holds no customer data — ideally one created for
this.** Payrexx has no sandbox: test mode is a per-PSP toggle inside a real
account, and their own testing guide says a trial subscription is enough, so "a
test instance" means a second real account.

Per weekly run the tests make about eight API calls, create one gateway at
CHF 1.00 and delete it again. Nothing is charged and no money moves. The reads
are reads. That part is harmless on any account.

What is not harmless is the credential and the data behind it:

- **The API secret is full account access.** It can read every transaction,
  invoice and contact, create and delete gateways, and cancel subscriptions.
  Putting an active client's production secret into a public repository's secrets
  gives that blast radius to anyone who can push a workflow change.
- **The tests read real records.** `TestLiveListsDecode` lists transactions,
  subscriptions and invoices. On a production account those carry names, email
  and postal addresses, phone numbers and masked card numbers.

The second point had teeth: the failure these tests exist to catch *is* a decode
failure, so the one moment a response body would be printed is the moment it
holds a populated customer record — and that output goes to a public Actions log
and into the issue this workflow files. The tests therefore never print a body.
`shape()` replaces every value with its type, keeping the keys and the top-level
error envelope, which is also the more useful artefact: a decode failure is a
disagreement about a field's *type*. `TestShapeRedactsCustomerData` runs offline
on every push and fails if that ever stops being true.

That makes a production account survivable rather than advisable. The ranking:

1. **A dedicated Payrexx account** for CI. No customer data to leak, and a
   secret worth nothing if exposed.
2. **The project's own account**, before it has customers. Fine, and it has the
   advantage of testing the instance you will actually ship against.
3. **An unrelated active project's credentials.** Works, but you are handing a
   public repository a key that can cancel that project's subscriptions, to test
   an SDK that project does not use.

The workflow triggers on `schedule` and `workflow_dispatch` only — never
`pull_request` or `pull_request_target` — so a fork cannot run it and cannot
reach the secrets.

### The keepalive commit

Public repositories have their scheduled workflows *"automatically disabled when
no repository activity has occurred in 60 days"* — and an SDK that is working
correctly is exactly a repository with no activity, so this watch would switch
itself off precisely while it was doing its job. Workflow runs do not count as
activity; commits do. So after a clean run, if the last commit is more than 45
days old, the workflow writes the date into `.github/spec-last-verified` and
pushes it. At most one commit every seven weeks, and none at all while anyone is
working on the repo.

It needs `contents: write` and push access to the default branch. If you protect
that branch, either exempt `github-actions[bot]` or drop the step and re-enable
the schedule by hand when GitHub emails about it.

## License

The generated client mirrors Payrexx's published API definition; the generated
portions follow openapi-generator's Apache-2.0 output licensing. The hand-written
`auth.go`, `webhook.go` and `scripts/build_spec.py` are part of this repository.
This is not an official Payrexx product.
