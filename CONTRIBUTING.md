# Contributing

## Regenerating after a Payrexx change

See [Regenerating from upstream](README.md#regenerating-from-upstream). In short:
`./gen.sh --refresh-spec`, then read `git diff openapi.json` before committing.
An added field is routine. A changed type is not, and usually means a correction
in `scripts/build_spec.py` needs revisiting.

`ci.yml` checks that regenerating from the committed `openapi.json` reproduces
the committed client byte for byte, so a hand-edit to a generated file fails the
build rather than silently disappearing at the next `./gen.sh`.

## Running the live tests

```sh
PAYREXX_LIVE_TEST=1 \
PAYREXX_INSTANCE=example \
PAYREXX_API_SECRET=... \
go test -run TestLive -v ./...
```

Without `PAYREXX_LIVE_TEST` they skip, so `go test ./...` stays offline and needs
no credentials.

Per run they make about eight API calls and create one gateway at CHF 1.00, read
it back, and delete it in a `t.Cleanup`. A gateway is a hosted checkout page:
nothing is charged and no money moves. Nothing else is written.

## Which Payrexx account to use

**One that holds no customer data, ideally created for this purpose.** Payrexx
has no sandbox. Test mode is a per-PSP toggle inside a real account, and Payrexx's
testing guide says a trial subscription is enough, so "a test instance" means a
second real account.

The API calls are harmless anywhere. The credential and the data behind it are
not:

- **The API secret is full account access.** It can read every transaction,
  invoice and contact, create and delete gateways, and cancel subscriptions.
  Putting an active client's production secret into a public repository's secrets
  hands that reach to anyone who can push a workflow change.
- **The tests read real records.** `TestLiveListsDecode` lists transactions,
  subscriptions and invoices. On a production account those carry names, email
  and postal addresses, phone numbers and masked card numbers.

The second point has teeth, because the failure these tests exist to catch *is* a
decode failure: the one moment a response body would be printed is the moment it
holds a populated customer record, and that output reaches a public Actions log
and the issue the watch files. So the tests never print a body. `shape()`
replaces every value with its type, keeping the keys and the top-level error
envelope, and `TestShapeRedactsCustomerData` fails the build if that ever stops
being true.

That makes a production account survivable rather than advisable:

1. **A dedicated Payrexx account.** No customer data to leak, and a secret worth
   nothing if it leaks.
2. **The project's own account, before it has customers.** Fine, with the
   advantage of testing the instance you actually ship against.
3. **An unrelated active project's credentials.** Works, but it hands a public
   repository a key that can cancel that project's subscriptions, to test an SDK
   that project does not use.

## CI secrets

`upstream-watch.yml` needs two repository secrets. Without them the live step
fails loudly rather than skipping quietly, because a watchdog that silently does
nothing is worse than none:

| Secret | Value |
|---|---|
| `PAYREXX_INSTANCE` | the instance name, e.g. `example` for `example.payrexx.com` |
| `PAYREXX_API_SECRET` | that instance's API secret |

```sh
gh secret set PAYREXX_INSTANCE   --repo zweiundeins/payrexx-go-sdk
gh secret set PAYREXX_API_SECRET --repo zweiundeins/payrexx-go-sdk
```

The workflow triggers on `schedule` and `workflow_dispatch` only, never
`pull_request` or `pull_request_target`, so a fork cannot run it and cannot reach
the secrets.

Run it on demand with `gh workflow run upstream-watch.yml`. Note that a
schedule-only workflow is not registered by the push that creates a repository;
it appears after a later push that touches the file.

## The keepalive commit

Public repositories have their scheduled workflows *"automatically disabled when
no repository activity has occurred in 60 days"*, and an SDK that is working
correctly is exactly a repository with no activity, so this watch would switch
itself off precisely while doing its job. Workflow runs do not count as activity;
commits do. After a clean run, if the last commit is more than 45 days old, the
workflow writes the date into `.github/spec-last-verified` and pushes it. That is
at most one commit every seven weeks, and none while anyone is working on the
repository.

It needs `contents: write` and push access to the default branch. If that branch
is protected, either exempt `github-actions[bot]` or drop the step and re-enable
the schedule by hand when GitHub sends the disablement notice.

## Conventions

- Generated files are never hand-edited. `payrexx/auth.go`, `payrexx/webhook.go`
  and the tests are the hand-written surface, protected by
  `payrexx/.openapi-generator-ignore`.
- Corrections to Payrexx's specification belong in `scripts/build_spec.py` as a
  named entry with the reason, not as an edit to `openapi.json`. The script
  prints what it did on every run, so an upstream change shows up as a diff in
  that output.
- The generated package has no third-party dependencies. Keep it that way; a
  schema construct that makes openapi-generator reach for a validation library is
  a reason to simplify the schema in `build_spec.py`.
