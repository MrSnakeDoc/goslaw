package pkg

import (
	"reflect"
	"testing"
)

func TestSort_Ints(t *testing.T) {
	input := []int{5, 2, 9, 1}
	expected := []int{1, 2, 5, 9}

	Sort(input, func(a, b int) bool { return a < b })

	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Sort() = %v, want %v", input, expected)
	}
}

func TestSort_Strings(t *testing.T) {
	input := []string{"banana", "apple", "cherry"}
	expected := []string{"apple", "banana", "cherry"}

	Sort(input, func(a, b string) bool { return a < b })

	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Sort() = %v, want %v", input, expected)
	}
}

func TestSort_EmptySlice(t *testing.T) {
	input := []int{}
	expected := []int{}

	Sort(input, func(a, b int) bool { return a < b })

	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Sort() = %v, want %v", input, expected)
	}
}
