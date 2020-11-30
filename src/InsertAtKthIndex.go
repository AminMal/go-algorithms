package main

import (
	"errors"
	"fmt"
)

func addElemAt(index int, element int, slice []int) ([]int, error) {
	if index > len(slice) || index < 0 {
		return []int{}, errors.New("index out of bounds")
	} else {
		if index == len(slice) {
			return append(slice, element), nil
		}
		leftSide := slice[:index]
		rightSide := slice[index + 1:]
		return append(append(leftSide, element), rightSide...), nil
	}
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	sl, err := addElemAt(7, 8, slice)
	if err != nil {
		fmt.Println("error occurred")
	} else {
		fmt.Println(sl)
	}

	sli, er := addElemAt(4, 11, slice)
	if er != nil {
		fmt.Println("error occurred")
	} else {
		fmt.Println(sli)
	}
}