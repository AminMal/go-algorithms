package main

import(
	"fmt"
	"errors"
)

func except(index int, slice []int) ([]int, error) {
	if index >= len(slice) || index < 0 {
		return []int{}, errors.New("index out of bounds")
	} else {
		leftSide := slice[:len(slice) - index - 1]
		rightSide := slice[index + 1:]
		return append(leftSide, rightSide...), nil
	}
}


func main() {

	originalSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(except(-2, originalSlice))
	fmt.Println(except(2, originalSlice))
	fmt.Println(except(7, originalSlice))
	/*
			[] index out of bounds
			[1 2 4 5] <nil>
			[] index out of bounds
	 */

}