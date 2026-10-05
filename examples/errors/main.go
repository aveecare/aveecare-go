// Command errors shows how to branch on typed errors.
//
//	AVEECARE_API_KEY=ak_... go run ./examples/errors
package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/aveecare/aveecare-go"
)

func main() {
	client, err := aveecare.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.Patients.Retrieve(context.Background(), "pat_missing")

	var notFound *aveecare.NotFoundError
	var rateLimited *aveecare.RateLimitError
	var apiErr *aveecare.APIError
	switch {
	case err == nil:
		fmt.Println("found")
	case errors.As(err, &notFound):
		fmt.Println("no such patient, request", notFound.RequestID)
	case errors.As(err, &rateLimited):
		fmt.Println("slow down, retry after", rateLimited.RetryAfter)
	case errors.As(err, &apiErr):
		fmt.Println(apiErr.Status, apiErr.Code, apiErr.Message)
	default:
		log.Fatal(err)
	}
}
