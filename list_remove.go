package tdd

import (
	"errors"
)

func removeElement(l []int, index int) ([]int, error) {
	if index < 0 || index >= len(l) {
		return nil, errors.New("Invalid index")
	}
	return append(l[:index], l[index+1:]...), nil
}