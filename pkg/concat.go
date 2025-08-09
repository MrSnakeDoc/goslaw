package pkg

// Concat concatenates multiple slices into a single slice.
// Parameters:
//   - slices: A variadic parameter that accepts multiple slices of any type
//
// Returns:
//   - A new slice containing all elements from the provided slices
func Concat[T any](slices ...[]T) []T {
	var totalLen int
	for _, s := range slices {
		totalLen += len(s)
	}

	out := make([]T, 0, totalLen)
	for _, s := range slices {
		out = append(out, s...)
	}
	return out
}
