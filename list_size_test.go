package tdd

import (
	"testing"
)

func TestListSize(t *testing.T) {
	var l []int
	l = append(l, 1, 2, 3)
	if len(l) != 3 {
		t.Errorf("Error in the length of the list: %v", l)
	}
}

func TestEmptyListSize(t *testing.T) {
	var l []int
	if len(l) != 0 {
		t.Errorf("The list should be empty: %v", l)
	}
}

func TestSizeRemovingElements(t *testing.T) {
	var l []int
	l = append(l, 1, 2, 3)
	// Remove the first element
	l = append(l[1:])
	if len(l) != 2 {
		t.Errorf("Error in the length of the list while removing element: %v", l)
	}
}