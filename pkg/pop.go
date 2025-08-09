package pkg

// Pop removes the last element from a slice and returns that element and the modified slice.
// Similar to JavaScript's Array.pop().
//
// Parameters:
//   - in: The input slice
//
// Returns:
//   - The removed element and the modified slice
func Pop[T any](in []T) (T, []T) {
	if len(in) == 0 {
		var zero T
		return zero, in[:0]
	}
	last := in[len(in)-1]
	rest := in[: len(in)-1 : len(in)-1]
	return last, rest
}

// PopFree removes the last element from a slice and returns that element and the modified slice.
// It additionally frees memory if the slice capacity is significantly larger than its length.
//
// Parameters:
//   - in: The input slice
//
// Returns:
//   - The removed element and the modified slice with potentially reduced capacity
func PopFree[T any](in []T) (T, []T) {
	if len(in) == 0 {
		var zero T
		return zero, in[:0]
	}

	last := in[len(in)-1]
	rest := in[:len(in)-1]

	if cap(in) > len(rest)*2 {
		tmp := make([]T, len(rest))
		copy(tmp, rest)
		rest = tmp
	}
	return last, rest
}
