# AveeCare Go Library

[![Go Reference](https://pkg.go.dev/badge/github.com/aveecare/aveecare-go.svg)](https://pkg.go.dev/github.com/aveecare/aveecare-go)
[![CI](https://github.com/aveecare/aveecare-go/actions/workflows/ci.yml/badge.svg)](https://github.com/aveecare/aveecare-go/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/aveecare/aveecare-go/blob/main/LICENSE)

The official Go library for the [AveeCare API](https://www.aveecare.com/docs/api).
Read and write your home care agency's patients, visits, invoices and more from your own
servers, with typed models, automatic pagination, retries and idempotency built in.

- Standard library only (`net/http`), Go 1.21 or newer
- Context-first methods, functional options, typed errors for `errors.As`
- Automatic retries with backoff and automatic idempotency keys

## Install

```sh
go get github.com/aveecare/aveecare-go
```

## Quickstart

Create an API key in AveeCare, then:

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aveecare/aveecare-go"
)

func main() {
	client, err := aveecare.NewClient(aveecare.WithAPIKey(os.Getenv("AVEECARE_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}

	page, err := client.Patients.List(context.Background(), &aveecare.PatientListParams{
		Active: aveecare.Bool(true),
		Limit:  aveecare.Int(25),
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, patient := range page.Data {
		fmt.Println(patient.ID)
	}
}
```

## Authentication

Every request is authenticated with your secret API key, sent as
`Authorization: Bearer ak_<keyId>_<secret>`. Pass it with `WithAPIKey`, or leave it
out and set `AVEECARE_API_KEY`: `aveecare.NewClient()` reads it for you and returns
`aveecare.ErrMissingAPIKey` when there is none. Keys are read-only or read and write;
see [Authentication](https://www.aveecare.com/docs/api/authentication).

Keep keys on your server. Never put an API key in browser or mobile code.

## Usage

```go
func example(ctx context.Context, client *aveecare.Client) error {
	// Retrieve, create, update and delete
	patient, err := client.Patients.Retrieve(ctx, "pat_8f2k1m9x")
	if err != nil {
		return err
	}
	created, err := client.Patients.Create(ctx, &aveecare.PatientCreate{
		FirstName: "Margaret",
		LastName:  "Alvarez",
	})
	if err != nil {
		return err
	}
	updated, err := client.Patients.Update(ctx, created.ID, &aveecare.PatientUpdate{
		PrimaryLanguage: aveecare.Set("Spanish"),
		Email:           aveecare.Null[string](), // clears the email
	})
	if err != nil {
		return err
	}
	if _, err := client.Patients.Delete(ctx, created.ID); err != nil {
		return err
	}

	// Visits are listed by start-time window (at most 93 days)
	visits, err := client.Visits.List(ctx, &aveecare.VisitListParams{
		StartFrom: "2026-10-01T00:00:00Z",
		StartTo:   "2026-10-08T00:00:00Z",
	})
	if err != nil {
		return err
	}

	// The API key making the request
	me, err := client.Me.Retrieve(ctx)
	if err != nil {
		return err
	}
	fmt.Println(patient.ID, updated.ID, len(visits.Data), me.ID)
	return nil
}
```

Fields are Go-cased (`FirstName`, `PatientID`) and map to the API's camelCase JSON.
In a response, a field that can be `null` is a pointer (`nil` when null or absent).
In a request, required fields are plain values and optional fields are pointers, so
"not given" is different from the zero value: use `aveecare.String`, `Int`, `Bool`,
`Float` or the generic `aveecare.Ptr` to make one. Fields that can be cleared on an
update are `aveecare.Nullable[T]`: leave it zero to omit the field,
`aveecare.Set(v)` to set it and `aveecare.Null[T]()` to send `null`.
String enums are typed strings, and a value added to the API later decodes as it
arrives rather than failing.

## Pagination

`List` returns one `Page`. Each list method has a `ListAutoPaging` twin returning an
`Iterator` over every item across every page, fetching pages lazily:

```go
it := client.Patients.ListAutoPaging(ctx, &aveecare.PatientListParams{Limit: aveecare.Int(100)})
for it.Next() {
	fmt.Println(it.Item().ID)
}
if err := it.Err(); err != nil {
	return err
}
```

On Go 1.23 and newer the iterator also works with `range`, yielding any error once:

```go
for patient, err := range client.Patients.ListAutoPaging(ctx, nil).All() {
	if err != nil {
		return err
	}
	fmt.Println(patient.ID)
}
```

To page by hand, use `page.HasNextPage()` and `page.NextPage(ctx)`.

## Errors

Errors are typed. Use `errors.As` to match one, or `*aveecare.APIError` to match any
non-2xx response:

```go
_, err := client.Patients.Retrieve(ctx, "pat_missing")

var notFound *aveecare.NotFoundError
var apiErr *aveecare.APIError
switch {
case err == nil:
case errors.As(err, &notFound):
	fmt.Println("no such patient", notFound.RequestID)
case errors.As(err, &apiErr):
	fmt.Println(apiErr.Status, apiErr.Code, apiErr.Message)
}
```

| Status | Type | Example codes |
|---|---|---|
| 400 | `*aveecare.InvalidRequestError` | `invalid_request`, `unknown_field` |
| 401 | `*aveecare.AuthenticationError` | `unauthorized` |
| 403 | `*aveecare.PermissionError` | `forbidden`, `insufficient_scope` |
| 404 | `*aveecare.NotFoundError` | `not_found` |
| 409 | `*aveecare.ConflictError` | `idempotency_conflict` |
| 429 | `*aveecare.RateLimitError` | `rate_limited` |
| other | `*aveecare.APIError` | `internal_error` |

Every one carries `Status`, `Code`, `Type`, `Message` and `RequestID`. Network failures
are `*aveecare.ConnectionError`, and timeouts `*aveecare.TimeoutError`; a cancelled
context returns the context's own error.

## Retries, timeouts and idempotency

Rate-limited (429), server error (5xx) and network failures are retried twice by
default with jittered exponential backoff. When the API sends `Retry-After`, the client
waits that long.

Every `Create` sends an `Idempotency-Key` (a random UUID), and the same key is reused
across retries, so a retried create never makes a duplicate. Pass your own key to make
a create safe to repeat across processes:

```go
client, err := aveecare.NewClient(
	aveecare.WithAPIKey(os.Getenv("AVEECARE_API_KEY")),
	aveecare.WithMaxRetries(4),
	aveecare.WithTimeout(20*time.Second),
)
if err != nil {
	return err
}
_, err = client.Patients.Create(ctx,
	&aveecare.PatientCreate{FirstName: "Margaret", LastName: "Alvarez"},
	aveecare.WithIdempotencyKey("import-row-1042"),
)
```

## Configuration

`NewClient` accepts `WithAPIKey` (default `AVEECARE_API_KEY`), `WithBaseURL` (default
`AVEECARE_BASE_URL`, then `https://api.aveecare.com`), `WithTimeout` (60 seconds per
attempt), `WithMaxRetries` (2), `WithHTTPClient` (your own `*http.Client`) and
`WithHeader`. Per call, pass `WithIdempotencyKey`, `WithRequestTimeout`,
`WithRequestMaxRetries` or `WithRequestHeader` as the last arguments. Every request
sends `User-Agent: aveecare-go/<version>`. A `Client` is safe for concurrent use.

## Documentation

The full API reference, with every resource and field, is at
[www.aveecare.com/docs/api](https://www.aveecare.com/docs/api), and the package docs
are on [pkg.go.dev](https://pkg.go.dev/github.com/aveecare/aveecare-go).
Building an integration to offer other agencies? See the
[AveeCare Partner Center](https://www.aveecare.com/partners).

## Other official SDKs

- [aveecare-node](https://github.com/aveecare/aveecare-node): Node.js and TypeScript
- [aveecare-python](https://github.com/aveecare/aveecare-python): Python
- [aveecare-rust](https://github.com/aveecare/aveecare-rust): Rust
- [aveecare-go](https://github.com/aveecare/aveecare-go): Go
- [aveecare-ruby](https://github.com/aveecare/aveecare-ruby): Ruby
- [aveecare-php](https://github.com/aveecare/aveecare-php): PHP
- [aveecare-java](https://github.com/aveecare/aveecare-java): Java and Kotlin
- [aveecare-dotnet](https://github.com/aveecare/aveecare-dotnet): .NET
- [aveecare-cli](https://github.com/aveecare/aveecare-cli): command line
- [aveecare-openapi](https://github.com/aveecare/aveecare-openapi): OpenAPI specification
- [aveecare-mcp](https://github.com/aveecare/aveecare-mcp): connect AI assistants
- [aveecare-examples](https://github.com/aveecare/aveecare-examples): sample apps

## License

MIT. See [LICENSE](https://github.com/aveecare/aveecare-go/blob/main/LICENSE).

---

Built by [AveeCare](https://www.aveecare.com) — home care software.
