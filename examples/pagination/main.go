// Command pagination walks every patient across all pages.
//
//	AVEECARE_API_KEY=ak_... go run ./examples/pagination
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aveecare/aveecare-go"
)

func main() {
	client, err := aveecare.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	// Automatic: pages are fetched lazily as the iterator advances.
	it := client.Patients.ListAutoPaging(ctx, &aveecare.PatientListParams{Limit: aveecare.Int(100)})
	count := 0
	for it.Next() {
		fmt.Println(it.Item().ID)
		count++
	}
	if err := it.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("patients:", count)

	// By hand: one page at a time.
	page, err := client.Patients.List(ctx, &aveecare.PatientListParams{Limit: aveecare.Int(50)})
	for err == nil {
		fmt.Println("page of", len(page.Data))
		if !page.HasNextPage() {
			return
		}
		page, err = page.NextPage(ctx)
	}
	log.Fatal(err)
}
