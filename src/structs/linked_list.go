package structs

import (
	"errors"
)

type Node struct {
	data int
	previous *Node
	next *Node
}

type LinkedList struct {
	nodes []Node
}

func (ll LinkedList) Append(element int) LinkedList {
	if len(ll.nodes) == 0 {
		return LinkedList{nodes: []Node{{data: element}}}
	} else {
		last := ll.nodes[len(ll.nodes) - 1]
		node := Node{data: element, previous: &(ll.nodes[len(ll.nodes) - 1])}
		updatedLeft := Node{data: last.data, previous: last.previous, next: &node}
		return LinkedList{nodes: append(ll.nodes[:len(ll.nodes) - 1], updatedLeft, node)}
	}
}

func (ll LinkedList) Prepend(element int) LinkedList {
	if len(ll.nodes) == 0 {
		return LinkedList{nodes: []Node{{data: element}}}
	} else {
		right := ll.nodes[0]
		node := Node{data: element, next: &(ll.nodes[0])}
		updatedRight := Node {data: right.data, next: right.next, previous: &node}
		slice := []Node{node, updatedRight}
		return LinkedList{
			nodes: append(slice, ll.nodes[1:]...),
		}
	}
}

func (ll LinkedList) Get(index int) (int, error) {
	if index < 0 || index >= len(ll.nodes) {
		return -1, errors.New("index out of bound")
	} else {
		elem := ll.nodes[index]
		return elem.data, nil
	}
}