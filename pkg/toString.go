package pkg

import (
	"reflect"
	"strconv"
	"strings"
)

// ToString returns a string representing the specified slice and its elements.
// Similar to JavaScript's Array.toString().
//
// Parameters:
//   - in: The input slice to convert to a string
//
// Returns:
//   - A string representation of the slice
func ToString[T any](in []T) string {
	if len(in) == 0 {
		return "[]"
	}

	var b strings.Builder
	b.WriteByte('[')

	for i, v := range in {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(elemToString(v))
	}

	b.WriteByte(']')
	return b.String()
}

// elemToString is a helper function that converts a single element to its string representation.
// It handles various basic types and provides appropriate formatting.
//
// Parameters:
//   - v: The value to convert to a string
//
// Returns:
//   - A string representation of the value
func elemToString(v any) string {
	switch v := v.(type) {
	case string:
		return `"` + v + `"`
	case int:
		return strconv.FormatInt(int64(v), 10)
	case int8, int16, int32, int64:
		return strconv.FormatInt(reflect.ValueOf(v).Int(), 10)
	case uint, uint8, uint16, uint32, uint64:
		return strconv.FormatUint(reflect.ValueOf(v).Uint(), 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(v)
	default:
		return "unknown"
	}
}
