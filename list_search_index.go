package tdd

import (
	"errors"
)

func searchElement(l []int, element int) (int, error) {
	if len(l) == 0 {
		return -1, errors.New("Invalid list")
	}
	foundIndex := -1
	for i, v := range l {
		if (element == v) {
			foundIndex = i
			break
		}
	}
	return foundIndex, nil
}