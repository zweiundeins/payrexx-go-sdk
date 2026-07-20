#!/usr/bin/env bash
#
# Regenerate the Payrexx Go SDK.
#
# Payrexx ships no OpenAPI document; scripts/build_spec.py assembles one from the
# per-endpoint fragments on developers.payrexx.com and applies the corrections
# documented in that file. The result is committed as openapi.json, so a plain
# ./gen.sh is reproducible offline and `git diff openapi.json` shows exactly what
# upstream changed when it is refreshed.
#
# The generated client lives in the payrexx/ subpackage; the module root keeps
# only project files (go.mod, README, this script, the spec, CI). The generator
# version is pinned in openapitools.json. Hand-written files (payrexx/auth.go,
# payrexx/webhook.go, payrexx/*_test.go) are protected by
# payrexx/.openapi-generator-ignore and are never touched.
#
# Usage:
#   ./gen.sh                 regenerate from the committed openapi.json
#   ./gen.sh --refresh-spec  first rebuild openapi.json from upstream, then regenerate
#
# After running, review `git diff` before committing.
set -euo pipefail
cd "$(dirname "$0")"

SPEC_FILE="openapi.json"
OUT_DIR="payrexx" # the generated package lives in ./payrexx

if [[ "${1:-}" == "--refresh-spec" ]]; then
  echo ">> rebuilding $SPEC_FILE from upstream"
  python3 scripts/build_spec.py -o "$SPEC_FILE"
fi

if [[ ! -f "$SPEC_FILE" ]]; then
  echo "!! $SPEC_FILE not found; run with --refresh-spec to build it" >&2
  exit 1
fi

echo ">> spec: $SPEC_FILE ($(wc -c < "$SPEC_FILE") bytes)"
if command -v sha256sum >/dev/null 2>&1; then
  echo ">> sha256: $(sha256sum "$SPEC_FILE" | cut -d' ' -f1)"
fi

# Drop the previous generation first. The generator overwrites but never deletes,
# so a model that upstream removed (or that a build_spec.py change stopped
# emitting) would linger and break the build. .openapi-generator/FILES is the
# manifest of what the last run wrote; anything in it is safe to remove, and
# .openapi-generator-ignore keeps hand-written files out of that manifest.
MANIFEST="$OUT_DIR/.openapi-generator/FILES"
if [[ -f "$MANIFEST" ]]; then
  echo ">> pruning $(wc -l < "$MANIFEST") files from the previous generation"
  while IFS= read -r f; do
    [[ -n "$f" ]] && rm -f "$OUT_DIR/$f"
  done < "$MANIFEST"
fi

# --skip-validate-spec: the assembled document inherits Payrexx's own quirks
# (empty inline schemas, missing operation descriptions) that the validator
# rejects and the generator handles fine.
# go.mod is hand-maintained at the module root (withGoMod=false), so the
# generator only writes package sources into $OUT_DIR.
npx --yes @openapitools/openapi-generator-cli generate \
  -i "$SPEC_FILE" \
  -g go \
  -o "$OUT_DIR" \
  --skip-validate-spec \
  --additional-properties=packageName=payrexx,withGoMod=false,isGoSubmodule=false,enumClassPrefix=true,hideGenerationTimestamp=true

# The generator writes scaffolding this repo does not use: stub tests that pull
# in testify (real tests are hand-written), a push script, a Travis config, and a
# package README that would shadow the real one at the module root. They are
# listed in .openapi-generator-ignore, but that only stops overwrites, not first
# creation on a clean checkout.
rm -rf "$OUT_DIR/test" "$OUT_DIR/git_push.sh" "$OUT_DIR/.travis.yml" \
  "$OUT_DIR/README.md" "$OUT_DIR/.gitignore"

echo ">> gofmt"
gofmt -w .

echo ">> go build"
go build ./...

echo ">> go vet"
go vet ./...

echo ">> done. Review 'git diff' before committing."
