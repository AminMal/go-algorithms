package main

import (
	"errors"
	"fmt"
)

type List struct {
	elements []int
}

func (l List) Add(element int) List {
	lPointer := &l
	lPointer.elements = append(l.elements, element)
	return List{elements: lPointer.elements}
}
func iterate(index int, slice []int) int { // not going to use the [index] on slice
	if index == 0 {
		return slice[0]
	} else {
		return iterate(index - 1, slice[1:])
	}
}

func (l List) Get(index int) (int, error) { // works fine
	if index < 0 || index >= len(l.elements) {
		return -1, errors.New("index out of bound for elements")
	} else {
		return iterate(index, l.elements), nil
	}
}

// I also wanted to implement map, foreach, ... methods, but it's just not possible
func (l List) IsEmpty() bool {
	if len(l.elements) == 0 {
		return true
	} else {
		return false
	}
}

func lengthIterator(elements []int, preLength int) int { // should be 0 by default
	listOfElems := List{elements: elements}
	if listOfElems.IsEmpty() {
		return preLength
	} else {
		return lengthIterator(elements[1:], preLength + 1)
	}
}

func (l List) Length() int {
	return lengthIterator(l.elements, 0)
}

func main() {
	list := List{elements: []int{1, 2, 3, 4, 5}}

	fmt.Println(list)
	fmt.Println("Length is: ", list.Length())

	appendedList := list.Add(6)
	fmt.Println("Appended list: ", appendedList)
	fmt.Println("Length of appended list is: ", appendedList.Length())

	emptyList := List{}
	fmt.Println("empty list is empty: ", emptyList.IsEmpty())

	fmt.Println("appended list is empty: ", appendedList.IsEmpty())

	fmt.Println("trying to get element at index 5 on appended list")
	element, err := appendedList.Get(5)
	if err == nil {
		fmt.Println("Error is nil, element is: ", element)
	}
	fmt.Println("Trying to get element at index 5 on first list (causes error)")
	element, err = list.Get(5)
	if err != nil {
		fmt.Println("We got an error")
	}
}