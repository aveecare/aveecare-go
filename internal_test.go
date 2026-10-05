package aveecare

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

func TestBackoffIsJitteredExponentialAndCapped(t *testing.T) {
	c := &core{jitter: func() float64 { return 0 }}
	want := []time.Duration{500 * time.Millisecond, time.Second, 2 * time.Second, 4 * time.Second, 8 * time.Second, 8 * time.Second}
	for attempt, w := range want {
		if got := c.backoff(attempt); got != w {
			t.Errorf("attempt %d: %v, want %v", attempt, got, w)
		}
	}
	c.jitter = func() float64 { return 1 }
	if got := c.backoff(0); got != 375*time.Millisecond {
		t.Errorf("full jitter takes 25%% off: got %v", got)
	}
	if got := c.backoff(1000); got <= 0 || got > 8*time.Second {
		t.Errorf("huge attempt numbers must stay capped: %v", got)
	}
}

func TestParseRetryAfter(t *testing.T) {
	now := time.Date(2026, 10, 5, 12, 0, 0, 0, time.UTC)
	cases := []struct {
		in   string
		want time.Duration
		ok   bool
	}{
		{"", 0, false},
		{"0", 0, true},
		{"7", 7 * time.Second, true},
		{"1.5", 1500 * time.Millisecond, true},
		{"-3", 0, false},
		{"soon", 0, false},
		{"Mon, 05 Oct 2026 12:00:30 GMT", 30 * time.Second, true},
		{"Mon, 05 Oct 2026 11:00:00 GMT", 0, true},
	}
	for _, tc := range cases {
		got, ok := parseRetryAfter(tc.in, now)
		if got != tc.want || ok != tc.ok {
			t.Errorf("parseRetryAfter(%q) = %v, %v; want %v, %v", tc.in, got, ok, tc.want, tc.ok)
		}
	}
}

func TestRetryDelaysComeFromRetryAfterThenBackoff(t *testing.T) {
	var calls int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch atomic.AddInt32(&calls, 1) {
		case 1:
			w.Header().Set("Retry-After", "3")
			w.WriteHeader(429)
		case 2:
			w.WriteHeader(500)
		case 3:
			w.Header().Set("Retry-After", "86400") // capped at one minute
			w.WriteHeader(503)
		default:
			_, _ = w.Write([]byte(`{}`))
		}
	}))
	defer srv.Close()
	c, err := newCore([]Option{WithAPIKey("k"), WithBaseURL(srv.URL), WithMaxRetries(3)})
	if err != nil {
		t.Fatal(err)
	}
	var slept []time.Duration
	c.sleep = func(d time.Duration, _ <-chan struct{}) bool {
		slept = append(slept, d)
		return true
	}
	c.jitter = func() float64 { return 0 }
	var out map[string]any
	if err := c.do(context.Background(), requestSpec{method: "GET", path: "/v1/me"}, newRequestOptions(nil), &out); err != nil {
		t.Fatal(err)
	}
	want := []time.Duration{3 * time.Second, time.Second, time.Minute}
	if len(slept) != len(want) {
		t.Fatalf("slept %v, want %v", slept, want)
	}
	for i := range want {
		if slept[i] != want[i] {
			t.Errorf("sleep %d = %v, want %v", i, slept[i], want[i])
		}
	}
}

func TestRetriesRequestTimeouts408(t *testing.T) {
	var calls int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if atomic.AddInt32(&calls, 1) == 1 {
			w.WriteHeader(408)
			return
		}
		_, _ = w.Write([]byte(`{}`))
	}))
	defer srv.Close()
	c, _ := newCore([]Option{WithAPIKey("k"), WithBaseURL(srv.URL)})
	c.sleep = func(time.Duration, <-chan struct{}) bool { return true }
	var out map[string]any
	if err := c.do(context.Background(), requestSpec{method: "GET", path: "/x"}, newRequestOptions(nil), &out); err != nil || calls != 2 {
		t.Fatalf("err %v calls %d", err, calls)
	}
}

func TestUUID4Format(t *testing.T) {
	seen := map[string]bool{}
	for i := 0; i < 50; i++ {
		u := uuid4()
		if len(u) != 36 || u[14] != '4' || (u[19] != '8' && u[19] != '9' && u[19] != 'a' && u[19] != 'b') {
			t.Fatalf("not a v4 UUID: %q", u)
		}
		if seen[u] {
			t.Fatalf("duplicate UUID %q", u)
		}
		seen[u] = true
	}
}
