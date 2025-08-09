package pkg

// Values returns a slice containing the values of the map.
// Similar to JavaScript's Object.values().
//
// Parameters:
//   - m: The map from which to extract values
//
// Returns:
//   - A slice containing all the values from the map
func Values[K comparable, V any](m map[K]V) []V {
	out := make([]V, len(m))
	i := 0
	for _, v := range m {
		out[i] = v
		i++
	}
	return out
}
