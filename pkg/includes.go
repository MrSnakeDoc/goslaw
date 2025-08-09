package pkg

// Includes determines whether a slice includes a certain value among its entries.
// Similar to JavaScript's Array.includes().
//
// Parameters:
//   - in: The input slice to search
//   - target: The value to search for
//
// Returns:
//   - true if the target value is found; otherwise, false
func Includes[T comparable](in []T, target T) bool {
	for _, v := range in {
		if v == target {
			return true
		}
	}
	return false
}
