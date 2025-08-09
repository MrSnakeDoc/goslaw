package pkg

import (
	"sort"
)

// Sort sorts the elements of a slice in place.
// Similar to JavaScript's Array.sort().
//
// Parameters:
//   - in: The input slice to sort
//   - less: A function that returns true if element a should come before element b
//
// Returns:
//   - None (modifies the slice in-place)
func Sort[T any](in []T, less func(a, b T) bool) {
	if len(in) < 2 {
		return
	}
	sort.Slice(in, func(i, j int) bool { return less(in[i], in[j]) })
}
