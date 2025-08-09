package pkg

// Keys returns a slice containing the keys of the map.
// Similar to JavaScript's Object.keys().
//
// Parameters:
//   - m: The map from which to extract keys
//
// Returns:
//   - A slice containing all the keys from the map
func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}
