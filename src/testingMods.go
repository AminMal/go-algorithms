package main

import(
	"fmt"
	str "me/structs"
)

func main() {

	stack := str.MyStack{Elements: []int{1, 2, 3, 4}}
	fmt.Println("stack is:", stack)
	list := str.List{Elements: []int{1, 2, 3}}
	val, _ := list.Get(2)
	fmt.Println("getting at index 2", val)

    newStack := stack.Push(7)
	fmt.Println("Stack is ", newStack)
}

