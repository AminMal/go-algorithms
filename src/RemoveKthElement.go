package main

import(
	"fmt"
	"errors"
)

func removeKthElem(index int, slice []int) ([]int, error) {
	if index >= len(slice) || index < 0  {
		return []int{}, errors.New("index out of bound")
	} else {
		leftSide := slice[:index]
		rightSide := slice[index + 1:]
		return append(leftSide, rightSide...), nil
	}
}

func main() {

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}

	index := 5
	element, err := removeKthElem(index, slice)
	if err != nil {
		fmt.Println("An error occurred:", err)
	} else {
		fmt.Println(element)
	}

}