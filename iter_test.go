//go:build go1.23

package aveecare_test

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/aveecare/aveecare-go"
)

func TestRangeOverFuncIteratesEveryItem(t *testing.T) {
	srv, _ := pagedServer(t)
	c := newClient(t, srv)
	var ids []string
	for patient, err := range c.Patients.ListAutoPaging(context.Background(), nil).All() {
		if err != nil {
			t.Fatal(err)
		}
		ids = append(ids, patient.ID)
	}
	if strings.Join(ids, ",") != "pat_1,pat_2,pat_3,pat_4,pat_5" {
		t.Fatalf("ids = %v", ids)
	}
}

func TestRangeOverFuncCanBreakEarly(t *testing.T) {
	srv, rec := pagedServer(t)
	c := newClient(t, srv)
	n := 0
	for _, err := range c.Patients.ListAutoPaging(context.Background(), nil).All() {
		if err != nil {
			t.Fatal(err)
		}
		n++
		if n == 2 {
			break
		}
	}
	if len(rec.all()) != 1 {
		t.Fatalf("breaking after the first page must not fetch more, got %d requests", len(rec.all()))
	}
}

func TestRangeOverFuncYieldsTheErrorOnce(t *testing.T) {
	srv, _ := newServer(t, func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, 401, errorBody(401, "authentication_error", "unauthorized", "Bad key."))
	})
	c := newClient(t, srv, aveecare.WithMaxRetries(0))
	var errs int
	for _, err := range c.Patients.ListAutoPaging(context.Background(), nil).All() {
		var auth *aveecare.AuthenticationError
		if !errors.As(err, &auth) {
			t.Fatalf("got %v", err)
		}
		errs++
	}
	if errs != 1 {
		t.Fatalf("error yielded %d times", errs)
	}
}
