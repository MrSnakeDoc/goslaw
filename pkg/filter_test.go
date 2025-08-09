package pkg

import (
	"reflect"
	"testing"
)

func TestFilter_EvenNumbers(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	expected := []int{2, 4, 6}

	result := Filter(input, func(x int) bool {
		return x%2 == 0
	})

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Filter() = %v, want %v", result, expected)
	}
}

func TestFilter_EmptyResult(t *testing.T) {
	input := []int{1, 3, 5}
	expected := []int{}

	result := Filter(input, func(x int) bool {
		return x%2 == 0
	})

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Filter() = %v, want %v", result, expected)
	}
}

func TestFilter_EmptySlice(t *testing.T) {
	input := []int{}
	expected := []int{}

	result := Filter(input, func(x int) bool {
		return true
	})

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Filter() = %v, want %v", result, expected)
	}
}
