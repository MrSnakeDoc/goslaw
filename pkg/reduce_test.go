package pkg

import (
	"testing"
)

func TestReduce_Sum(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := 15

	result := Reduce(input, 0, func(acc, val int) int {
		return acc + val
	})

	if result != expected {
		t.Errorf("Reduce() = %v, want %v", result, expected)
	}
}

func TestReduce_ConcatStrings(t *testing.T) {
	input := []string{"Go", "Lang", "Rocks"}
	expected := "GoLangRocks"

	result := Reduce(input, "", func(acc, val string) string {
		return acc + val
	})

	if result != expected {
		t.Errorf("Reduce() = %v, want %v", result, expected)
	}
}

func TestReduce_EmptySlice(t *testing.T) {
	input := []int{}
	expected := 42

	result := Reduce(input, 42, func(acc, val int) int {
		return acc + val
	})

	if result != expected {
		t.Errorf("Reduce() = %v, want %v", result, expected)
	}
}
