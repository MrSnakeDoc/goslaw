package pkg

// FlatMap first maps each element using a mapping function, then flattens the result into a new slice.
// Similar to JavaScript's Array.flatMap().
//
// Parameters:
//   - in: The input slice to transform
//   - fn: A function that maps each element to a slice of elements of type R
//
// Returns:
//   - A new flattened slice of type R containing all mapped elements
func FlatMap[T any, R any](in []T, fn func(T) []R) []R {
	if len(in) == 0 {
		return nil
	}
	var out []R
	for _, v := range in {
		out = append(out, fn(v)...)
	}

	if len(out) == 0 {
		return []R{}
	}

	return out
}
