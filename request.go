package aveecare

import (
	"bytes"
	"context"
	crand "crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const maxRetryDelay = 60 * time.Second

type requestSpec struct {
	method     string
	path       string
	query      url.Values
	body       any // nil sends no body
	idempotent bool
}

type response struct {
	status int
	header http.Header
	body   []byte
}

// do sends one request, retrying 408, 429, 5xx and network errors, and decodes
// a 2xx JSON body into out (when out is not nil).
func (c *core) do(ctx context.Context, spec requestSpec, ro *requestOptions, out any) error {
	target := c.baseURL + spec.path
	if len(spec.query) > 0 {
		target += "?" + spec.query.Encode()
	}
	var payload []byte
	if spec.body != nil {
		b, err := json.Marshal(spec.body)
		if err != nil {
			return fmt.Errorf("aveecare: could not encode the request body: %w", err)
		}
		payload = b
	}

	headers := http.Header{}
	headers.Set("Accept", "application/json")
	headers.Set("Authorization", "Bearer "+c.apiKey)
	headers.Set("User-Agent", "aveecare-go/"+Version)
	for k, vs := range c.headers {
		headers[k] = append([]string(nil), vs...)
	}
	for k, vs := range ro.headers {
		headers[k] = append([]string(nil), vs...)
	}
	if payload != nil {
		headers.Set("Content-Type", "application/json")
	}
	if spec.method == http.MethodPost || spec.idempotent {
		key := ro.idempotencyKey
		if key == "" {
			key = headers.Get("Idempotency-Key")
		}
		if key == "" {
			key = uuid4()
		}
		headers.Set("Idempotency-Key", key)
	}

	maxRetries := c.maxRetries
	if ro.maxRetries != nil {
		maxRetries = *ro.maxRetries
	}
	timeout := c.timeout
	if ro.timeout != nil {
		timeout = *ro.timeout
	}

	for attempt := 0; ; attempt++ {
		res, err := c.send(ctx, spec.method, target, headers, payload, timeout)
		if err != nil {
			var conn *ConnectionError
			if errors.As(err, &conn) && ctx.Err() == nil && attempt < maxRetries {
				if !c.sleep(c.backoff(attempt), ctx.Done()) {
					return canceled(ctx)
				}
				continue
			}
			return err
		}
		if res.status >= 200 && res.status < 300 {
			if out == nil || len(bytes.TrimSpace(res.body)) == 0 {
				return nil
			}
			if err := json.Unmarshal(res.body, out); err != nil {
				return &InvalidResponseError{Status: res.status, RequestID: res.header.Get("Request-Id"), Err: err}
			}
			return nil
		}
		if shouldRetry(res.status) && attempt < maxRetries {
			delay, ok := parseRetryAfter(res.header.Get("Retry-After"), time.Now())
			if !ok {
				delay = c.backoff(attempt)
			}
			if delay > maxRetryDelay {
				delay = maxRetryDelay
			}
			if !c.sleep(delay, ctx.Done()) {
				return canceled(ctx)
			}
			continue
		}
		return newAPIError(res.status, res.body, res.header)
	}
}

// send makes one attempt under its own timeout. Transport failures come back
// as *ConnectionError or *TimeoutError.
func (c *core) send(ctx context.Context, method, target string, headers http.Header, payload []byte, timeout time.Duration) (*response, error) {
	actx := ctx
	if timeout > 0 {
		var cancel context.CancelFunc
		actx, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()
	}
	var body io.Reader
	if payload != nil {
		body = bytes.NewReader(payload)
	}
	req, err := http.NewRequestWithContext(actx, method, target, body)
	if err != nil {
		return nil, fmt.Errorf("aveecare: could not build the request: %w", err)
	}
	req.Header = headers.Clone()
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, classify(ctx, actx, timeout, err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, classify(ctx, actx, timeout, err)
	}
	return &response{status: resp.StatusCode, header: resp.Header, body: data}, nil
}

func classify(ctx, attemptCtx context.Context, timeout time.Duration, err error) error {
	if ctx.Err() != nil {
		return canceled(ctx)
	}
	if attemptCtx.Err() == context.DeadlineExceeded {
		return &TimeoutError{ConnectionError{
			Message: fmt.Sprintf("request timed out after %s", timeout),
			Err:     context.DeadlineExceeded,
		}}
	}
	return &ConnectionError{Message: "could not reach the AveeCare API", Err: err}
}

func canceled(ctx context.Context) error {
	return &ConnectionError{Message: "request canceled", Err: ctx.Err()}
}

func shouldRetry(status int) bool {
	return status == http.StatusRequestTimeout || status == http.StatusTooManyRequests || status >= 500
}

// backoff is jittered exponential backoff: 0.5s, 1s, 2s, ... capped at 8s,
// minus up to 25%.
func (c *core) backoff(attempt int) time.Duration {
	if attempt > 10 {
		attempt = 10
	}
	base := math.Min(0.5*math.Pow(2, float64(attempt)), 8)
	return time.Duration(base * (1 - c.jitter()*0.25) * float64(time.Second))
}

// parseRetryAfter reads Retry-After as seconds or an HTTP date.
func parseRetryAfter(value string, now time.Time) (time.Duration, bool) {
	value = strings.TrimSpace(value)
	if value == "" {
		return 0, false
	}
	if secs, err := strconv.ParseFloat(value, 64); err == nil {
		if secs < 0 || math.IsNaN(secs) || math.IsInf(secs, 0) {
			return 0, false
		}
		return time.Duration(secs * float64(time.Second)), true
	}
	if t, err := http.ParseTime(value); err == nil {
		if d := t.Sub(now); d > 0 {
			return d, true
		}
		return 0, true
	}
	return 0, false
}

// sleepOrDone waits d, returning false if done closes first.
func sleepOrDone(d time.Duration, done <-chan struct{}) bool {
	if d <= 0 {
		select {
		case <-done:
			return false
		default:
			return true
		}
	}
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-t.C:
		return true
	case <-done:
		return false
	}
}

func randFloat() float64 { return rand.Float64() }

func uuid4() string {
	var b [16]byte
	if _, err := crand.Read(b[:]); err != nil {
		for i := range b {
			b[i] = byte(rand.Intn(256))
		}
	}
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	h := hex.EncodeToString(b[:])
	return h[0:8] + "-" + h[8:12] + "-" + h[12:16] + "-" + h[16:20] + "-" + h[20:]
}

// pathEscape percent-encodes one path segment.
func pathEscape(segment string) string { return url.PathEscape(segment) }
