package main

import (
	"errors"
	"fmt"
)

// only for integers
type Node struct {
	data int
	next *Node
	previous *Node
}

type LinkedList struct {
	nodes []Node
}

func (node Node) NextVal() (int, error) {
	if node.next == nil {
		return -1, errors.New("last value")
	} else {
		return node.next.data, nil
	}
}

func (node Node) PreviousVal() (int, error) {
	if node.previous  == nil {
		return -1, errors.New("first element")
	} else {
		return node.previous.data, nil
	}
}

func (ll LinkedList) AttachLeft(node Node) LinkedList {
	if len(ll.nodes) == 0 {
		return LinkedList{[]Node{node}}
	} else {
		first := ll.nodes[0]
		firstUpdated := Node{first.data, &node, first.next}
		firstAsSlice := []Node{firstUpdated}
		return LinkedList{nodes: append(firstAsSlice, ll.nodes...)}
	}
}

func (ll LinkedList) Attach(node Node) LinkedList {
	if len(ll.nodes) == 0 {
		return LinkedList{nodes: []Node{node}}
	} else {
		last := ll.nodes[len(ll.nodes) - 1]
		lastUpdated := Node{last.data, &node, last.previous}
		return LinkedList{nodes: append(ll.nodes, lastUpdated)}
	}
}

func (ll LinkedList) Get(index int) (Node, error) {
	if index < 0 || index >= len(ll.nodes) {
		return Node{}, errors.New("index out of bound")
	} else {
		return ll.nodes[index], nil
	}
}

func main() {

	firstNode := Node{data: 0}
	secondNode := Node{data: 1}
	thirdNode := Node{data: 2}
	fourthNode := Node{data: 3}
	fifthNode := Node{data: 4}

	emptyList := LinkedList{}
	list1 := emptyList.Attach(secondNode).Attach(thirdNode).Attach(firstNode)

	fmt.Println("first list", list1)

	third, err := list1.Get(0)
	if err == nil {
		fmt.Println("third node:", third)
	}

	list2 := list1.Attach(fourthNode).Attach(fifthNode)
	fmt.Println(list2)

    fmt.Println("end")
	fmt.Println("end2")
}