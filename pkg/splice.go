package pkg

// Splice removes a specified number of elements from a slice starting at a given index and inserts new elements in their place.
// Similar to JavaScript's Array.splice().
//
// Parameters:
//   - in: The input slice from which elements will be removed and new elements will be inserted.
//   - start: The index at which to start removing elements. If negative, it is treated as an offset from the end of the slice.
//   - deleteCount: The number of elements to remove from the slice. If negative, it is treated as zero, meaning no elements will be removed.
//   - items: A variadic parameter that accepts new elements to be inserted into the slice at the specified index.
//
// Returns:
//   - A new slice containing the elements from the original slice with the specified elements removed and new elements inserted.
//   - If `start` is greater than the length of the slice, it will behave as if `start` is equal to the length of the slice, meaning no elements will be removed.
func Splice[T any](in []T, start, deleteCount int, items ...T) []T {
	if start < 0 {
		start = 0
	}
	if deleteCount < 0 {
		deleteCount = 0
	}
	if start > len(in) {
		start = len(in)
	}

	end := start + deleteCount
	if end > len(in) {
		end = len(in)
	}

	out := make([]T, 0, len(in)-deleteCount+len(items))
	out = append(out, in[:start]...)
	out = append(out, items...)
	out = append(out, in[end:]...)

	return out
}
