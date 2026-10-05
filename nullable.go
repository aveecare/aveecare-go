package aveecare

import (
	"bytes"
	"encoding/json"
)

type nullableState uint8

const (
	nullableOmitted nullableState = iota
	nullableSet
	nullableNull
)

// Nullable is an optional request field the API also accepts as null. It has
// three states, so an update can leave a field alone, change it, or clear it:
//
//	aveecare.PatientUpdate{
//		FirstName: aveecare.Set("Margaret"), // change it
//		Email:     aveecare.Null[string](),  // send null, clearing it
//		// MiddleName is left at its zero value: not sent at all.
//	}
//
// The zero value omits the field from the request.
type Nullable[T any] struct {
	value T
	state nullableState
}

// Set returns a Nullable that sends v.
func Set[T any](v T) Nullable[T] { return Nullable[T]{value: v, state: nullableSet} }

// Null returns a Nullable that sends an explicit JSON null.
func Null[T any]() Nullable[T] { return Nullable[T]{state: nullableNull} }

// IsOmitted reports whether the field is left out of the request.
func (n Nullable[T]) IsOmitted() bool { return n.state == nullableOmitted }

// IsNull reports whether the field is sent as null.
func (n Nullable[T]) IsNull() bool { return n.state == nullableNull }

// IsSet reports whether the field is sent with a value.
func (n Nullable[T]) IsSet() bool { return n.state == nullableSet }

// Value returns the value and true when the field is set.
func (n Nullable[T]) Value() (T, bool) { return n.value, n.state == nullableSet }

// MarshalJSON implements json.Marshaler. An omitted Nullable encodes as null
// when marshaled on its own; the generated request types skip omitted fields.
func (n Nullable[T]) MarshalJSON() ([]byte, error) {
	if n.state != nullableSet {
		return []byte("null"), nil
	}
	return json.Marshal(n.value)
}

// UnmarshalJSON implements json.Unmarshaler.
func (n *Nullable[T]) UnmarshalJSON(data []byte) error {
	if bytes.Equal(bytes.TrimSpace(data), []byte("null")) {
		*n = Nullable[T]{state: nullableNull}
		return nil
	}
	var v T
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	*n = Nullable[T]{value: v, state: nullableSet}
	return nil
}

// objectWriter builds a JSON object field by field, in the order added. The
// generated request types use it to leave out fields that were not given.
type objectWriter struct {
	buf bytes.Buffer
	n   int
	err error
}

func (w *objectWriter) add(name string, value any) {
	if w.err != nil {
		return
	}
	key, err := json.Marshal(name)
	if err != nil {
		w.err = err
		return
	}
	val, err := json.Marshal(value)
	if err != nil {
		w.err = err
		return
	}
	if w.n > 0 {
		w.buf.WriteByte(',')
	}
	w.buf.Write(key)
	w.buf.WriteByte(':')
	w.buf.Write(val)
	w.n++
}

func (w *objectWriter) bytes() ([]byte, error) {
	if w.err != nil {
		return nil, w.err
	}
	out := make([]byte, 0, w.buf.Len()+2)
	out = append(out, '{')
	out = append(out, w.buf.Bytes()...)
	return append(out, '}'), nil
}
