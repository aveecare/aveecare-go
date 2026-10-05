package aveecare_test

// Live tests against a real AveeCare API.
//
// They are skipped unless both AVEECARE_API_KEY and AVEECARE_BASE_URL are set,
// and they create, update and delete one test patient. Use a sandbox or local
// API, never production data:
//
//	AVEECARE_API_KEY=ak_... AVEECARE_BASE_URL=https://... go test -run Live -v

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/aveecare/aveecare-go"
)

func liveClient(t *testing.T) *aveecare.Client {
	t.Helper()
	if os.Getenv("AVEECARE_API_KEY") == "" || os.Getenv("AVEECARE_BASE_URL") == "" {
		t.Skip("set AVEECARE_API_KEY and AVEECARE_BASE_URL to run live tests")
	}
	c, err := aveecare.NewClient()
	if err != nil {
		t.Fatal(err)
	}
	return c
}

func TestLiveMe(t *testing.T) {
	c := liveClient(t)
	me, err := c.Me.Retrieve(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if me.ID == "" {
		t.Fatal("empty key id")
	}
}

func TestLiveListPageTwiceAndRetrieve(t *testing.T) {
	c := liveClient(t)
	ctx := context.Background()
	first, err := c.Patients.List(ctx, &aveecare.PatientListParams{Limit: aveecare.Int(1)})
	if err != nil {
		t.Fatal(err)
	}
	if first.Object != "list" {
		t.Fatalf("object = %q", first.Object)
	}
	if first.HasNextPage() {
		second, err := first.NextPage(ctx)
		if err != nil {
			t.Fatal(err)
		}
		if len(first.Data) > 0 && len(second.Data) > 0 && first.Data[0].ID == second.Data[0].ID {
			t.Fatal("the next page repeated an item")
		}
	}
	if len(first.Data) > 0 {
		got, err := c.Patients.Retrieve(ctx, first.Data[0].ID)
		if err != nil || got.ID != first.Data[0].ID {
			t.Fatalf("retrieve: %v %+v", err, got)
		}
	}
}

func TestLiveCreateIdempotentUpdateDelete(t *testing.T) {
	c := liveClient(t)
	ctx := context.Background()
	key := fmt.Sprintf("live-%d", time.Now().UnixNano())
	last := "Livetest " + key
	created, err := c.Patients.Create(ctx, &aveecare.PatientCreate{FirstName: "Sdk", LastName: last}, aveecare.WithIdempotencyKey(key))
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if _, err := c.Patients.Delete(ctx, created.ID); err != nil {
			t.Errorf("cleanup delete: %v", err)
		}
		_, err := c.Patients.Retrieve(ctx, created.ID)
		var nf *aveecare.NotFoundError
		if !errors.As(err, &nf) {
			t.Errorf("a deleted patient must be 404, got %v", err)
		}
	}()
	replay, err := c.Patients.Create(ctx, &aveecare.PatientCreate{FirstName: "Sdk", LastName: last}, aveecare.WithIdempotencyKey(key))
	if err != nil {
		t.Fatal(err)
	}
	if replay.ID != created.ID {
		t.Fatalf("the same Idempotency-Key must return the same patient: %s vs %s", replay.ID, created.ID)
	}
	updated, err := c.Patients.Update(ctx, created.ID, &aveecare.PatientUpdate{FirstName: aveecare.String("Sdk2")})
	if err != nil {
		t.Fatal(err)
	}
	if updated.FirstName == nil || *updated.FirstName != "Sdk2" {
		t.Fatalf("update did not apply: %+v", updated.FirstName)
	}
}
