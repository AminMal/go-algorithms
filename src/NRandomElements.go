package main

import(
	"fmt"
	"math/rand"
)

func createRandomSlice(size int, max int) []int {
	if size <= 0 {
		return []int{}
	} else {
		return append(createRandomSlice(size - 1, max), rand.Intn(max))
	}
}

func main() {
	fmt.Println(createRandomSlice(6, 50))
}
