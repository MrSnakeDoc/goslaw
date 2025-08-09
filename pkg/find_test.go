package pkg

import (
	"testing"
)

func TestFind_Found(t *testing.T) {
	input := []int{3, 7, 10, 2}
	expected := 10

	result, ok := Find(input, func(x int) bool {
		return x > 8
	})
	if !ok {
		t.Fatalf("Expected to find a value but got %v", result)
	}

	if result != expected {
		t.Errorf("Find() = %v, want %v", result, expected)
	}
}

func TestFind_NotFound(t *testing.T) {
	input := []string{"apple", "banana", "cherry"}
	expected := ""

	result, ok := Find(input, func(s string) bool {
		return s == "kiwi"
	})
	if ok {
		t.Fatalf("Expected not to find a value but got %q", result)
	}

	if result != expected {
		t.Errorf("Find() = %q, want %q", result, expected)
	}
}

func TestFind_EmptySlice(t *testing.T) {
	input := []int{}
	expected := 0

	result, ok := Find(input, func(x int) bool {
		return x == 1
	})
	if ok {
		t.Fatalf("Expected not to find a value in an empty slice but got %v", result)
	}

	if result != expected {
		t.Errorf("Find() = %v, want %v", result, expected)
	}
}
