package pkg

import (
	"testing"
)

func TestSome_TrueMatch(t *testing.T) {
	input := []int{1, 3, 5, 8}
	result := Some(input, func(x int) bool {
		return x%2 == 0
	})

	if !result {
		t.Errorf("Some() = false, want true (even number exists)")
	}
}

func TestSome_NoMatch(t *testing.T) {
	input := []int{1, 3, 5}
	result := Some(input, func(x int) bool {
		return x%2 == 0
	})

	if result {
		t.Errorf("Some() = true, want false (no even number)")
	}
}

func TestSome_EmptySlice(t *testing.T) {
	input := []int{}
	result := Some(input, func(x int) bool {
		return true
	})

	if result {
		t.Errorf("Some() = true on empty slice, want false")
	}
}
