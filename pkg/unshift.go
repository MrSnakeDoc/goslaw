package pkg

// Unshift adds one element to the beginning of a slice and returns the new slice.
// Similar to JavaScript's Array.unshift().
//
// Parameters:
//   - in: The input slice
//   - item: The element to add to the beginning of the slice
//
// Returns:
//   - A new slice with the element added at the beginning
func Unshift[T any](in []T, item T) []T {
	n := len(in)
	if n == 0 {
		return []T{item}
	}

	if cap(in) > n {
		in = in[:n+1]
		copy(in[1:], in[:n])
		in[0] = item
		return in
	}

	out := make([]T, n+1)
	out[0] = item
	copy(out[1:], in)
	return out
}
