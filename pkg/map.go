package pkg

// Map applies a transformation function to each element in a slice and returns a new slice
// containing the results. Similar to JavaScript's Array.map().
//
// Parameters:
//   - in: The input slice to transform
//   - f: A function that transforms each element from type T to type R
//
// Returns:
//   - A new slice of type R containing the transformed elements
func Map[T any, R any](in []T, f func(T) R) []R {
	out := make([]R, len(in))
	for i, v := range in {
		out[i] = f(v)
	}
	return out
}
