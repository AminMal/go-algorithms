package main

import (
	"fmt"
)

func lastElem (slice []int) int {
	if len(slice) == 0 {
		return 0
	} else {
		return slice[len(slice) - 1]
	}
}

func reverseSlice (slice []int) []int {

	var result []int

	for k := range slice {
		sliceToPass := slice[:len(slice) - k]
		result = append(result, lastElem(sliceToPass))
	}
	return result
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(reverseSlice(slice))
}