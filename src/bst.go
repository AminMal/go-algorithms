package main

import(
    "fmt"
    )

func search(elems []int, target int, iterationCount int) (bool, int) {// int is going to be index

	fmt.Println(iterationCount, "th iterate")
	if len(elems) == 1 {
		if target != elems[0] {return false, 0}
		return true, 0
	}
    var currentIndex int = len(elems) / 2 // division of ints => int
    if elems[currentIndex] == target {
	return true, currentIndex
    } else {
	if elems[currentIndex] > target {
	    return search(elems[:currentIndex], target, iterationCount + 1)
	} else {
	    return search(elems[currentIndex:], target, iterationCount + 1)
	}
    }
    return false, -1
}

func main() {

    var source []int
    source = append(source, 1)
    source = append(source, 2)
    source = append(source, 3)
    source = append(source, 4)
    source = append(source, 5)
    source = append(source, 6)
    source = append(source, 7)
    source = append(source, 8)
    source = append(source, 9)
    source = append(source, 19)

    target := 19
    stat, index := search(source, target, 0)
    if stat == false {
        fmt.Println("element ", target, "not found")
    } else {
	fmt.Println("element ", target, "is at index: ", index)
    }
}
