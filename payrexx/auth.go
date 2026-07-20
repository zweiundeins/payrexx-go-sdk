/*
Payrexx REST API — authentication layer.

This file is hand-written and is NOT produced by the OpenAPI generator (it is
listed in .openapi-generator-ignore so `gen.sh` never overwrites it).

Payrexx has no security scheme in its published API reference, so the generated
client ships without auth. Two schemes are live, both selectable via Config.Auth:

  - X-API-KEY (default) — the API secret in a header. What Payrexx recommends and
    what its own PHP SDK sends.
  - API signature (legacy) — an HMAC-SHA256 over the request's parameters, passed
    as an extra ApiSignature parameter.

Both are applied as an http.RoundTripper, so they cover every generated call.
The transport also appends the `instance` query parameter that every Payrexx
endpoint requires: it identifies the account, is constant for the lifetime of a
client, and has to be excluded from the signed parameter set — so the transport,
which knows both the instance and the scheme, is the one place that can get it
right. scripts/build_spec.py strips it from the generated operation signatures
for the same reason.
*/

package payrexx

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// DefaultBaseURL is the production Payrexx REST API base, including the API
// version the generated client was built against.
const DefaultBaseURL = "https://api.payrexx.com/v1.16"

// AuthScheme selects how requests are authenticated.
type AuthScheme int

const (
	// AuthAPIKey sends the API secret in the X-API-KEY header. This is the zero
	// value, is what Payrexx recommends, and is what payrexx-php sends.
	AuthAPIKey AuthScheme = iota
	// AuthSignature signs each request's parameters and passes the result as an
	// additional ApiSignature parameter, leaving the secret off the wire. This is
	// Payrexx's older scheme; prefer AuthAPIKey unless you have a reason not to.
	AuthSignature
)

// Config holds the credentials and options for a Payrexx API client.
type Config struct {
	// Instance is the Payrexx instance name — the subdomain you reach your
	// account on. For example.payrexx.com that is "example".
	Instance string
	// APISecret is the instance's API secret, from the Payrexx admin panel under
	// the API settings. Server-side only; it must never reach a browser.
	APISecret string
	// Auth selects the authentication scheme. The zero value (AuthAPIKey) is
	// recommended.
	Auth AuthScheme
	// BaseURL overrides the API base, including the version segment. Empty uses
	// DefaultBaseURL. Set it to pin an older API version, to reach a platform
	// account on its own domain, or to a test server URL in tests.
	BaseURL string
	// HTTPClient is the underlying client. If nil, a new one is used. Its
	// Transport is wrapped, so set Transport here to layer your own RoundTripper
	// (proxy, tracing, retries) beneath the authentication.
	HTTPClient *http.Client
	// UserAgent overrides the default User-Agent sent by the generated client.
	UserAgent string
}

// NewClient builds an *APIClient that authenticates every request and carries
// the instance name for you.
func NewClient(cfg Config) (*APIClient, error) {
	if cfg.Instance == "" {
		return nil, errors.New("payrexx: Config.Instance is required")
	}
	if cfg.APISecret == "" {
		return nil, errors.New("payrexx: Config.APISecret is required")
	}

	base := cfg.BaseURL
	if base == "" {
		base = DefaultBaseURL
	}

	httpClient := cfg.HTTPClient
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	next := httpClient.Transport
	if next == nil {
		next = http.DefaultTransport
	}
	httpClient.Transport = &authTransport{
		instance: cfg.Instance,
		secret:   cfg.APISecret,
		scheme:   cfg.Auth,
		next:     next,
	}

	gc := NewConfiguration()
	gc.Servers = ServerConfigurations{{URL: base, Description: "Payrexx REST API"}}
	gc.HTTPClient = httpClient
	if cfg.UserAgent != "" {
		gc.UserAgent = cfg.UserAgent
	}
	return NewAPIClient(gc), nil
}

// authTransport authenticates each outgoing request. It satisfies
// http.RoundTripper.
type authTransport struct {
	instance string
	secret   string
	scheme   AuthScheme
	next     http.RoundTripper
}

func (t *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// The RoundTripper contract forbids mutating the passed request, so work on
	// a clone.
	r := req.Clone(req.Context())

	if t.scheme == AuthSignature {
		if err := t.sign(r); err != nil {
			return nil, err
		}
	} else {
		r.Header.Set("X-API-KEY", t.secret)
	}

	// `instance` is added after signing: Payrexx excludes it from the signed
	// parameter set. Appended textually rather than through url.Values so the
	// order of the parameters the generated client built — which is the order the
	// signature was computed over — survives.
	if r.URL.RawQuery != "" {
		r.URL.RawQuery += "&"
	}
	r.URL.RawQuery += "instance=" + phpEscape(t.instance)

	return t.next.RoundTrip(r)
}

// sign attaches an ApiSignature parameter computed over the request's other
// parameters.
//
// Payrexx defines the signature as base64(HMAC-SHA256(query_string, secret)),
// where query_string is the parameters rendered as PHP's http_build_query would
// render them — nested values as name[key], lists as name[0], spaces as "+".
// The parameters live in the query string on GET and DELETE and in the JSON body
// everywhere else, which is where the two branches below come from.
func (t *authTransport) sign(r *http.Request) error {
	switch r.Method {
	case http.MethodGet, http.MethodDelete:
		if r.URL.RawQuery == "" {
			r.URL.RawQuery = "ApiSignature=" + phpEscape(t.signature(""))
			return nil
		}
		r.URL.RawQuery += "&ApiSignature=" + phpEscape(t.signature(phpQueryFromURLQuery(r.URL.RawQuery)))
		return nil
	}

	body, err := readRequestBody(r)
	if err != nil {
		return err
	}
	trimmed := bytes.TrimSpace(body)
	if len(trimmed) == 0 {
		trimmed = []byte("{}")
	}
	if trimmed[0] != '{' {
		return fmt.Errorf("payrexx: cannot sign a %s body that is not a JSON object", r.Method)
	}

	encoded, err := phpQueryFromJSON(trimmed)
	if err != nil {
		return fmt.Errorf("payrexx: signing %s %s: %w", r.Method, r.URL.Path, err)
	}
	sig, err := json.Marshal(t.signature(encoded))
	if err != nil {
		return err
	}

	// Splice ApiSignature in as the last member rather than re-marshalling, so
	// every other byte the caller produced reaches Payrexx untouched.
	var out bytes.Buffer
	out.Write(trimmed[:len(trimmed)-1])
	if !bytes.Equal(trimmed, []byte("{}")) {
		out.WriteByte(',')
	}
	out.WriteString(`"ApiSignature":`)
	out.Write(sig)
	out.WriteByte('}')
	setRequestBody(r, out.Bytes())
	return nil
}

func (t *authTransport) signature(query string) string {
	mac := hmac.New(sha256.New, []byte(t.secret))
	mac.Write([]byte(query))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func readRequestBody(r *http.Request) ([]byte, error) {
	if r.Body == nil {
		return nil, nil
	}
	body, err := io.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		return nil, err
	}
	setRequestBody(r, body)
	return body, nil
}

func setRequestBody(r *http.Request, body []byte) {
	r.Body = io.NopCloser(bytes.NewReader(body))
	r.ContentLength = int64(len(body))
	r.GetBody = func() (io.ReadCloser, error) {
		return io.NopCloser(bytes.NewReader(body)), nil
	}
}

// phpEscape encodes a string the way PHP's urlencode does: every byte outside
// [A-Za-z0-9_-.] percent-encoded, and a space as "+". Go's url.QueryEscape is
// almost this but leaves "~" alone, which would change the signed bytes.
func phpEscape(s string) string {
	const unreserved = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_."
	var b strings.Builder
	b.Grow(len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case strings.IndexByte(unreserved, c) >= 0:
			b.WriteByte(c)
		case c == ' ':
			b.WriteByte('+')
		default:
			b.WriteString(fmt.Sprintf("%%%02X", c))
		}
	}
	return b.String()
}

// phpQueryFromURLQuery re-renders an already-encoded query string in PHP's
// encoding, preserving parameter order. Payrexx's server parses what it receives
// and rebuilds the string itself, so only the order and the values have to
// match — but the escaping of those values has to match exactly, and Go and PHP
// disagree on a handful of characters.
func phpQueryFromURLQuery(raw string) string {
	var parts []string
	for _, pair := range strings.Split(raw, "&") {
		if pair == "" {
			continue
		}
		name, value, _ := strings.Cut(pair, "=")
		parts = append(parts, phpEscape(urlDecode(name))+"="+phpEscape(urlDecode(value)))
	}
	return strings.Join(parts, "&")
}

func urlDecode(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for i := 0; i < len(s); i++ {
		switch {
		case s[i] == '+':
			b.WriteByte(' ')
		case s[i] == '%' && i+2 < len(s):
			if v, err := strconv.ParseUint(s[i+1:i+3], 16, 8); err == nil {
				b.WriteByte(byte(v))
				i += 2
				continue
			}
			b.WriteByte(s[i])
		default:
			b.WriteByte(s[i])
		}
	}
	return b.String()
}

// phpQueryFromJSON renders a JSON object as PHP's http_build_query would render
// the equivalent array: member order preserved, nested objects and arrays as
// name[key] and name[0], booleans as 1 and 0, and nulls dropped entirely.
func phpQueryFromJSON(body []byte) (string, error) {
	dec := json.NewDecoder(bytes.NewReader(body))
	dec.UseNumber()
	if _, err := dec.Token(); err != nil { // consume '{'
		return "", err
	}
	var parts []string
	if err := appendObject(dec, "", &parts); err != nil {
		return "", err
	}
	return strings.Join(parts, "&"), nil
}

// appendObject consumes members up to the closing '}' of the object whose
// opening token has already been read.
func appendObject(dec *json.Decoder, prefix string, parts *[]string) error {
	for dec.More() {
		tok, err := dec.Token()
		if err != nil {
			return err
		}
		key, ok := tok.(string)
		if !ok {
			return fmt.Errorf("expected an object key, got %v", tok)
		}
		name := key
		if prefix != "" {
			name = prefix + "[" + key + "]"
		}
		if err := appendValue(dec, name, parts); err != nil {
			return err
		}
	}
	_, err := dec.Token() // consume '}'
	return err
}

func appendValue(dec *json.Decoder, name string, parts *[]string) error {
	tok, err := dec.Token()
	if err != nil {
		return err
	}
	switch v := tok.(type) {
	case json.Delim:
		switch v {
		case '{':
			return appendObject(dec, name, parts)
		case '[':
			for i := 0; dec.More(); i++ {
				if err := appendValue(dec, name+"["+strconv.Itoa(i)+"]", parts); err != nil {
					return err
				}
			}
			_, err := dec.Token() // consume ']'
			return err
		}
		return fmt.Errorf("unexpected delimiter %v", v)
	case nil:
		return nil // http_build_query skips null members
	case string:
		*parts = append(*parts, phpEscape(name)+"="+phpEscape(v))
	case json.Number:
		*parts = append(*parts, phpEscape(name)+"="+phpEscape(v.String()))
	case bool:
		n := "0"
		if v {
			n = "1"
		}
		*parts = append(*parts, phpEscape(name)+"="+n)
	default:
		return fmt.Errorf("unexpected value %v for %q", tok, name)
	}
	return nil
}
