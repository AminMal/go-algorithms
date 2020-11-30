package structs

import (
    "errors"
    "fmt"
)

type MyStack struct {
    elements []int
}

func (s MyStack) push(element int) {
    s.elements = append(s.elements, element)
}

func (s MyStack) pop() (int, MyStack, error) {
    if len(s.elements) == 0 {
	return -1, MyStack{elements: nil}, errors.New("stack is empty")
    } else {
        value := s.elements[len(s.elements) - 1]
        return value, MyStack{elements: s.elements[:len(s.elements) - 1]}, nil
    }
}

func main() {

    source := make([]int, 4)
    for i := 0; i < 4; i++ { // 0, 2, 4, 6
	source[i] = i * 2
    }

    var stack = MyStack{elements: source}
    lastElem, stackAfterThat, _ := stack.pop()


    fmt.Println("Stack: ", stack)
    fmt.Println("Last element: ", lastElem)
    fmt.Println("Stack after popping one element: ", stackAfterThat)
}
