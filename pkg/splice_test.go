package pkg

import (
	"reflect"
	"testing"
)

func TestSplice_RemoveAndInsert(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := []int{1, 9, 8, 4, 5}

	result := Splice(input, 1, 2, 9, 8)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Splice() = %v, want %v", result, expected)
	}
}

func TestSplice_StartBeyondLength(t *testing.T) {
	input := []int{1, 2, 3}
	expected := []int{1, 2, 3, 99}

	result := Splice(input, 10, 2, 99)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Splice() = %v, want %v", result, expected)
	}
}

func TestSplice_DeleteOnly(t *testing.T) {
	input := []string{"a", "b", "c", "d"}
	expected := []string{"a", "d"}

	result := Splice(input, 1, 2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Splice() = %v, want %v", result, expected)
	}
}
