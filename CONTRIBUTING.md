# Contributing to aveecare-go

Thanks for helping improve the AveeCare Go library.

## How this library is built

- **Models and resource services are generated** from the public AveeCare OpenAPI
  specification, published in
  [aveecare-openapi](https://github.com/aveecare/aveecare-openapi). The files ending
  in `_gen.go` (`models_gen.go`, `resources_gen.go`, `client_gen.go`) are overwritten
  on each release, so please do not edit them by hand. If a model, field or endpoint
  is wrong or missing, open an issue here or in aveecare-openapi and we will fix it
  at the source.
- **The runtime is hand-written** (configuration, authentication, retries,
  pagination, errors). Fixes and improvements to it are very welcome as pull
  requests.

## Development

You need Go 1.21 or newer.

```sh
go vet ./...
go test ./...
gofmt -l .        # must print nothing
```

Unit tests use a local test server and never touch the network. Live tests run
only when both `AVEECARE_API_KEY` and `AVEECARE_BASE_URL` are set, and they
create, update and delete one test patient, so point them at a sandbox:

```sh
AVEECARE_API_KEY=ak_... AVEECARE_BASE_URL=https://... go test -run Live -v
```

## Pull requests

1. Open an issue first for anything larger than a small fix.
2. Add or update tests for the behaviour you change.
3. Keep the public API backwards compatible within a major version.
4. Add a line to `CHANGELOG.md` under "Unreleased".

By contributing you agree that your contributions are licensed under the MIT
License, and you agree to follow our [Code of Conduct](CODE_OF_CONDUCT.md).

Security issues: please follow [SECURITY.md](SECURITY.md) instead of opening a
public issue.

## Releasing (maintainers)

Go modules are published by tag; there is no registry account or secret to
configure. Update `Version` in `version.go` and `CHANGELOG.md`, commit, then push
a tag matching the version:

```sh
git tag v0.1.1 && git push origin v0.1.1
```

The release workflow checks the tag against `version.go`, runs the tests and
creates a GitHub release. pkg.go.dev picks the new version up on first request
(or run `GOPROXY=https://proxy.golang.org go list -m github.com/aveecare/aveecare-go@v0.1.1`).
