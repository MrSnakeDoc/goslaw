package pkg

// Entries returns a slice of key-value pairs from a map.
// Similar to JavaScript's Object.entries().
//
// Parameters:
//   - m: The map from which to extract key-value pairs
//
// Returns:
//   - A slice of struct objects containing Key and Value fields
func Entries[K comparable, V any](m map[K]V) []struct {
	Key   K
	Value V
} {
	if len(m) == 0 {
		return nil
	}
	out := make([]struct {
		Key   K
		Value V
	}, 0, len(m))
	for k, v := range m {
		out = append(out, struct {
			Key   K
			Value V
		}{k, v})
	}
	return out
}
