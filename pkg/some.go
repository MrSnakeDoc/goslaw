package pkg

// Some tests whether at least one element in the slice passes the test implemented by the provided function.
// Similar to JavaScript's Array.some().
//
// Parameters:
//   - in: The input slice to test
//   - pred: A function that tests each element and returns a boolean
//
// Returns:
//   - true if any element passes the test; otherwise, false
func Some[T any](in []T, pred func(T) bool) bool {
	for _, v := range in {
		if pred(v) {
			return true
		}
	}
	return false
}
