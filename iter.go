//go:build go1.23

package aveecare

import "iter"

// All returns the remaining items as an iter.Seq2 for range-over-func (Go 1.23
// and newer). A failure is yielded once as (zero, err) and ends the sequence.
//
//	for patient, err := range client.Patients.ListAutoPaging(ctx, params).All() {
//		if err != nil {
//			return err
//		}
//		fmt.Println(patient.ID)
//	}
func (it *Iterator[T]) All() iter.Seq2[T, error] {
	return func(yield func(T, error) bool) {
		for it.Next() {
			if !yield(it.Item(), nil) {
				return
			}
		}
		if err := it.Err(); err != nil {
			var zero T
			yield(zero, err)
		}
	}
}
