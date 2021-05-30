package tdd

import (
	"errors"
)

func elementExists(l []int, element int) (bool, error) {
	if len(l) == 0 {
		return false, errors.New("Invalid list")
	}
	for _, v := range l {
		if (element == v) {
			return true, nil
		}
	}
	return false, nil
}