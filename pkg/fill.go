package pkg

// Fill creates a new slice with all elements set to the specified value.
// Similar to JavaScript's Array.fill().
//
// Parameters:
//   - in: The input slice to fill
//   - value: The value to fill the slice with
//
// Returns:
//   - A new slice with all elements set to the specified value
func Fill[T any](template []T, value T) []T {
	out := make([]T, len(template))
	for i := range out {
		out[i] = value
	}
	return out
}
