package pkg

// Reduce executes a reducer function on each element of the slice, resulting in a single output value.
// Similar to JavaScript's Array.reduce().
//
// Parameters:
//   - in: The input slice to reduce
//   - init: The initial value for the accumulator
//   - fn: A function that takes an accumulator and the current element and returns a new accumulator
//
// Returns:
//   - The final value of the accumulator
func Reduce[T, R any](in []T, init R, fn func(R, T) R) R {
	acc := init
	for _, v := range in {
		acc = fn(acc, v)
	}
	return acc
}

// ReduceRight works like Reduce, but processes the slice from right to left.
// Similar to JavaScript's Array.reduceRight().
//
// Parameters:
//   - in: The input slice to reduce
//   - init: The initial value for the accumulator
//   - fn: A function that takes an accumulator and the current element and returns a new accumulator
//
// Returns:
//   - The final value of the accumulator
func ReduceRight[T, R any](in []T, init R, fn func(R, T) R) R {
	acc := init
	for i := len(in) - 1; i >= 0; i-- {
		acc = fn(acc, in[i])
	}
	return acc
}
