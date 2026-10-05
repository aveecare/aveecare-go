package aveecare

import (
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	// DefaultBaseURL is the API origin used when none is configured.
	DefaultBaseURL = "https://api.aveecare.com"
	// DefaultTimeout is the default per-attempt timeout.
	DefaultTimeout = 60 * time.Second
	// DefaultMaxRetries is the default number of retries for 429, 5xx and network errors.
	DefaultMaxRetries = 2
)

// Option configures a Client. Pass options to NewClient.
type Option func(*config)

type config struct {
	apiKey     string
	baseURL    string
	timeout    time.Duration
	maxRetries int
	httpClient *http.Client
	headers    http.Header
}

// WithAPIKey sets your secret key, ak_<keyId>_<secret>. It defaults to the
// AVEECARE_API_KEY environment variable.
func WithAPIKey(key string) Option {
	return func(c *config) { c.apiKey = key }
}

// WithBaseURL sets the API origin. It defaults to the AVEECARE_BASE_URL
// environment variable, then https://api.aveecare.com.
func WithBaseURL(url string) Option {
	return func(c *config) { c.baseURL = url }
}

// WithTimeout sets the per-attempt timeout (default 60 seconds). Zero or
// negative disables it, leaving only the request context to bound a call.
func WithTimeout(d time.Duration) Option {
	return func(c *config) { c.timeout = d }
}

// WithMaxRetries sets how many times a request is retried after a 429, 5xx or
// network error (default 2). Zero disables retries.
func WithMaxRetries(n int) Option {
	return func(c *config) { c.maxRetries = n }
}

// WithHTTPClient uses your own *http.Client (proxies, custom TLS, tracing).
func WithHTTPClient(hc *http.Client) Option {
	return func(c *config) { c.httpClient = hc }
}

// WithHeader adds a header sent with every request.
func WithHeader(key, value string) Option {
	return func(c *config) { c.headers.Set(key, value) }
}

func env(name string) string { return strings.TrimSpace(os.Getenv(name)) }

// core is the transport shared by every resource service.
type core struct {
	apiKey     string
	baseURL    string
	timeout    time.Duration
	maxRetries int
	http       *http.Client
	headers    http.Header

	// Overridable in tests.
	sleep  func(d time.Duration, done <-chan struct{}) bool
	jitter func() float64
}

func newCore(opts []Option) (*core, error) {
	cfg := config{
		timeout:    DefaultTimeout,
		maxRetries: DefaultMaxRetries,
		headers:    http.Header{},
	}
	for _, opt := range opts {
		opt(&cfg)
	}
	apiKey := strings.TrimSpace(cfg.apiKey)
	if apiKey == "" {
		apiKey = env("AVEECARE_API_KEY")
	}
	if apiKey == "" {
		return nil, ErrMissingAPIKey
	}
	baseURL := strings.TrimSpace(cfg.baseURL)
	if baseURL == "" {
		baseURL = env("AVEECARE_BASE_URL")
	}
	if baseURL == "" {
		baseURL = DefaultBaseURL
	}
	hc := cfg.httpClient
	if hc == nil {
		hc = &http.Client{}
	}
	return &core{
		apiKey:     apiKey,
		baseURL:    strings.TrimRight(baseURL, "/"),
		timeout:    cfg.timeout,
		maxRetries: cfg.maxRetries,
		http:       hc,
		headers:    cfg.headers,
		sleep:      sleepOrDone,
		jitter:     randFloat,
	}, nil
}

// NewClient creates a Client. With no options it reads AVEECARE_API_KEY and
// AVEECARE_BASE_URL from the environment; it returns ErrMissingAPIKey when no
// key is available. A Client is safe for concurrent use.
//
//	client, err := aveecare.NewClient(aveecare.WithAPIKey(os.Getenv("AVEECARE_API_KEY")))
func NewClient(opts ...Option) (*Client, error) {
	c, err := newCore(opts)
	if err != nil {
		return nil, err
	}
	return newClient(c), nil
}

// BaseURL returns the API origin requests go to.
func (c *Client) BaseURL() string { return c.core.baseURL }

// RequestOption overrides settings for a single call. Pass it as the last
// arguments of any resource method.
type RequestOption func(*requestOptions)

type requestOptions struct {
	idempotencyKey string
	timeout        *time.Duration
	maxRetries     *int
	headers        http.Header
}

func newRequestOptions(opts []RequestOption) *requestOptions {
	ro := &requestOptions{}
	for _, opt := range opts {
		opt(ro)
	}
	return ro
}

// WithIdempotencyKey sets the Idempotency-Key of a POST. One is generated for
// you when omitted. The key is reused across that request's retries; pass your
// own to make a create safe to repeat across processes.
func WithIdempotencyKey(key string) RequestOption {
	return func(r *requestOptions) { r.idempotencyKey = key }
}

// WithRequestTimeout overrides the per-attempt timeout for one call.
func WithRequestTimeout(d time.Duration) RequestOption {
	return func(r *requestOptions) { r.timeout = &d }
}

// WithRequestMaxRetries overrides the retry count for one call.
func WithRequestMaxRetries(n int) RequestOption {
	return func(r *requestOptions) { r.maxRetries = &n }
}

// WithRequestHeader adds a header to one call.
func WithRequestHeader(key, value string) RequestOption {
	return func(r *requestOptions) {
		if r.headers == nil {
			r.headers = http.Header{}
		}
		r.headers.Set(key, value)
	}
}
