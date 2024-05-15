package main

import "fmt"

type Node[T comparable] struct {
	v    T
	next *Node[T]
}

func (n *Node[T]) Add(value T) {
	node := &Node[T]{}
	node.v = value
	node.next = nil

	for n.next != nil {
		n = n.next
	}
	n.next = node
}

func (n *Node[T]) Insert(value T, pos int) {
	node := &Node[T]{}
	node.v = value

	for i := 0; i < pos; i++ {
		n = n.next
	}

	node.next = n.next
	n.next = node
}

func (n *Node[T]) Index(value T) int {
	i := 0
	n = n.next
	for n != nil {
		if n.v == value {
			return i
		}
		i++
		n = n.next
	}

	return -1
}

func (n *Node[T]) PrintValues() {
	n = n.next
	for n != nil {
		fmt.Printf("%v\t", n.v)
		n = n.next
	}
	fmt.Printf("\n")
}
