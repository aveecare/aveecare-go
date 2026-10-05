// Package aveecare is the official Go library for the AveeCare API
// (https://www.aveecare.com/docs/api).
//
// Read and write your home care agency's patients, caregivers, visits, care
// documentation, billing records and messages from your own servers, with typed
// models, automatic pagination, retries and idempotency built in.
//
//	client, err := aveecare.NewClient(aveecare.WithAPIKey(os.Getenv("AVEECARE_API_KEY")))
//	if err != nil {
//		log.Fatal(err)
//	}
//	page, err := client.Patients.List(ctx, &aveecare.PatientListParams{Limit: aveecare.Int(25)})
//
// Every method takes a context.Context first. Errors are typed: use errors.As
// with *NotFoundError, *RateLimitError and the other error types, or with
// *APIError to match any non-2xx response.
//
// Keep API keys on your server; never put one in browser or mobile code. See
// https://www.aveecare.com/docs/api/authentication.
package aveecare
