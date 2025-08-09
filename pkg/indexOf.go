package pkg

import "slices"

// IndexOf returns the first index at which a given element can be found in the slice.
// Similar to JavaScript's Array.indexOf().
//
// Parameters:
//   - in: The input slice to search
//   - target: The value to locate
//
// Returns:
//   - The first index of the element in the slice if found; otherwise, -1
func IndexOf[T comparable](in []T, target T) int {
	return slices.Index(in, target)
}

// IndexOfFunc returns the first index at which an element satisfies the provided testing function.
//
// Parameters:
//   - in: The input slice to search
//   - pred: A function that tests each element and returns true when the desired element is found
//
// Returns:
//   - The first index of an element that satisfies the predicate if found; otherwise, -1
func IndexOfFunc[T any](in []T, pred func(T) bool) int {
	return slices.IndexFunc(in, pred)
}

// LastIndexOf returns the last index at which a given element can be found in the slice.
// Similar to JavaScript's Array.lastIndexOf().
//
// Parameters:
//   - in: The input slice to search
//   - target: The value to locate
//
// Returns:
//   - The last index of the element in the slice if found; otherwise, -1
func LastIndexOf[T comparable](in []T, target T) int {
	for i := len(in) - 1; i >= 0; i-- {
		if in[i] == target {
			return i
		}
	}
	return -1
}

// LastIndexOfFunc returns the last index for which the predicate returns true.
//
// Parameters:
//   - in: The input slice to search
//   - pred: A function that tests each element and returns true when the desired element is
//
// Returns:
//   - The last index of an element that satisfies the predicate if found; otherwise, -
//     -1
//
// Note: This function iterates from the end of the slice to the beginning, returning the
// first index where the predicate returns true. If no such element is found, it returns -1.
//
// Example usage:
//
//	index := LastIndexOfFunc([]int{1, 2, 3, 2, 1}, func(x int) bool { return x == 2 })
//	fmt.Println(index) // Output: 3
//
// This function is useful for finding the last occurrence of an element that meets a specific condition.
func LastIndexOfFunc[T any](in []T, pred func(T) bool) int {
	for i := len(in) - 1; i >= 0; i-- {
		if pred(in[i]) {
			return i
		}
	}
	return -1
}
