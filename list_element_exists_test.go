package tdd

import (
	"testing"
)

func TestFindExistingElement(t *testing.T) {
	var l []int
	l = append(l, 1, 2, 3)
	element := 2
	found, err := elementExists(l, element)
	if err != nil {
		t.Errorf("There should be no existing error, got: %v", err)
	}
	if found == false {
		t.Errorf("Element %v was not found on list %v", element, l)
	}
}

func TestFindNotExistingElement(t *testing.T) {
	var l []int
	l = append(l, 1, 2, 3)
	element := 4
	found, err := elementExists(l, element)
	if err != nil {
		t.Errorf("There should be no existing error, got: %v", err)
	}
	if found == true {
		t.Errorf("Element %v was found on list %v", element, l)
	}
}

func TestFindInEmptyList(t *testing.T) {
	var l []int
	l = nil
	element := 1
	_, err := elementExists(l, element)
	if err == nil {
		t.Errorf("There should be an error, got: %v", err)
	}
}