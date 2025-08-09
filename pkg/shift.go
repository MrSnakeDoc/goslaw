package pkg

// Shift removes the first element from a slice and returns that element and the modified slice.
// Similar to JavaScript's Array.shift().
//
// Parameters:
//   - in: The input slice
//
// Returns:
//   - The removed element and the modified slice
func Shift[T any](in []T) (T, []T) {
	if len(in) == 0 {
		var zero T
		return zero, nil
	}
	return in[0], in[1:]
}
