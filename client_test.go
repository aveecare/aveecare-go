package aveecare_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/aveecare/aveecare-go"
)

const testKey = "ak_test_secret123"

type recorded struct {
	Method string
	Path   string // escaped path
	Query  string
	Header http.Header
	Body   string
}

type recorder struct {
	mu   sync.Mutex
	reqs []recorded
}

func (r *recorder) add(req *http.Request) {
	b, _ := io.ReadAll(req.Body)
	r.mu.Lock()
	defer r.mu.Unlock()
	r.reqs = append(r.reqs, recorded{
		Method: req.Method,
		Path:   req.URL.EscapedPath(),
		Query:  req.URL.RawQuery,
		Header: req.Header.Clone(),
		Body:   string(b),
	})
}

func (r *recorder) all() []recorded {
	r.mu.Lock()
	defer r.mu.Unlock()
	return append([]recorded(nil), r.reqs...)
}

func (r *recorder) last(t *testing.T) recorded {
	t.Helper()
	all := r.all()
	if len(all) == 0 {
		t.Fatal("no requests recorded")
	}
	return all[len(all)-1]
}

func newServer(t *testing.T, h http.HandlerFunc) (*httptest.Server, *recorder) {
	t.Helper()
	rec := &recorder{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec.add(r)
		h(w, r)
	}))
	t.Cleanup(srv.Close)
	return srv, rec
}

func newClient(t *testing.T, srv *httptest.Server, opts ...aveecare.Option) *aveecare.Client {
	t.Helper()
	all := append([]aveecare.Option{aveecare.WithAPIKey(testKey), aveecare.WithBaseURL(srv.URL)}, opts...)
	c, err := aveecare.NewClient(all...)
	if err != nil {
		t.Fatal(err)
	}
	return c
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func patientJSON(id string) map[string]any {
	return map[string]any{"object": "patient", "id": id, "firstName": "Margaret", "lastName": "Alvarez", "email": nil, "active": true}
}

func errorBody(status int, typ, code, msg string) map[string]any {
	return map[string]any{"error": map[string]any{"type": typ, "code": code, "message": msg, "request_id": "req_body_1"}}
}

func TestHeadersAndMe(t *testing.T) {
	srv, rec := newServer(t, func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, 200, map[string]any{
			"object": "me", "id": "Xk3mP9qR2tLw", "name": "Payroll export", "scopes": []string{"read", "write"},
			"role": "RootUser", "email": "avery@example.com", "expiresAt": "2027-10-05T00:00:00.000Z",
			"organization": map[string]any{"id": "org_1", "name": "Sunrise Home Care"},
		})
	})
	c := newClient(t, srv)
	me, err := c.Me.Retrieve(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if me.Organization.Name == nil || *me.Organization.Name != "Sunrise Home Care" || me.Name == nil || *me.Name != "Payroll export" {
		t.Fatalf("unexpected me: %+v", me)
	}
	if len(me.Scopes) != 2 || me.Scopes[1] != aveecare.MeScopeWrite || me.Role != aveecare.MeRoleRootUser {
		t.Fatalf("enums did not decode: %+v", me)
	}
	req := rec.last(t)
	if req.Method != "GET" || req.Path != "/v1/me" {
		t.Fatalf("unexpected request %s %s", req.Method, req.Path)
	}
	if got := req.Header.Get("Authorization"); got != "Bearer "+testKey {
		t.Fatalf("Authorization = %q", got)
	}
	if got := req.Header.Get("User-Agent"); got != "aveecare-go/"+aveecare.Version {
		t.Fatalf("User-Agent = %q", got)
	}
	if got := req.Header.Get("Accept"); got != "application/json" {
		t.Fatalf("Accept = %q", got)
	}
	if req.Header.Get("Idempotency-Key") != "" {
		t.Fatal("GET must not carry an Idempotency-Key")
	}
	if !regexp.MustCompile(`^\d+\.\d+\.\d+$`).MatchString(aveecare.Version) {
		t.Fatalf("Version %q is not semver", aveecare.Version)
	}
}

func TestListSendsWireParamsAndDecodesModels(t *testing.T) {
	srv, rec := newServer(t, func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, 200, map[string]any{"object": "list", "data": []any{patientJSON("pat_1")}, "next_cursor": nil, "has_more": false})
	})
	c := newClient(t, srv)
	page, err := c.Patients.List(context.Background(), &aveecare.PatientListParams{
		LifecycleStage: aveecare.Ptr(aveecare.PatientLifecycleStageProspect),
		Active:         aveecare.Bool(true),
		Limit:          aveecare.Int(25),
	})
	if err != nil {
		t.Fatal(err)
	}
	q := rec.last(t).Query
	if q != "active=true&lifecycle_stage=Prospect&limit=25" {
		t.Fatalf("query = %q", q)
	}
	if page.Object != "list" || len(page.Data) != 1 || page.HasNextPage() {
		t.Fatalf("unexpected page: %+v", page)
	}
	p := page.Data[0]
	if p.ID != "pat_1" || p.FirstName == nil || *p.FirstName != "Margaret" || p.Email != nil || p.Active == nil || !*p.Active {
		t.Fatalf("model did not decode: %+v", p)
	}
}

func TestListWithNilParamsSendsNoQuery(t *testing.T) {
	srv, rec := newServer(t, func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, 200, map[string]any{"object": "list", "data": []any{}, "next_cursor": nil, "has_more": false})
	})
	c := newClient(t, srv)
	page, err := c.Patients.List(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if rec.last(t).Query != "" || page.Data == nil || len(page.Data) != 0 {
		t.Fatalf("query %q data %v", rec.last(t).Query, page.Data)
	}
}

func TestRequiredListParamsAreSent(t *testing.T) {
	srv, rec := newServer(t, func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, 200, map[string]any{"object": "list", "data": []any{}, "next_cursor": nil, "has_more": false})
	})
	c := newClient(t, srv)
	_, err := c.Visits.List(context.Background(), &aveecare.VisitListParams{
		StartFrom: "2026-10-01T00:00:00Z",
		StartTo:   "2026-10-08T00:00:00Z",
		Status:    aveecare.Ptr(aveecare.VisitStatusInProgress),
	})
	if err != nil {
		t.Fatal(err)
	}
	if got := rec.last(t).Query; got != "start_from=2026-10-01T00%3A00%3A00Z&start_to=2026-10-08T00%3A00%3A00Z&status=InProgress" {
		t.Fatalf("query = %q", got)
	}
}

// pagedServer serves 5 patients, two per page, keyed by cursor.
func pagedServer(t *testing.T) (*httptest.Server, *recorder) {
	return newServer(t, func(w http.ResponseWriter, r *http.Request) {
		pages := map[string][]string{"": {"pat_1", "pat_2"}, "c2": {"pat_3", "pat_4"}, "c3": {"pat_5"}}
		next := map[string]any{"": "c2", "c2": "c3", "c3": nil}
		cursor := r.URL.Query().Get("cursor")
		var data []any
		for _, id := range pages[cursor] {
			data = append(data, patientJSON(id))
		}
		writeJSON(w, 200, map[string]any{"object": "list", "data": data, "next_cursor": next[cursor], "has_more": next[cursor] != nil})
	})
}

func TestAutoPagingWalksEveryItemAndKeepsFilters(t *testing.T) {
	srv, rec := pagedServer(t)
	c := newClient(t, srv)
	it := c.Patients.ListAutoPaging(context.Background(), &aveecare.PatientListParams{Active: aveecare.Bool(true), Limit: aveecare.Int(2)})
	var ids []string
	for it.Next() {
		ids = append(ids, it.Item().ID)
	}
	if err := it.Err(); err != nil {
		t.Fatal(err)
	}
	if strings.Join(ids, ",") != "pat_1,pat_2,pat_3,pat_4,pat_5" {
		t.Fatalf("ids = %v", ids)
	}
	reqs := rec.all()
	if len(reqs) != 3 {
		t.Fatalf("expected 3 requests, got %d", len(reqs))
	}
	for i, want := range []string{"active=true&limit=2", "active=true&cursor=c2&limit=2", "active=true&cursor=c3&limit=2"} {
		if reqs[i].Query != want {
			t.Fatalf("request %d query = %q, want %q", i, reqs[i].Query, want)
		}
	}
	if it.Next() {
		t.Fatal("Next must stay false after the end")
	}
}

func TestAutoPagingIsLazy(t *testing.T) {
	srv, rec := pagedServer(t)
	c := newClient(t, srv)
	it := c.Patients.ListAutoPaging(context.Background(), nil)
	if len(rec.all()) != 0 {
		t.Fatal("creating the iterator must not send a request")
	}
	if !it.Next() || it.Item().ID != "pat_1" {
		t.Fatal("expected pat_1")
	}
	if !it.Next() || it.Item().ID != "pat_2" {
		t.Fatal("expected pat_2")
	}
	if n := len(rec.all()); n != 1 {
		t.Fatalf("the second page must not be fetched yet, %d requests", n)
	}
}

func TestAutoPagingStopsOnError(t *testing.T) {
	var calls int32
	srv, _ := newServer(t, func(w http.ResponseWriter, r *http.Request) {
		if atomic.AddInt32(&calls, 1) == 1 {
			writeJSON(w, 200, map[string]any{"object": "list", "data": []any{patientJSON("pat_1")}, "next_cursor": "c2", "has_more": true})
			return
		}
		writeJSON(w, 403, errorBody(403, "permission_error", "forbidden", "No."))
	})
	c := newClient(t, srv)
	it := c.Patients.ListAutoPaging(context.Background(), nil)
	n := 0
	for it.Next() {
		n++
	}
	var perm *aveecare.PermissionError
	if n != 1 || !errors.As(it.Err(), &perm) {
		t.Fatalf("n=%d err=%v", n, it.Err())
	}
}

func TestManualPaging(t *testing.T) {
	srv, _ := pagedServer(t)
	c := newClient(t, srv)
	page, err := c.Patients.List(context.Background(), &aveecare.PatientListParams{Limit: aveecare.Int(2)})
	if err != nil {
		t.Fatal(err)
	}
	counts := []int{len(page.Data)}
	for page.HasNextPage() {
		if page, err = page.NextPage(context.Background()); err != nil {
			t.Fatal(err)
		}
		counts = append(counts, len(page.Data))
	}
	if fmt.Sprint(counts) != "[2 2 1]" {
		t.Fatalf("counts = %v", counts)
	}
	if _, err := page.NextPage(context.Background()); !errors.Is(err, aveecare.ErrNoMorePages) {
		t.Fatalf("expected ErrNoMorePages, got %v", err)
	}
}

func TestPageAutoPagingContinuesFromAPage(t *testing.T) {
	srv, _ := pagedServer(t)
	c := newClient(t, srv)
	first, err := c.Patients.List(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	second, err := first.NextPage(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	it := second.AutoPaging(context.Background())
	var ids []string
	for it.Next() {
		ids = append(ids, it.Item().ID)
	}
	if it.Err() != nil || strings.Join(ids, ",") != "pat_3,pat_4,pat_5" {
		t.Fatalf("ids = %v, err = %v", ids, it.Err())
	}
}

func TestRetrieveUpdateDelete(t *testing.T) {
	srv, rec := newServer(t, func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "DELETE":
			writeJSON(w, 200, map[string]any{"object": "patient", "id": "pat_1", "deleted": true})
		default:
			writeJSON(w, 200, patientJSON("pat_1"))
		}
	})
	c := newClient(t, srv)
	ctx := context.Background()

	if _, err := c.Patients.Retrieve(ctx, "a/b c"); err != nil {
		t.Fatal(err)
	}
	if got := rec.last(t).Path; got != "/v1/patients/a%2Fb%20c" {
		t.Fatalf("path not escaped: %q", got)
	}

	_, err := c.Patients.Update(ctx, "pat_1", &aveecare.PatientUpdate{
		FirstName:  aveecare.String("Sdk2"),
		Email:      aveecare.Null[string](),
		GoesByName: aveecare.Set("Peggy"),
	})
	if err != nil {
		t.Fatal(err)
	}
	req := rec.last(t)
	if req.Method != "PATCH" || req.Path != "/v1/patients/pat_1" {
		t.Fatalf("unexpected request %s %s", req.Method, req.Path)
	}
	var body map[string]any
	if err := json.Unmarshal([]byte(req.Body), &body); err != nil {
		t.Fatal(err)
	}
	if len(body) != 3 || body["firstName"] != "Sdk2" || body["goesByName"] != "Peggy" {
		t.Fatalf("body = %s", req.Body)
	}
	if v, ok := body["email"]; !ok || v != nil {
		t.Fatalf("email must be an explicit null: %s", req.Body)
	}
	if _, ok := body["middleName"]; ok {
		t.Fatalf("an omitted field must not be sent: %s", req.Body)
	}
	if req.Header.Get("Idempotency-Key") != "" {
		t.Fatal("PATCH must not carry an Idempotency-Key")
	}

	deleted, err := c.Patients.Delete(ctx, "pat_1")
	if err != nil {
		t.Fatal(err)
	}
	if !deleted.Deleted || deleted.ID != "pat_1" || rec.last(t).Method != "DELETE" {
		t.Fatalf("unexpected delete: %+v", deleted)
	}
}

func TestUpdateWithNilParamsSendsEmptyObject(t *testing.T) {
	srv, rec := newServer(t, func(w http.ResponseWriter, r *http.Request) { writeJSON(w, 200, patientJSON("pat_1")) })
	c := newClient(t, srv)
	if _, err := c.Patients.Update(context.Background(), "pat_1", nil); err != nil {
		t.Fatal(err)
	}
	if got := rec.last(t).Body; got != "{}" {
		t.Fatalf("body = %q", got)
	}
}

var uuidV4 = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)

func TestCreateSendsIdempotencyKeyAutomaticallyOrTheOneGiven(t *testing.T) {
	srv, rec := newServer(t, func(w http.ResponseWriter, r *http.Request) { writeJSON(w, 201, patientJSON("pat_1")) })
	c := newClient(t, srv)
	ctx := context.Background()
	params := &aveecare.PatientCreate{FirstName: "Margaret", LastName: "Alvarez", Email: aveecare.Set("m@example.com")}

	if _, err := c.Patients.Create(ctx, params); err != nil {
		t.Fatal(err)
	}
	first := rec.last(t)
	if !uuidV4.MatchString(first.Header.Get("Idempotency-Key")) {
		t.Fatalf("Idempotency-Key = %q, want a UUID v4", first.Header.Get("Idempotency-Key"))
	}
	if first.Header.Get("Content-Type") != "application/json" {
		t.Fatalf("Content-Type = %q", first.Header.Get("Content-Type"))
	}
	if first.Body != `{"firstName":"Margaret","lastName":"Alvarez","email":"m@example.com"}` {
		t.Fatalf("body = %s", first.Body)
	}
	if _, err := c.Patients.Create(ctx, params); err != nil {
		t.Fatal(err)
	}
	if rec.last(t).Header.Get("Idempotency-Key") == first.Header.Get("Idempotency-Key") {
		t.Fatal("two creates must not share an automatic key")
	}
	if _, err := c.Patients.Create(ctx, params, aveecare.WithIdempotencyKey("import-row-1042")); err != nil {
		t.Fatal(err)
	}
	if got := rec.last(t).Header.Get("Idempotency-Key"); got != "import-row-1042" {
		t.Fatalf("Idempotency-Key = %q", got)
	}
}

func TestErrorsAreTyped(t *testing.T) {
	cases := []struct {
		status int
		code   string
		check  func(error) bool
	}{
		{400, "invalid_request", func(e error) bool { var x *aveecare.InvalidRequestError; return errors.As(e, &x) }},
		{422, "invalid_request", func(e error) bool { var x *aveecare.InvalidRequestError; return errors.As(e, &x) }},
		{401, "unauthorized", func(e error) bool { var x *aveecare.AuthenticationError; return errors.As(e, &x) }},
		{403, "insufficient_scope", func(e error) bool { var x *aveecare.PermissionError; return errors.As(e, &x) }},
		{404, "not_found", func(e error) bool { var x *aveecare.NotFoundError; return errors.As(e, &x) }},
		{409, "idempotency_conflict", func(e error) bool { var x *aveecare.ConflictError; return errors.As(e, &x) }},
		{429, "rate_limited", func(e error) bool { var x *aveecare.RateLimitError; return errors.As(e, &x) }},
		{500, "internal_error", func(e error) bool {
			var x *aveecare.APIError
			var nf *aveecare.NotFoundError
			return errors.As(e, &x) && !errors.As(e, &nf)
		}},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprint(tc.status), func(t *testing.T) {
			srv, _ := newServer(t, func(w http.ResponseWriter, r *http.Request) {
				writeJSON(w, tc.status, errorBody(tc.status, "some_error", tc.code, "Something went wrong."))
			})
			c := newClient(t, srv, aveecare.WithMaxRetries(0))
			_, err := c.Patients.Retrieve(context.Background(), "pat_1")
			if err == nil || !tc.check(err) {
				t.Fatalf("status %d produced %T: %v", tc.status, err, err)
			}
			var apiErr *aveecare.APIError
			if !errors.As(err, &apiErr) {
				t.Fatalf("every HTTP error must match *APIError, got %T", err)
			}
			if apiErr.Status != tc.status || apiErr.Code != tc.code || apiErr.Type != "some_error" || apiErr.RequestID != "req_body_1" || apiErr.Message != "Something went wrong." {
				t.Fatalf("unexpected fields: %+v", apiErr)
			}
			if !strings.Contains(err.Error(), tc.code) || !strings.Contains(err.Error(), "req_body_1") {
				t.Fatalf("Error() = %q", err.Error())
			}
		})
	}
}

func TestRequestIDFallsBackToHeader(t *testing.T) {
	srv, _ := newServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Request-Id", "req_header_9")
		w.WriteHeader(502)
		_, _ = w.Write([]byte("<html>bad gateway</html>"))
	})
	c := newClient(t, srv, aveecare.WithMaxRetries(0))
	_, err := c.Me.Retrieve(context.Background())
	var apiErr *aveecare.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("got %T", err)
	}
	if apiErr.RequestID != "req_header_9" || apiErr.Status != 502 || apiErr.Message != "Request failed with status 502" || apiErr.Code != "" {
		t.Fatalf("unexpected: %+v", apiErr)
	}
}

func TestRetries429HonoringRetryAfterWithSameIdempotencyKey(t *testing.T) {
	var calls int32
	srv, rec := newServer(t, func(w http.ResponseWriter, r *http.Request) {
		if atomic.AddInt32(&calls, 1) < 3 {
			w.Header().Set("Retry-After", "0")
			writeJSON(w, 429, errorBody(429, "rate_limit_error", "rate_limited", "Slow down."))
			return
		}
		writeJSON(w, 201, patientJSON("pat_1"))
	})
	c := newClient(t, srv)
	if _, err := c.Patients.Create(context.Background(), &aveecare.PatientCreate{FirstName: "A", LastName: "B"}); err != nil {
		t.Fatal(err)
	}
	reqs := rec.all()
	if len(reqs) != 3 {
		t.Fatalf("expected 3 attempts, got %d", len(reqs))
	}
	key := reqs[0].Header.Get("Idempotency-Key")
	for i, r := range reqs {
		if r.Header.Get("Idempotency-Key") != key || key == "" {
			t.Fatalf("attempt %d changed the Idempotency-Key", i)
		}
		if r.Body != reqs[0].Body || r.Body == "" {
			t.Fatalf("attempt %d resent a different body: %q", i, r.Body)
		}
	}
}

func TestGivesUpAfterMaxRetriesWithRetryAfter(t *testing.T) {
	var calls int32
	srv, _ := newServer(t, func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&calls, 1)
		w.Header().Set("Retry-After", "0")
		writeJSON(w, 429, errorBody(429, "rate_limit_error", "rate_limited", "Slow down."))
	})
	c := newClient(t, srv, aveecare.WithMaxRetries(1))
	_, err := c.Me.Retrieve(context.Background())
	var rl *aveecare.RateLimitError
	if !errors.As(err, &rl) {
		t.Fatalf("got %T: %v", err, err)
	}
	if calls != 2 {
		t.Fatalf("calls = %d, want 2 (one retry)", calls)
	}
	if rl.RetryAfter != 0 || rl.Status != 429 {
		t.Fatalf("unexpected: %+v", rl)
	}
}

func TestRateLimitErrorExposesRetryAfter(t *testing.T) {
	srv, _ := newServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Retry-After", "7")
		writeJSON(w, 429, errorBody(429, "rate_limit_error", "rate_limited", "Slow down."))
	})
	c := newClient(t, srv, aveecare.WithMaxRetries(0))
	_, err := c.Me.Retrieve(context.Background())
	var rl *aveecare.RateLimitError
	if !errors.As(err, &rl) || rl.RetryAfter != 7*time.Second {
		t.Fatalf("got %v", err)
	}
}

func TestRetries5xxThenSucceedsAndSkips4xx(t *testing.T) {
	var calls int32
	srv, _ := newServer(t, func(w http.ResponseWriter, r *http.Request) {
		if atomic.AddInt32(&calls, 1) == 1 {
			writeJSON(w, 503, errorBody(503, "api_error", "internal_error", "Try again."))
			return
		}
		writeJSON(w, 200, patientJSON("pat_1"))
	})
	c := newClient(t, srv, aveecare.WithMaxRetries(2))
	start := time.Now()
	if _, err := c.Patients.Retrieve(context.Background(), "pat_1"); err != nil {
		t.Fatal(err)
	}
	if calls != 2 {
		t.Fatalf("calls = %d", calls)
	}
	if time.Since(start) > 3*time.Second {
		t.Fatal("backoff for the first retry must stay under a second")
	}

	for _, status := range []int{400, 401, 403, 404, 409} {
		var n int32
		srv2, _ := newServer(t, func(w http.ResponseWriter, r *http.Request) {
			atomic.AddInt32(&n, 1)
			writeJSON(w, status, errorBody(status, "x", "y", "no"))
		})
		c2 := newClient(t, srv2, aveecare.WithMaxRetries(3))
		if _, err := c2.Me.Retrieve(context.Background()); err == nil {
			t.Fatalf("%d must be an error", status)
		}
		if n != 1 {
			t.Fatalf("status %d was retried %d times", status, n-1)
		}
	}
}

func TestNetworkErrorsAreRetriedThenReported(t *testing.T) {
	var calls int32
	srv, _ := newServer(t, func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&calls, 1)
		conn, _, err := w.(http.Hijacker).Hijack()
		if err == nil {
			conn.Close()
		}
	})
	c := newClient(t, srv, aveecare.WithMaxRetries(1))
	_, err := c.Me.Retrieve(context.Background())
	var ce *aveecare.ConnectionError
	var te *aveecare.TimeoutError
	if !errors.As(err, &ce) || errors.As(err, &te) {
		t.Fatalf("expected a ConnectionError that is not a timeout, got %T: %v", err, err)
	}
	if calls != 2 {
		t.Fatalf("calls = %d, want 2", calls)
	}
}

func TestUnreachableServerIsAConnectionError(t *testing.T) {
	srv := httptest.NewServer(http.NotFoundHandler())
	url := srv.URL
	srv.Close()
	c, err := aveecare.NewClient(aveecare.WithAPIKey(testKey), aveecare.WithBaseURL(url), aveecare.WithMaxRetries(0))
	if err != nil {
		t.Fatal(err)
	}
	_, err = c.Me.Retrieve(context.Background())
	var ce *aveecare.ConnectionError
	if !errors.As(err, &ce) {
		t.Fatalf("got %T: %v", err, err)
	}
}

func TestTimeoutsAreReportedAsTimeoutError(t *testing.T) {
	release := make(chan struct{})
	srv, _ := newServer(t, func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-release:
		case <-r.Context().Done():
		}
	})
	defer close(release)
	c := newClient(t, srv, aveecare.WithTimeout(50*time.Millisecond), aveecare.WithMaxRetries(0))
	_, err := c.Me.Retrieve(context.Background())
	var te *aveecare.TimeoutError
	var ce *aveecare.ConnectionError
	if !errors.As(err, &te) || !errors.As(err, &ce) {
		t.Fatalf("got %T: %v", err, err)
	}
	if !te.Timeout() || !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("a timeout must satisfy Timeout() and context.DeadlineExceeded: %v", err)
	}
}

func TestPerRequestTimeoutAndRetriesOverrideTheClient(t *testing.T) {
	var calls int32
	srv, _ := newServer(t, func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&calls, 1)
		<-r.Context().Done()
	})
	c := newClient(t, srv, aveecare.WithTimeout(time.Minute), aveecare.WithMaxRetries(5))
	_, err := c.Me.Retrieve(context.Background(), aveecare.WithRequestTimeout(40*time.Millisecond), aveecare.WithRequestMaxRetries(0))
	var te *aveecare.TimeoutError
	if !errors.As(err, &te) || calls != 1 {
		t.Fatalf("err %v, calls %d", err, calls)
	}
}

func TestCanceledContextStopsWithoutRetrying(t *testing.T) {
	var calls int32
	srv, _ := newServer(t, func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&calls, 1)
		<-r.Context().Done()
	})
	c := newClient(t, srv, aveecare.WithMaxRetries(3))
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()
	_, err := c.Me.Retrieve(ctx)
	var ce *aveecare.ConnectionError
	if !errors.As(err, &ce) || !errors.Is(err, context.Canceled) {
		t.Fatalf("got %T: %v", err, err)
	}
	if n := atomic.LoadInt32(&calls); n != 1 {
		t.Fatalf("a canceled request was retried: %d calls", n)
	}
}

func TestAlreadyCanceledContextSendsNothing(t *testing.T) {
	srv, rec := newServer(t, func(w http.ResponseWriter, r *http.Request) { writeJSON(w, 200, map[string]any{}) })
	c := newClient(t, srv)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := c.Me.Retrieve(ctx); !errors.Is(err, context.Canceled) {
		t.Fatalf("got %v", err)
	}
	if len(rec.all()) != 0 {
		t.Fatal("no request should be sent")
	}
}

func TestNonJSONSuccessIsAnInvalidResponseError(t *testing.T) {
	srv, _ := newServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Request-Id", "req_x")
		_, _ = w.Write([]byte("<html>not json</html>"))
	})
	c := newClient(t, srv)
	_, err := c.Me.Retrieve(context.Background())
	var ir *aveecare.InvalidResponseError
	if !errors.As(err, &ir) || ir.Status != 200 || ir.RequestID != "req_x" {
		t.Fatalf("got %T: %v", err, err)
	}
}

func TestClientConfiguration(t *testing.T) {
	t.Setenv("AVEECARE_API_KEY", "")
	t.Setenv("AVEECARE_BASE_URL", "")
	if _, err := aveecare.NewClient(); !errors.Is(err, aveecare.ErrMissingAPIKey) {
		t.Fatalf("expected ErrMissingAPIKey, got %v", err)
	}
	c, err := aveecare.NewClient(aveecare.WithAPIKey("k"))
	if err != nil || c.BaseURL() != aveecare.DefaultBaseURL {
		t.Fatalf("default base URL: %v %q", err, c.BaseURL())
	}

	srv, rec := newServer(t, func(w http.ResponseWriter, r *http.Request) { writeJSON(w, 200, map[string]any{"object": "me"}) })
	t.Setenv("AVEECARE_API_KEY", "  ak_env_key  ")
	t.Setenv("AVEECARE_BASE_URL", srv.URL+"/")
	c, err = aveecare.NewClient()
	if err != nil {
		t.Fatal(err)
	}
	if c.BaseURL() != srv.URL {
		t.Fatalf("BaseURL = %q, want trailing slash trimmed", c.BaseURL())
	}
	if _, err := c.Me.Retrieve(context.Background()); err != nil {
		t.Fatal(err)
	}
	if got := rec.last(t).Header.Get("Authorization"); got != "Bearer ak_env_key" {
		t.Fatalf("Authorization = %q", got)
	}
	// Explicit options win over the environment.
	c, err = aveecare.NewClient(aveecare.WithAPIKey("ak_opt_key"), aveecare.WithBaseURL("http://example.invalid"))
	if err != nil || c.BaseURL() != "http://example.invalid" {
		t.Fatalf("options must win: %v %q", err, c.BaseURL())
	}
}

func TestHeadersAndHTTPClientOptions(t *testing.T) {
	srv, rec := newServer(t, func(w http.ResponseWriter, r *http.Request) { writeJSON(w, 200, map[string]any{"object": "me"}) })
	var used int32
	hc := &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
		atomic.AddInt32(&used, 1)
		return http.DefaultTransport.RoundTrip(r)
	})}
	c := newClient(t, srv, aveecare.WithHTTPClient(hc), aveecare.WithHeader("X-Trace", "abc"))
	if _, err := c.Me.Retrieve(context.Background(), aveecare.WithRequestHeader("X-Once", "1")); err != nil {
		t.Fatal(err)
	}
	h := rec.last(t).Header
	if h.Get("X-Trace") != "abc" || h.Get("X-Once") != "1" || atomic.LoadInt32(&used) != 1 {
		t.Fatalf("headers %v, custom client used %d times", h, used)
	}
	if _, err := c.Me.Retrieve(context.Background()); err != nil {
		t.Fatal(err)
	}
	if rec.last(t).Header.Get("X-Once") != "" {
		t.Fatal("a per-request header must not leak into the next request")
	}
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func TestUnknownEnumValuesAndFieldsDecode(t *testing.T) {
	srv, _ := newServer(t, func(w http.ResponseWriter, r *http.Request) {
		p := patientJSON("pat_1")
		p["lifecycleStage"] = "Alumni"
		p["somethingNew"] = map[string]any{"a": 1}
		writeJSON(w, 200, p)
	})
	c := newClient(t, srv)
	p, err := c.Patients.Retrieve(context.Background(), "pat_1")
	if err != nil {
		t.Fatal(err)
	}
	if p.LifecycleStage == nil || *p.LifecycleStage != aveecare.PatientLifecycleStage("Alumni") {
		t.Fatalf("an enum value from a newer API must survive: %+v", p.LifecycleStage)
	}
}

func TestNullableStates(t *testing.T) {
	var omitted aveecare.Nullable[string]
	set := aveecare.Set("x")
	null := aveecare.Null[string]()
	if !omitted.IsOmitted() || omitted.IsNull() || omitted.IsSet() {
		t.Fatal("zero value must be omitted")
	}
	if v, ok := set.Value(); !ok || v != "x" || !set.IsSet() {
		t.Fatal("Set must hold its value")
	}
	if !null.IsNull() {
		t.Fatal("Null must be null")
	}
	b, _ := json.Marshal(struct {
		A aveecare.Nullable[string]
		B aveecare.Nullable[string]
	}{set, null})
	if string(b) != `{"A":"x","B":null}` {
		t.Fatalf("json = %s", b)
	}
	var back struct{ A, B, C aveecare.Nullable[string] }
	if err := json.Unmarshal([]byte(`{"A":"x","B":null}`), &back); err != nil {
		t.Fatal(err)
	}
	if !back.A.IsSet() || !back.B.IsNull() || !back.C.IsOmitted() {
		t.Fatalf("round trip: %+v", back)
	}
}
