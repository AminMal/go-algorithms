package main

import "fmt"

func main() {

    // input: [1, 2, 3, 4] (or whatever) => output: [24, 12, 8, 6]
	input := []int{1, 2, 3, 4, 5}
	var output [5]int

	for i := range input {
		rightSide := input[i + 1:]
		leftSide := input[:i]

		fmt.Println("Right side: ", rightSide)
		fmt.Println("left side: ", leftSide)

		rightProd := 1
		leftProd := 1
		for j := range rightSide {
			rightProd *= rightSide[j]
		}
		for k := range leftSide {
			leftProd *= leftSide[k]
		}
		output[i] = leftProd * rightProd
		rightProd = 1
		leftProd = 1
	}

    fmt.Println(output)
}