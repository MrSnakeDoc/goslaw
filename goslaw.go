package goslaw

import (
	"fmt"

	inner "github.com/MrSnakeDoc/goslaw/pkg"
)

func Concat[T any](slices ...[]T) []T {
	return inner.Concat(slices...)
}

func CopyWithin[T any](in []T, target int, start int, end ...int) []T {
	return inner.CopyWithin(in, target, start, end...)
}

func Entries[K comparable, V any](m map[K]V) []struct {
	Key   K
	Value V
} {
	return inner.Entries(m)
}

func Every[T any](in []T, pred func(T) bool) bool {
	return inner.Every(in, pred)
}

func EveryIndex[T any](in []T, pred func(T) bool) (bool, int) {
	return inner.EveryIndex(in, pred)
}

func Fill[T any](template []T, value T) []T {
	return inner.Fill(template, value)
}

func Filter[T any](in []T, keep func(T) bool) []T {
	return inner.Filter(in, keep)
}

func Find[T any](in []T, pred func(T) bool) (T, bool) {
	return inner.Find(in, pred)
}

func FindIndex[T any](in []T, pred func(T) bool) int {
	return inner.FindIndex(in, pred)
}

func Flat[T any](chunks [][]T) []T {
	return inner.Flat(chunks)
}

func FlatMap[T any, R any](in []T, fn func(T) []R) []R {
	return inner.FlatMap(in, fn)
}

func ForEach[T any](in []T, fn func(T)) {
	inner.ForEach(in, fn)
}

func Includes[T comparable](in []T, target T) bool {
	return inner.Includes(in, target)
}

func IndexOf[T comparable](in []T, target T) int {
	return inner.IndexOf(in, target)
}

func IndexOfFunc[T any](in []T, pred func(T) bool) int {
	return inner.IndexOfFunc(in, pred)
}

func Join[T fmt.Stringer](in []T, sep string) string {
	return inner.Join(in, sep)
}

func Keys[K comparable, V any](m map[K]V) []K {
	return inner.Keys(m)
}

func LastIndexOf[T comparable](in []T, target T) int {
	return inner.LastIndexOf(in, target)
}

func LastIndexOfFunc[T any](in []T, pred func(T) bool) int {
	return inner.LastIndexOfFunc(in, pred)
}

func Map[T any, R any](in []T, f func(T) R) []R {
	return inner.Map(in, f)
}

func Pop[T any](in []T) (T, []T) {
	return inner.Pop(in)
}

func PopFree[T any](in []T) (T, []T) {
	return inner.PopFree(in)
}

func Reduce[T, R any](in []T, init R, fn func(R, T) R) R {
	return inner.Reduce(in, init, fn)
}

func ReduceRight[T, R any](in []T, init R, fn func(R, T) R) R {
	return inner.ReduceRight(in, init, fn)
}

func Reverse[T any](in []T) []T {
	return inner.Reverse(in)
}

func Shift[T any](in []T) (T, []T) {
	return inner.Shift(in)
}

func Slice[T any](in []T, start, end int) []T {
	return inner.Slice(in, start, end)
}

func Some[T any](in []T, pred func(T) bool) bool {
	return inner.Some(in, pred)
}

func Sort[T any](in []T, less func(a, b T) bool) {
	inner.Sort(in, less)
}

func Splice[T any](in []T, start, deleteCount int, items ...T) []T {
	return inner.Splice(in, start, deleteCount, items...)
}

func ToReversed[T any](in []T) []T {
	return inner.ToReversed(in)
}

func ToString[T any](in []T) string {
	return inner.ToString(in)
}

func Unshift[T any](in []T, item T) []T {
	return inner.Unshift(in, item)
}

func Values[K comparable, V any](m map[K]V) []V {
	return inner.Values(m)
}

func With[T any](in []T, idx int, val T) []T {
	return inner.With(in, idx, val)
}
