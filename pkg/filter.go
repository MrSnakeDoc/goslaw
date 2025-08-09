package pkg

import "slices"

// Filter creates a new slice with all elements that pass the test implemented by the provided function.
// Similar to JavaScript's Array.filter().
//
// Parameters:
//   - in: The input slice to filter
//   - keep: A function that tests each element and returns true to keep it in the result
//
// Returns:
//   - A new slice containing only the elements that passed the test
func Filter[T any](in []T, keep func(T) bool) []T {
	out := make([]T, 0, len(in)) // cap max, len 0

	for _, v := range in {
		if keep(v) {
			out = append(out, v)
		}
	}

	if len(out) >= cap(out)*3/4 { // optimize to avoid over-allocating small filtered results
		return out
	}

	return slices.Clone(out)
}
