package tdd

import (
	"testing"
)

func TestAddEmptyList(t *testing.T) {
	var l []int
	l = append(l, 1, 2, 3)
	if len(l) != 3 {
		t.Errorf("The list should have 3 elements: %v", l)
	}
}

func TestAddListWithElements(t *testing.T) {
	var l [] int
	l = append(l, 1, 2, 3)
	l = append(l, 4, 5)
	if len(l) != 5 {
		t.Errorf("The list should have 5 elements: %v", l)
	}
}

func TestAddToListWithCapacity(t *testing.T) {
	// Make a list with initial capacity of 3 but length 0
	l := make([]int, 0, 3)
	if len(l) != 0 {
		t.Errorf("The list should be empty at this point: %v", l)
	}
	l = append(l, 1, 2)
	if len(l) != 2 && cap(l) != 3 {
		t.Errorf("The list should have 2 elements: %v", l)
	}
}
