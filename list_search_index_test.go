package tdd

import (
	"testing"
)

func TestSearchIndex(t *testing.T) {
	var l []int
	l = append(l, 1, 2, 3)
	element := 2
	index, err := searchElement(l, element)
	if err != nil {
		t.Errorf("There should be no error")
	}
	if index != 1 {
		t.Errorf("The found index should be 1, got: %v", index)
	}
}

func TestIndexInvalidElement(t *testing.T) {
	var l []int
	l = append(l, 1, 2, 3)
	element := 5
	index, err := searchElement(l, element)
	if err != nil {
		t.Errorf("There should be no error")
	}
	if index != -1 {
		t.Errorf("Index should be -1, got: %v", index)
	}
}

func TestIndexEmptyList(t *testing.T) {
	var l []int
	l = nil
	element := 0
	_, err := searchElement(l, element)
	if err == nil {
		t.Errorf("Error should exist.")
	}
}
