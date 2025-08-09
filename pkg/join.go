package pkg

import (
	"fmt"
	"strings"
)

// Join appends item to the end of in and returns the resulting slice.
// It behaves like JavaScript’s Array.push()/concat(); it does NOT produce
// a string the way Array.join() does.
func Join[T fmt.Stringer](in []T, sep string) string {
	if len(in) == 0 {
		return ""
	}
	var b strings.Builder
	b.WriteString(in[0].String())
	for _, v := range in[1:] {
		b.WriteString(sep)
		b.WriteString(v.String())
	}
	return b.String()
}
