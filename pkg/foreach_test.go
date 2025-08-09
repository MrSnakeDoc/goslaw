package pkg

import (
	"testing"
)

func TestForEach_OrderPreserved(t *testing.T) {
	input := []int{1, 2, 3}
	var output []int

	ForEach(input, func(x int) {
		output = append(output, x)
	})

	expected := []int{1, 2, 3}
	for i := range expected {
		if output[i] != expected[i] {
			t.Errorf("ForEach order mismatch at index %d: got %d, want %d", i, output[i], expected[i])
		}
	}
}

func TestForEach_SideEffects(t *testing.T) {
	input := []string{"a", "b", "c"}
	count := 0

	ForEach(input, func(s string) {
		if s != "" {
			count++
		}
	})

	if count != len(input) {
		t.Errorf("ForEach did not apply function to each element: got %d, want %d", count, len(input))
	}
}

func TestForEach_EmptySlice(t *testing.T) {
	var called bool

	ForEach([]int{}, func(x int) {
		called = true
	})

	if called {
		t.Errorf("ForEach should not call function on empty slice")
	}
}
