package pkg

import "slices"

// With returns a new slice with the element at the specified index replaced with the provided value.
// Similar to JavaScript's Array.with().
//
// Parameters:
//   - in: The input slice
//   - idx: The index of the element to replace
//   - val: The new value to replace with
//
// Returns:
//   - A new slice with the element at the specified index replaced
//
// Panics:
//   - If the index is out of range
func With[T any](in []T, idx int, val T) []T {
	if idx < 0 || idx >= len(in) {
		panic("index out of range")
	}
	out := slices.Clone(in)
	out[idx] = val
	return out
}
