package main

import(
    "fmt"
    )

func search(elems []int, target int, iterationCount int) bool {

	if len(elems) == 1 {
		if target != elems[0] {return false}
		return true
	}
    var currentIndex int = len(elems) / 2 // division of ints => int
    if elems[currentIndex] == target {
	return true
    } else {
	if elems[currentIndex] > target {
	    return search(elems[:currentIndex], target, iterationCount + 1)
	} else {
	    return search(elems[currentIndex:], target, iterationCount + 1)
	}
    }
}

func main() {

    input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    target := 10

    stat := search(input, target, 0)
    if stat == false {
        fmt.Println("element ", target, "not found")
    } else {
	fmt.Println("element", target, "exists")
    }

}
