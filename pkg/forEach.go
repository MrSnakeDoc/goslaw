package pkg

// ForEach executes a provided function once for each slice element.
// Similar to JavaScript's Array.forEach().
//
// Parameters:
//   - in: The input slice to iterate over
//   - fn: A function to execute for each element
//
// Returns:
//   - None (void)
func ForEach[T any](in []T, fn func(T)) {
	for _, v := range in {
		fn(v)
	}
}
