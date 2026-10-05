package aveecare

// Ptr returns a pointer to v. Optional request fields and query parameters are
// pointers so that "not given" can be told apart from the zero value:
//
//	&aveecare.PatientListParams{Limit: aveecare.Ptr[int64](25)}
func Ptr[T any](v T) *T { return &v }

// String returns a pointer to s.
func String(s string) *string { return &s }

// Int returns a pointer to n.
func Int(n int64) *int64 { return &n }

// Float returns a pointer to f.
func Float(f float64) *float64 { return &f }

// Bool returns a pointer to b.
func Bool(b bool) *bool { return &b }
