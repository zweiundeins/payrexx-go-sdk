#!/usr/bin/env python3
"""Assemble a single OpenAPI document for the Payrexx REST API.

Payrexx publishes no downloadable spec. What it does publish is a ReadMe.io
documentation site whose every API-reference page carries that one endpoint's
OpenAPI 3.1 fragment in a fenced ``json`` block, reachable by appending ``.md``
to the page URL. ``llms.txt`` indexes those pages. This script fetches them,
merges the fragments into one document, and applies the corrections listed
below so that a code generator can turn it into an idiomatic client.

Every transformation is recorded here rather than applied by hand, so
regeneration stays reproducible and ``git diff`` on ``openapi.json`` stays
meaningful when Payrexx changes something upstream.

Usage::

    scripts/build_spec.py                 # fetch upstream, write openapi.json
    scripts/build_spec.py --from-dir DIR  # rebuild from previously fetched .md files
    scripts/build_spec.py --cache-dir DIR # also save the raw fragments to DIR
"""

from __future__ import annotations

import argparse
import collections
import json
import pathlib
import re
import sys
import urllib.request

LLMS_TXT = "https://developers.payrexx.com/llms.txt"
REFERENCE_RE = re.compile(r"https://developers\.payrexx\.com/reference/[a-z0-9-]+\.md")
JSON_BLOCK_RE = re.compile(r"```json\n(.*?)\n```", re.S)

# The API version the assembled document targets. Fragments all declare the same
# server; this is asserted at merge time rather than assumed.
EXPECTED_SERVER = "https://api.payrexx.com/v1.16"

# --- corrections ------------------------------------------------------------

# 1. Payrexx spells the same pagination field ``isList`` on five pages and
#    ``is_list`` on one. Both land on a Go field named IsList, which does not
#    compile. The majority spelling wins.
FIELD_TYPOS = {'"is_list"': '"isList"'}

# 2. Operation naming. The fragments derive operationIds from page slugs
#    ("create-a-new-subscription"), which generate methods like
#    CreateANewSubscription. This table renames them resource-first so the
#    generated services read as SubscriptionAPI.SubscriptionCreate(...), and
#    assigns the tag that groups operations into services in the first place --
#    without tags every operation lands in one DefaultAPI god-object.
#
#    An upstream endpoint missing from this table is not an error: it keeps its
#    slug-derived name, lands in DefaultAPI, and is reported at the end of the
#    run so the omission is visible in the regeneration output.
OPERATIONS = {
    ("/Bill/", "get"): ("Bill", "BillList"),
    ("/Bill/", "post"): ("Bill", "BillCreate"),
    ("/Bill/{id}/", "get"): ("Bill", "BillRetrieve"),
    ("/Bill/{id}/", "patch"): ("Bill", "BillUpdate"),
    ("/Bill/{id}/", "delete"): ("Bill", "BillDelete"),
    ("/Gateway/", "post"): ("Gateway", "GatewayCreate"),
    ("/Gateway/{id}/", "get"): ("Gateway", "GatewayRetrieve"),
    ("/Gateway/{id}/", "delete"): ("Gateway", "GatewayDelete"),
    ("/Invoice/", "post"): ("Invoice", "InvoiceCreatePaylink"),
    ("/PaymentProvider/", "get"): ("PaymentProvider", "PaymentProviderList"),
    ("/Payout/", "get"): ("Payout", "PayoutList"),
    ("/Payout/{uuid}/", "get"): ("Payout", "PayoutRetrieve"),
    ("/Payout/{uuid}/details", "get"): ("Payout", "PayoutRetrieveWithDetails"),
    ("/Subscription/", "get"): ("Subscription", "SubscriptionList"),
    ("/Subscription/", "post"): ("Subscription", "SubscriptionCreate"),
    ("/Subscription/{id}/", "get"): ("Subscription", "SubscriptionRetrieve"),
    ("/Subscription/{id}/", "put"): ("Subscription", "SubscriptionUpdate"),
    ("/Subscription/{id}/", "delete"): ("Subscription", "SubscriptionCancel"),
    ("/Transaction/", "get"): ("Transaction", "TransactionList"),
    ("/Transaction/{id}/", "get"): ("Transaction", "TransactionRetrieve"),
    ("/Transaction/{id}/", "put"): ("Transaction", "TransactionUpdateContactDetails"),
    ("/Transaction/{id}/cancel", "patch"): ("Transaction", "TransactionCancel"),
    ("/Transaction/{id}/updateTokenization", "patch"): ("Transaction", "TransactionUpdateTokenization"),
    ("/ecr/{serialNumber}/pair", "get"): ("Ecr", "EcrPairingRetrieve"),
    ("/ecr/{serialNumber}/pair", "post"): ("Ecr", "EcrPair"),
    ("/ecr/{serialNumber}/pair", "delete"): ("Ecr", "EcrUnpair"),
    ("/ecr/{serialNumber}/payment", "post"): ("Ecr", "EcrPaymentInitiate"),
    ("/ecr/{serialNumber}/payment/{paymentId}", "get"): ("Ecr", "EcrPaymentRetrieve"),
    ("/ecr/{serialNumber}/payment/{paymentId}/cancel", "post"): ("Ecr", "EcrPaymentCancel"),
    ("/ecr/{serialNumber}/payment/{paymentId}/void", "post"): ("Ecr", "EcrPaymentVoid"),
    ("/ecr/{serialNumber}/paymentMethods", "get"): ("Ecr", "EcrPaymentMethodList"),
}

# 3. Response entity names. Every Payrexx response is an envelope --
#    ``{"status": "success", "data": [ ... ]}`` -- and the entity schema sits
#    inline under ``data.items``, repeated per page. Left alone, the generator
#    emits GatewayCreate200ResponseDataInner and GatewayRetrieve200ResponseDataInner
#    as two unrelated types, so a gateway read cannot be passed where a gateway
#    create result is expected. This table hoists those inline schemas into
#    named components, merging the occurrences (see merge_schema).
ENTITIES = {
    ("/Bill/", "get"): "Invoice",
    ("/Bill/", "post"): "Invoice",
    ("/Bill/{id}/", "get"): "Invoice",
    ("/Bill/{id}/", "patch"): "Invoice",
    ("/Bill/{id}/", "delete"): "InvoiceDeletion",
    ("/Gateway/", "post"): "Gateway",
    ("/Gateway/{id}/", "get"): "Gateway",
    ("/Invoice/", "post"): "Paylink",
    ("/PaymentProvider/", "get"): "PaymentProvider",
    ("/Payout/", "get"): "Payout",
    ("/Payout/{uuid}/", "get"): "Payout",
    ("/Payout/{uuid}/details", "get"): "Payout",
    ("/Subscription/", "get"): "Subscription",
    ("/Subscription/", "post"): "Subscription",
    ("/Subscription/{id}/", "get"): "Subscription",
    ("/Subscription/{id}/", "put"): "Subscription",
    ("/Transaction/", "get"): "Transaction",
    ("/Transaction/{id}/", "get"): "Transaction",
    ("/Transaction/{id}/", "put"): "Transaction",
    ("/Transaction/{id}/cancel", "patch"): "Transaction",
    ("/Transaction/{id}/updateTokenization", "patch"): "Transaction",
    ("/ecr/{serialNumber}/pair", "get"): "EcrPairing",
    ("/ecr/{serialNumber}/pair", "post"): "EcrPairing",
    ("/ecr/{serialNumber}/payment", "post"): "EcrPayment",
    ("/ecr/{serialNumber}/payment/{paymentId}", "get"): "EcrPayment",
    ("/ecr/{serialNumber}/payment/{paymentId}/cancel", "post"): "EcrPayment",
    ("/ecr/{serialNumber}/payment/{paymentId}/void", "post"): "EcrPayment",
    ("/ecr/{serialNumber}/paymentMethods", "get"): "EcrPaymentMethod",
}

# 4. Endpoints the reference pages omit. Both are exercised by Payrexx's own PHP
#    SDK (examples/SubscriptionCancel.php, examples/GatewayDelete.php), which
#    maps its `cancel`/`delete` verbs onto HTTP DELETE against the entity URL.
#    Cancelling a subscription is not an optional part of a payment integration,
#    so the two are declared here rather than left as a hole callers have to work
#    around with a raw http.Client.
SDK_SOURCED = {
    "/Subscription/{id}/": {
        "delete": {
            "summary": "Cancel a Subscription",
            "description": (
                "Cancels a subscription. Not present in Payrexx's published API reference; "
                "declared here after payrexx-php, whose `cancel` verb issues DELETE against "
                "this URL (see examples/SubscriptionCancel.php)."
            ),
            "parameters": [
                {
                    "name": "id",
                    "in": "path",
                    "required": True,
                    "schema": {"type": "integer"},
                    "description": "The subscription id.",
                }
            ],
            "responses": {
                "200": {
                    "description": "OK",
                    "content": {
                        "application/json": {
                            "schema": {
                                "type": "object",
                                "properties": {
                                    "status": {"type": "string", "example": "success"},
                                    "data": {
                                        "type": "array",
                                        "items": {"$ref": "#/components/schemas/Subscription"},
                                    },
                                },
                            }
                        }
                    },
                }
            },
        }
    },
    "/Gateway/{id}/": {
        "delete": {
            "summary": "Delete a Gateway",
            "description": (
                "Removes a gateway that has not been used yet. Not present in Payrexx's "
                "published API reference; declared here after payrexx-php, whose `delete` "
                "verb issues DELETE against this URL (see examples/GatewayDelete.php)."
            ),
            "parameters": [
                {
                    "name": "id",
                    "in": "path",
                    "required": True,
                    "schema": {"type": "integer"},
                    "description": "The gateway id.",
                }
            ],
            "responses": {
                "200": {
                    "description": "OK",
                    "content": {
                        "application/json": {
                            "schema": {
                                "type": "object",
                                "properties": {
                                    "status": {"type": "string", "example": "success"},
                                    "data": {
                                        "type": "array",
                                        "items": {"$ref": "#/components/schemas/Gateway"},
                                    },
                                },
                            }
                        }
                    },
                }
            },
        }
    },
}


# 5. Fields the pages type inconsistently, where neither spelling is safe to
#    pick. ``EcrPayment.slip`` is an object on the initiate/cancel/void pages and
#    an array on the payment-methods page; a Go struct typed either way fails to
#    decode the other. Typing them as a free-form schema yields `interface{}`,
#    which decodes both. Keyed by "<entity>.<dotted property path>".
POLYMORPHIC = {
    "EcrPayment.slip": "object on the payment pages, array on the paymentMethods page",
}

# 6. camelCase vs snake_case twins. The Bill pages document the same entity twice
#    over: `paymentLink` on the create/update pages, `payment_link` on the
#    list/retrieve pages. Unioned as-is they are two properties that both
#    PascalCase to Go's PaymentLink, which does not compile, and x-go-name is not
#    honoured by the Go generator. Payrexx's own PHP SDK settles which spelling is
#    canonical: Models\Base::fromArray strips underscores and title-cases before
#    dispatching to a setter, so `payment_link` and `paymentLink` both drive
#    setPaymentLink -- camelCase is the model, snake_case an accepted alias. This
#    keeps the camelCase key and folds the twin into it.
#
#    Consequence worth knowing: if an endpoint really does answer in snake_case,
#    that field decodes as nil rather than failing. Only the Bill/Invoice entity
#    is documented both ways.
#    The ECR pages add a third spelling of the same idea, `payment-status`
#    alongside `payment_status`, so the fold keys on "same letters ignoring
#    separators" rather than on snake_case specifically.


def fold_alias_spellings(schema, where: str, folded: list[str]) -> None:
    """Merge properties that differ only in separators, preferring camelCase."""
    if not isinstance(schema, dict):
        return
    props = schema.get("properties")
    if isinstance(props, dict):
        groups: dict[str, list[str]] = collections.defaultdict(list)
        for key in props:
            groups[key.replace("_", "").replace("-", "").lower()].append(key)
        for spellings in groups.values():
            if len(spellings) == 1:
                continue
            # Prefer the separator-free spelling; failing that, the first listed.
            canonical = next((s for s in spellings if "_" not in s and "-" not in s), spellings[0])
            for alias in spellings:
                if alias == canonical:
                    continue
                props[canonical] = merge_schema(
                    props[canonical], props.pop(alias), f"{where}.{canonical}", []
                )
                folded.append(f"{where}.{alias} -> {canonical}")
        for name, sub in props.items():
            fold_alias_spellings(sub, f"{where}.{name}", folded)
    if isinstance(schema.get("items"), dict):
        fold_alias_spellings(schema["items"], f"{where}[]", folded)


# 7. A single error type. The pages describe failure bodies in three ways: an
#    empty `{"type": "object", "properties": {}}` (most of them, carrying no
#    information at all), a `{status, message}` envelope, and a `{error: {type,
#    param, message}}` envelope on the newer Bill and ECR endpoints -- the last
#    spelled as a six-branch `oneOf` whose branches differ only in their examples,
#    which the generator turns into a discriminated union pulling in a validation
#    dependency. Collapsing all of them onto one schema gives callers a single
#    type to assert on, and keeps the generated package dependency-free.
ERROR_SCHEMA = {
    "type": "object",
    "description": (
        "Failure body. Payrexx documents two shapes and returns the one that "
        "matches the endpoint family, so only one group of fields is populated: "
        "`status`/`message`/`reason` on the classic endpoints, `error` on the Bill "
        "and ECR endpoints. Assembled by scripts/build_spec.py; the reference "
        "pages describe these per endpoint."
    ),
    "properties": {
        "status": {"type": "string", "example": "error"},
        "message": {"type": "string", "example": "Invalid value passed for parameter 'currency'"},
        "reason": {"type": "string"},
        "error": {
            "type": "object",
            "properties": {
                "type": {"type": "string", "example": "invalid_value"},
                "param": {"type": "string", "example": "currency"},
                "message": {
                    "type": "string",
                    "example": "Invalid value passed for parameter 'currency'",
                },
            },
        },
    },
}


def fetch(url: str) -> str:
    req = urllib.request.Request(url, headers={"User-Agent": "payrexx-go-sdk/build_spec"})
    with urllib.request.urlopen(req, timeout=60) as resp:
        return resp.read().decode("utf-8")


def load_fragments(from_dir: str | None, cache_dir: str | None) -> dict[str, str]:
    """Return {name: markdown} for every API-reference page."""
    if from_dir:
        paths = sorted(pathlib.Path(from_dir).glob("*.md"))
        if not paths:
            sys.exit(f"!! no .md files in {from_dir}")
        return {p.name: p.read_text(encoding="utf-8") for p in paths}

    print(f">> index: {LLMS_TXT}")
    urls = sorted(set(REFERENCE_RE.findall(fetch(LLMS_TXT))))
    if not urls:
        sys.exit("!! llms.txt listed no /reference/ pages -- upstream layout changed")
    print(f">> fetching {len(urls)} reference pages")
    out = {}
    for url in urls:
        name = url.rsplit("/", 1)[1]
        out[name] = fetch(url)
    if cache_dir:
        d = pathlib.Path(cache_dir)
        d.mkdir(parents=True, exist_ok=True)
        for name, text in out.items():
            (d / name).write_text(text, encoding="utf-8")
        print(f">> cached raw fragments in {cache_dir}/")
    return out


def merge_schema(a, b, where: str, conflicts: list[str]):
    """Union two occurrences of the same entity schema.

    The same entity is documented separately on every page that returns it, and
    the pages disagree: ``Gateway`` carries ``reservation`` on the create page and
    ``applicationFee`` on the retrieve page, ``Invoice`` is camelCase on two pages
    and snake_case on two others. None of the response schemas mark any property
    required, so a union is decode-safe -- a property the endpoint did not send
    simply stays nil -- and it is the only representation under which one Go type
    can carry the entity across calls. Property type disagreements are collected
    and reported rather than silently resolved.
    """
    if a == b:
        return a
    if not isinstance(a, dict) or not isinstance(b, dict):
        conflicts.append(f"{where}: non-object schema mismatch, keeping first")
        return a

    out = dict(a)
    for key, bv in b.items():
        if key not in out:
            out[key] = bv
        elif key == "properties":
            merged = dict(out[key])
            for prop, bprop in bv.items():
                merged[prop] = (
                    merge_schema(merged[prop], bprop, f"{where}.{prop}", conflicts)
                    if prop in merged
                    else bprop
                )
            out[key] = merged
        elif key == "required":
            out[key] = sorted(set(out[key]) | set(bv))
        elif key == "items":
            out[key] = merge_schema(out[key], bv, f"{where}[]", conflicts)
        elif out[key] != bv:
            if key == "type":
                conflicts.append(f"{where}: type {out[key]!r} vs {bv!r}, keeping {out[key]!r}")
            # description/example/default differences are noise; first wins.
    return out


def build(fragments: dict[str, str]) -> dict:
    paths: dict[str, dict] = collections.defaultdict(dict)
    entities: dict[str, dict] = {}
    conflicts: list[str] = []
    servers: set[str] = set()
    skipped: list[str] = []

    for name, text in sorted(fragments.items()):
        block = JSON_BLOCK_RE.search(text)
        if not block:
            skipped.append(name)  # guide pages (changelogs, "Getting started")
            continue
        doc = json.loads(FIELD_TYPOS_RE.sub(lambda m: FIELD_TYPOS[m.group(0)], block.group(1)))
        for server in doc.get("servers", []):
            servers.add(server["url"])
        for path, ops in doc.get("paths", {}).items():
            for method, op in ops.items():
                if method in paths[path]:
                    conflicts.append(f"duplicate operation {method.upper()} {path} (from {name})")
                    continue
                paths[path][method] = prepare(path, method, op, entities, conflicts)

    if servers != {EXPECTED_SERVER}:
        sys.exit(f"!! unexpected servers {sorted(servers)}, expected [{EXPECTED_SERVER}]")

    for path, ops in SDK_SOURCED.items():
        for method, op in ops.items():
            if method in paths[path]:
                print(f">> upstream now documents {method.upper()} {path}; drop it from SDK_SOURCED")
                continue
            tag, op_id = OPERATIONS[(path, method)]
            paths[path][method] = {"tags": [tag], "operationId": op_id, **op}

    folded: list[str] = []
    for name, schema in entities.items():
        fold_alias_spellings(schema, name, folded)
    if folded:
        print(f">> folded {len(folded)} alias spellings into a canonical property:")
        for entry in sorted(folded):
            print(f"     {entry}")

    for dotted, reason in POLYMORPHIC.items():
        entity, *trail = dotted.split(".")
        node = entities.get(entity)
        for step in trail:
            node = (node or {}).get("properties", {}).get(step)
        if node is None:
            conflicts.append(f"POLYMORPHIC entry {dotted} matched nothing -- upstream changed")
            continue
        node.clear()
        node["description"] = f"Free-form: {reason}."

    untagged = [
        f"{m.upper()} {p}" for p, ops in paths.items() for m, o in ops.items() if not o.get("tags")
    ]
    if untagged:
        print(">> NOT IN OPERATIONS TABLE (kept slug-derived naming, grouped under DefaultAPI):")
        for entry in sorted(untagged):
            print(f"     {entry}")
    # Every hoisted reference has to resolve, or the generator emits a client
    # that will not compile -- cheaper to catch here than three minutes later.
    referenced = set(re.findall(r'"#/components/schemas/([A-Za-z0-9_]+)"', json.dumps(paths)))
    missing = sorted(referenced - set(entities) - {"Error"})
    if missing:
        sys.exit(f"!! referenced but never defined: {missing}; check the ENTITIES table")

    conflicts = [c for c in conflicts if c.split(":")[0] not in POLYMORPHIC]
    if conflicts:
        print(">> schema quirks resolved:")
        for entry in conflicts:
            print(f"     {entry}")
    if skipped:
        print(f">> {len(skipped)} pages carry no OpenAPI block (guides/changelogs): "
              f"{', '.join(sorted(skipped))}")

    return {
        "openapi": "3.0.3",
        "info": {
            "title": "Payrexx REST API",
            "version": EXPECTED_SERVER.rsplit("/v", 1)[1],
            "description": (
                "Assembled by scripts/build_spec.py from the per-endpoint OpenAPI fragments "
                "published on developers.payrexx.com. Not an official Payrexx artefact."
            ),
        },
        "servers": [{"url": EXPECTED_SERVER, "description": "Payrexx REST API"}],
        "tags": [{"name": t} for t in sorted({t for t, _ in OPERATIONS.values()})],
        "paths": {p: dict(sorted(ops.items())) for p, ops in sorted(paths.items())},
        "components": {"schemas": dict(sorted({**entities, "Error": ERROR_SCHEMA}.items()))},
    }


FIELD_TYPOS_RE = re.compile("|".join(re.escape(k) for k in FIELD_TYPOS))


def prepare(path: str, method: str, op: dict, entities: dict, conflicts: list[str]) -> dict:
    op = json.loads(json.dumps(op))  # the fragment dicts are reused; work on a copy

    # `instance` identifies the Payrexx account and is the same for every call a
    # client makes. Carrying it as a required argument on all 31 operations would
    # be noise, and the API-signature scheme has to exclude it from the signed
    # query string anyway -- so it is dropped here and appended by the transport
    # in payrexx/auth.go, which is the only place that knows both.
    op["parameters"] = [p for p in op.get("parameters", []) if p.get("name") != "instance"]
    if not op["parameters"]:
        del op["parameters"]

    # The fragments declare `security: [{}]` -- "no security scheme" -- because
    # ReadMe has nowhere to express Payrexx's header/signature auth. Authentication
    # is a transport concern here (payrexx/auth.go); leaving the empty declaration
    # in only makes the generator emit dead scaffolding.
    op.pop("security", None)

    # Payrexx accepts the same body as JSON or as form-urlencoded and recommends
    # JSON; its own PHP SDK sends JSON. Keeping both makes the generator pick one
    # arbitrarily and emit a second set of request types for the other.
    body = op.get("requestBody", {}).get("content", {})
    if "application/json" in body and "application/x-www-form-urlencoded" in body:
        del body["application/x-www-form-urlencoded"]

    if (path, method) in OPERATIONS:
        tag, op_id = OPERATIONS[(path, method)]
        op["tags"] = [tag]
        op["operationId"] = op_id

    for code, response in op.get("responses", {}).items():
        if not code.startswith("2") and "application/json" in response.get("content", {}):
            response["content"]["application/json"]["schema"] = {
                "$ref": "#/components/schemas/Error"
            }

    entity = ENTITIES.get((path, method))
    if entity:
        schema = (
            op.get("responses", {})
            .get("200", {})
            .get("content", {})
            .get("application/json", {})
            .get("schema", {})
        )
        props = schema.get("properties") or {}
        data = props.get("data")
        inline = None
        if isinstance(data, dict):
            inline = data.get("items") if data.get("type") == "array" else data
        if isinstance(inline, dict) and inline.get("properties"):
            entities[entity] = (
                merge_schema(entities[entity], inline, entity, conflicts)
                if entity in entities
                else inline
            )
            ref = {"$ref": f"#/components/schemas/{entity}"}
            if data.get("type") == "array":
                data["items"] = ref
            else:
                props["data"] = ref
        elif data is not None:
            # The page describes no entity under `data` -- the Subscription
            # retrieve and list pages type it as a bare string, which their own
            # examples contradict. Payrexx's envelope is always
            # {"status": ..., "data": [entity, ...]}; payrexx-php iterates
            # $response['body']['data'] unconditionally. Point it at the entity
            # the other pages document rather than generate an accessor that
            # cannot decode the response.
            conflicts.append(
                f"{method.upper()} {path}: 200 `data` documented as "
                f"{json.dumps(data)[:60]}; coerced to an array of {entity}"
            )
            props["data"] = {
                "type": "array",
                "items": {"$ref": f"#/components/schemas/{entity}"},
            }
    return op


def main() -> int:
    ap = argparse.ArgumentParser(description=__doc__)
    ap.add_argument("-o", "--out", default="openapi.json")
    ap.add_argument("--from-dir", help="rebuild from previously fetched .md fragments")
    ap.add_argument("--cache-dir", help="save the raw fetched fragments here")
    args = ap.parse_args()

    spec = build(load_fragments(args.from_dir, args.cache_dir))
    text = json.dumps(spec, indent=2, ensure_ascii=False) + "\n"
    pathlib.Path(args.out).write_text(text, encoding="utf-8")

    ops = sum(len(v) for v in spec["paths"].values())
    print(
        f">> wrote {args.out}: {len(spec['paths'])} paths, {ops} operations, "
        f"{len(spec['components']['schemas'])} entities, {len(text)} bytes"
    )
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
