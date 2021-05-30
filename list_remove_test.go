package tdd

import (
	"testing"
)

func TestRemoveExistingElement(t *testing.T) {
	var l []int
	l = append(l, 1, 2, 3)
	index := 1
	newL, _ := removeElement(l, index)
	if len(newL) != 2 {
		t.Errorf("The length should be 2: %v", l)
	}
	if l[0] != 1 || l[1] != 3 {
		t.Errorf("Incorrect element was removed")
	}
}

func TestRemoveNotExistingElement(t *testing.T) {
	var l []int
	l = append(l, 1, 2, 3)
	index := 5
	_, err := removeElement(l, index)
	if err == nil {
		t.Errorf("No error was raised when using invalid index.")
	}
}

func TestRemoveOnEmptyList(t *testing.T) {
	var l []int
	l = nil
	index := 0
	_, err := removeElement(l, index)
	if err == nil {
		t.Errorf("No error was raised when using invalid index.")
	}
}