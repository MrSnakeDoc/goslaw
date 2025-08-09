package pkg

import "slices"

// ToReversed returns a new slice with the elements in reversed order.
// Unlike Reverse, this method does not mutate the original slice.
// Similar to JavaScript's Array.toReversed().
//
// Parameters:
//   - in: The input slice to reverse
//
// Returns:
//   - A new slice with elements in reversed order
func ToReversed[T any](in []T) []T {
	out := slices.Clone(in)
	slices.Reverse(out)
	return out
}
