package pkg

// Every tests whether all elements in the slice pass the test implemented by the provided function.
// Similar to JavaScript's Array.every().
//
// Parameters:
//   - in: The input slice to test
//   - pred: A function that tests each element and returns a boolean
//
// Returns:
//   - true if all elements pass the test; otherwise, false
func Every[T any](in []T, pred func(T) bool) bool {
	for _, v := range in {
		if !pred(v) {
			return false
		}
	}
	return true
}

// EveryIndex tests whether all elements in the slice pass the test implemented by the provided function.
// It also returns the index of the first element that fails the test.
// Similar to JavaScript's Array.every().
//
// Parameters:
//   - in: The input slice to test
//   - pred: A function that tests each element and returns a boolean
//
// Returns:
//   - true if all elements pass the test; otherwise, false
//   - The index of the first element that fails the test, or -1 if all pass
func EveryIndex[T any](in []T, pred func(T) bool) (bool, int) {
	for i, v := range in {
		if !pred(v) {
			return false, i
		}
	}
	return true, -1
}
