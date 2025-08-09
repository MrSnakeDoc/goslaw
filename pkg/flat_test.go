package pkg

import (
	"reflect"
	"testing"
)

func TestFlat(t *testing.T) {
	input := [][]string{
		{"a", "b"},
		{},
		{"c"},
	}
	expected := []string{"a", "b", "c"}

	result := Flat(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Flat() = %v, want %v", result, expected)
	}
}
