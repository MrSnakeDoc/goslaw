package pkg

// Slice returns a shallow copy of a portion of a slice into a new slice object.
// Similar to JavaScript's Array.slice().
//
// Parameters:
//   - in: The input slice
//   - start: Zero-based index at which to start extraction (inclusive)
//     If negative, it is treated as an offset from the end of the slice
//   - end: Zero-based index at which to end extraction (exclusive)
//     If negative, it is treated as an offset from the end of the slice
//
// Returns:
//   - A new slice containing the extracted elements
func Slice[T any](in []T, start, end int) []T {
	n := len(in)

	if start < 0 {
		start += n
	}
	if start < 0 {
		start = 0
	}
	if start > n {
		start = n
	}

	if end < 0 {
		end += n
	}
	if end < start || end > n {
		end = n
	}
	if start >= end {
		return in[:0]
	}
	return in[start:end]
}
