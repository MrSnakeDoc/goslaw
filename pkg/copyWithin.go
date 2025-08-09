package pkg

// CopyWithin copies part of a slice to another location in the same slice without modifying its length.
// Similar to JavaScript's Array.copyWithin().
//
// Parameters:
//   - in: The slice to modify in place
//   - target: Index at which to copy the elements to
//   - start: Index at which to start copying from
//   - end (optional): Index to stop copying (exclusive). Defaults to len(in) if not provided.
//
// Returns:
//   - The same slice with modified elements.
func CopyWithin[T any](in []T, target int, start int, end ...int) []T {
	n := len(in)
	if n == 0 {
		return in
	}

	// Normalize negative indices
	if target < 0 {
		target += n
	}
	if start < 0 {
		start += n
	}

	// Determine end index
	stop := n
	if len(end) > 0 {
		stop = end[0]
		if stop < 0 {
			stop += n
		}
	}

	// Clamp indices
	if target >= n || start >= n || stop <= start || target == start {
		return in
	}

	// Compute actual copy length without overflow
	copyLen := min(stop-start, n-target)

	// Use copy to handle overlapping memory safely
	copy(in[target:target+copyLen], in[start:start+copyLen])

	return in
}
