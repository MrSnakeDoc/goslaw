package pkg

import (
	"reflect"
	"testing"
)

func TestFlatMap(t *testing.T) {
	input := []int{1, 2, 3}
	expected := []int{1, 1, 2, 2, 3, 3}

	result := FlatMap(input, func(x int) []int {
		return []int{x, x}
	})

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("FlatMap() = %v, want %v", result, expected)
	}
}

func TestFlatMap_EmptyResult(t *testing.T) {
	input := []int{1, 2, 3}
	expected := []int{}

	result := FlatMap(input, func(x int) []int {
		return nil
	})

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("FlatMap() = %v, want %v", result, expected)
	}
}
