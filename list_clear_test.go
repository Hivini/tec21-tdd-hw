package tdd

import (
	"testing"
)

func TestClearList(t *testing.T) {
	var l []int
	l = append(l, 1, 2, 3)
	// Note for professor:
	// When I "clear the list" I do not remove the allocated memory.
	l = l[:0]
	if len(l) != 0 {
		t.Errorf("The list should be empty: %v", l)
	}
}

func TestClearEmptyList(t *testing.T) {
	var l [] int
	l = l[:0]
	if len(l) != 0 {
		t.Errorf("The list should be empty: %v", l)
	}
}

func TestClearNilList(t *testing.T) {
	var l [] int
	// Setting it up to nil will make the garbage collector get rid
	// of any elements. Setting the capacity to 0.
	l = nil
	if len(l) != 0 {
		t.Errorf("The list should be empty: %v", l)
	}
}

