package pkg

import "slices"

// Flat creates a new slice with all sub-slice elements concatenated into it recursively.
// Similar to JavaScript's Array.flat().
//
// Parameters:
//   - chunks: A slice of slices to flatten
//
// Returns:
//   - A new flattened slice containing all elements from the sub-slices
func Flat[T any](chunks [][]T) []T {
	var out []T
	for _, c := range chunks {
		out = append(out, c...)
	}

	if len(out) >= cap(out)*3/4 {
		return out
	}

	return slices.Clone(out)
}
