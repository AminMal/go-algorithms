package main

import(
	"fmt"
)

func checkEq(slice []int, firstIndex int, secondIndex int) bool {
	if firstIndex == -1 {
		return true
	} else {
		if slice[firstIndex] != slice[secondIndex] {
			return false
		} else {
			return checkEq(slice, firstIndex - 1, secondIndex + 1)
		}
	}
}

func isPalindrome(slice []int) bool {
	if len(slice) % 2 == 0 {
		return false
	} else {
		middleIndex := len(slice) / 2
		aIndex := middleIndex - 1 // I know namings suck
		bIndex := middleIndex + 1 // means the element after middle one and so on

		return checkEq(slice, aIndex, bIndex)
	}
}


func main() {

	aPalindrome := []int{1, 2, 3, 2, 1}
	aNonPalindrome := []int{1, 2, 3, 3, 2, 1}

	fmt.Println("a palindrome is palindrome: ", isPalindrome(aPalindrome))
	fmt.Println("a non palindrome is palindrome: ", isPalindrome(aNonPalindrome))

}