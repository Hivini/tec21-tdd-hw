package tdd

import (
	"testing"
)

func TestGetElement(t *testing.T) {
	var l []int
	l = append(l, 1, 2, 3)
	element := l[0]
	if element != 1 {
		t.Errorf("Element is not correct: %v element", element)
	}
}

func TestGetNotExistingElement(t *testing.T) {
	var l []int
	l = append(l, 1, 2)
	index := 5
	// There is no way to check if the index exists, so do it manually.
	if index < len(l) {
		t.Errorf("The index should not exist on list: %v", l)
	}
}

func TestGetOnEmptyList(t *testing.T) {
	var l []int
	l = nil
	index := 0
	// There is no way to check if the index exists, so do it manually.
	if index < len(l) {
		t.Errorf("The index should not exist on list: %v", l)
	}
}
