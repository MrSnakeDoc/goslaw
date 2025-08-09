package pkg

// Find returns the first element in the provided slice that satisfies the provided testing function.
// Similar to JavaScript's Array.find().
//
// Parameters:
//   - in: The input slice to search
//   - pred: A function that tests each element and returns true when the desired element is found
//
// Returns:
//   - The found element and true if an element passes the test; otherwise, zero value and false
func Find[T any](in []T, pred func(T) bool) (T, bool) {
	for _, v := range in {
		if pred(v) {
			return v, true
		}
	}
	var zero T
	return zero, false
}

// FindIndex returns the index of the first element in the provided slice that satisfies the provided testing function.
// Similar to JavaScript's Array.findIndex().
//
// Parameters:
//   - in: The input slice to search
//   - pred: A function that tests each element and returns true when the desired element is found
//
// Returns:
//   - The index of the first element that satisfies the predicate if found; otherwise, -1
func FindIndex[T any](in []T, pred func(T) bool) int {
	for i, v := range in {
		if pred(v) {
			return i
		}
	}
	return -1
}
