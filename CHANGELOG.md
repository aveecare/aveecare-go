# Changelog

All notable changes to this project are documented here. The format follows
[Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and the project uses
[Semantic Versioning](https://semver.org/).

## [Unreleased]

## [0.1.0] - 2026-10-05

### Added

- First release of the AveeCare Go library.
- Typed services for every public API resource, with context-first methods and
  functional options.
- Cursor pagination with `Page`, `NextPage` and an `Iterator` (plus `iter.Seq2`
  on Go 1.23 and newer) for automatic paging.
- Typed errors that work with `errors.As`, carrying the error code, HTTP status
  and request ID.
- Automatic retries with jittered exponential backoff for rate limits (honoring
  `Retry-After`), server errors and network failures.
- Automatic `Idempotency-Key` on every create request.
