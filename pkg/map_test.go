package pkg

import (
	"reflect"
	"testing"
)

func TestMap_IntToInt(t *testing.T) {
	input := []int{1, 2, 3}
	expected := []int{2, 4, 6}

	result := Map(input, func(x int) int {
		return x * 2
	})

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Map() = %v, want %v", result, expected)
	}
}

func TestMap_StringToLength(t *testing.T) {
	input := []string{"hi", "world", ""}
	expected := []int{2, 5, 0}

	result := Map(input, func(s string) int {
		return len(s)
	})

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Map() = %v, want %v", result, expected)
	}
}

func TestMap_EmptySlice(t *testing.T) {
	input := []int{}
	expected := []int{}

	result := Map(input, func(x int) int {
		return x * 2
	})

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Map() = %v, want %v", result, expected)
	}
}
