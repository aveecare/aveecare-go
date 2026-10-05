package aveecare

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// ErrMissingAPIKey is returned by NewClient when no API key was given and
// AVEECARE_API_KEY is not set.
var ErrMissingAPIKey = errors.New("aveecare: no API key provided: use WithAPIKey or set AVEECARE_API_KEY (see https://www.aveecare.com/docs/api/authentication)")

// ErrNoMorePages is returned by Page.NextPage when there is no following page.
var ErrNoMorePages = errors.New("aveecare: no more pages; check HasNextPage before calling NextPage")

// APIError is an error response from the API: any non-2xx status. The
// specific types below (NotFoundError, RateLimitError, ...) embed it and are
// chosen by status code, so both of these work:
//
//	var notFound *aveecare.NotFoundError
//	if errors.As(err, &notFound) { ... }
//
//	var apiErr *aveecare.APIError // any non-2xx response
//	if errors.As(err, &apiErr) { log.Println(apiErr.Status, apiErr.Code, apiErr.RequestID) }
//
// Status 400 and 422 are InvalidRequestError, 401 AuthenticationError, 403
// PermissionError, 404 NotFoundError, 409 ConflictError and 429
// RateLimitError. Every other status, including 5xx, is a plain APIError.
// Failures before a response arrives are ConnectionError and TimeoutError.
type APIError struct {
	// Status is the HTTP status code.
	Status int
	// Code is the machine-readable code such as "not_found" or "insufficient_scope".
	Code string
	// Type is the broad category such as "invalid_request_error".
	Type string
	// Message is the human-readable explanation.
	Message string
	// RequestID is the Request-Id of the failed request. Quote it when contacting support.
	RequestID string
	// Header holds the response headers.
	Header http.Header
}

// Error implements the error interface.
func (e *APIError) Error() string {
	var b strings.Builder
	b.WriteString(e.Message)
	fmt.Fprintf(&b, " (status %d", e.Status)
	if e.Code != "" {
		fmt.Fprintf(&b, ", code %s", e.Code)
	}
	if e.RequestID != "" {
		fmt.Fprintf(&b, ", request %s", e.RequestID)
	}
	b.WriteString(")")
	return b.String()
}

// InvalidRequestError is a 400 or 422 response, for example codes
// invalid_request and unknown_field.
type InvalidRequestError struct{ APIError }

// AuthenticationError is a 401 response (unauthorized).
type AuthenticationError struct{ APIError }

// PermissionError is a 403 response (forbidden, insufficient_scope).
type PermissionError struct{ APIError }

// NotFoundError is a 404 response (not_found).
type NotFoundError struct{ APIError }

// ConflictError is a 409 response (idempotency_conflict).
type ConflictError struct{ APIError }

// RateLimitError is a 429 response (rate_limited).
type RateLimitError struct {
	APIError
	// RetryAfter is how long the API asked us to wait (from Retry-After); zero when not given.
	RetryAfter time.Duration
}

// Unwrap returns the embedded APIError so errors.As matches *APIError too.
func (e *InvalidRequestError) Unwrap() error { return &e.APIError }

// Unwrap returns the embedded APIError so errors.As matches *APIError too.
func (e *AuthenticationError) Unwrap() error { return &e.APIError }

// Unwrap returns the embedded APIError so errors.As matches *APIError too.
func (e *PermissionError) Unwrap() error { return &e.APIError }

// Unwrap returns the embedded APIError so errors.As matches *APIError too.
func (e *NotFoundError) Unwrap() error { return &e.APIError }

// Unwrap returns the embedded APIError so errors.As matches *APIError too.
func (e *ConflictError) Unwrap() error { return &e.APIError }

// Unwrap returns the embedded APIError so errors.As matches *APIError too.
func (e *RateLimitError) Unwrap() error { return &e.APIError }

// ConnectionError means the request never got a response (DNS, TLS, a reset
// connection) or was canceled through its context. Err holds the cause, so
// errors.Is(err, context.Canceled) works.
type ConnectionError struct {
	Message string
	Err     error
}

// Error implements the error interface.
func (e *ConnectionError) Error() string {
	if e.Err == nil {
		return e.Message
	}
	return e.Message + ": " + e.Err.Error()
}

// Unwrap returns the underlying cause.
func (e *ConnectionError) Unwrap() error { return e.Err }

// TimeoutError means a request attempt exceeded its timeout. It is also a
// ConnectionError, so errors.As matches either, and errors.Is(err,
// context.DeadlineExceeded) is true.
type TimeoutError struct{ ConnectionError }

// Timeout reports true, satisfying the net.Error convention.
func (e *TimeoutError) Timeout() bool { return true }

// Unwrap returns the embedded ConnectionError.
func (e *TimeoutError) Unwrap() error { return &e.ConnectionError }

// InvalidResponseError means a 2xx response body was not the JSON this library
// expected.
type InvalidResponseError struct {
	Status    int
	RequestID string
	Err       error
}

// Error implements the error interface.
func (e *InvalidResponseError) Error() string {
	msg := fmt.Sprintf("aveecare: unexpected response (status %d", e.Status)
	if e.RequestID != "" {
		msg += ", request " + e.RequestID
	}
	return msg + "): " + e.Err.Error()
}

// Unwrap returns the decoding error.
func (e *InvalidResponseError) Unwrap() error { return e.Err }

type errorEnvelope struct {
	Error *struct {
		Type      string `json:"type"`
		Code      string `json:"code"`
		Message   string `json:"message"`
		RequestID string `json:"request_id"`
	} `json:"error"`
}

// newAPIError turns a non-2xx response into the typed error for its status.
func newAPIError(status int, body []byte, header http.Header) error {
	base := APIError{Status: status, Header: header, RequestID: header.Get("Request-Id")}
	var env errorEnvelope
	if err := json.Unmarshal(body, &env); err == nil && env.Error != nil {
		base.Code = env.Error.Code
		base.Type = env.Error.Type
		base.Message = env.Error.Message
		if env.Error.RequestID != "" {
			base.RequestID = env.Error.RequestID
		}
	}
	if base.Message == "" {
		base.Message = fmt.Sprintf("Request failed with status %d", status)
	}
	switch status {
	case http.StatusBadRequest, http.StatusUnprocessableEntity:
		return &InvalidRequestError{base}
	case http.StatusUnauthorized:
		return &AuthenticationError{base}
	case http.StatusForbidden:
		return &PermissionError{base}
	case http.StatusNotFound:
		return &NotFoundError{base}
	case http.StatusConflict:
		return &ConflictError{base}
	case http.StatusTooManyRequests:
		d, _ := parseRetryAfter(header.Get("Retry-After"), time.Now())
		return &RateLimitError{APIError: base, RetryAfter: d}
	default:
		return &base
	}
}
