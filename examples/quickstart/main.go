// Command quickstart lists a few patients.
//
//	AVEECARE_API_KEY=ak_... go run ./examples/quickstart
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aveecare/aveecare-go"
)

func main() {
	// With no options, NewClient reads AVEECARE_API_KEY from the environment.
	client, err := aveecare.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	page, err := client.Patients.List(context.Background(), &aveecare.PatientListParams{
		Active: aveecare.Bool(true),
		Limit:  aveecare.Int(25),
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, patient := range page.Data {
		fmt.Println(patient.ID)
	}
}
