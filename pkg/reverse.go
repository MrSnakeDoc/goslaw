package pkg

// Reverse returns a new slice with the elements in reversed order.
// Note: Unlike JavaScript's Array.reverse(), this does not modify the original slice in-place.
//
// Parameters:
//   - in: The input slice to reverse
//
// Returns:
//   - A new slice with elements in reversed order
func Reverse[T any](in []T) []T {
	if len(in) == 0 {
		return in
	}
	out := make([]T, len(in))
	for i, v := range in {
		out[len(in)-1-i] = v
	}
	return out
}
