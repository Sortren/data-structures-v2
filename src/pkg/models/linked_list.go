package models

import (
	"fmt"
)

type Queue[T any] interface {
	Push(value T)
}

type Stack[T any] interface {
	Pop() T
	Append(value T)
}
type Node[T any] struct {
	Value    T
	Next     *Node[T]
	Previous *Node[T]
}

func (node *Node[T]) printExplicitly() {
	// print the node in the explicit way with each its properties (Next and Previous)

	fmt.Println("[", node.Value, " next->", node.Next, " previous->", node.Previous, "]")
}

func (node *Node[T]) print(listLength int, delimeter string, isFirst bool) {
	// print the node in the regular list way [val1, val2, val3, ...]

	if listLength > 1 {
		if isFirst {
			fmt.Print(node.Value)
		} else {
			fmt.Print(delimeter, node.Value)
		}
	} else {
		fmt.Println("[", node.Value, "]")
	}
}

// LIFO and FIFO queueing
type LinkedList[T any] struct {
	Head   *Node[T]
	Length int
}

func (list *LinkedList[T]) GetTail() *Node[T] {
	if list.Head == nil || list.Head.Next == nil {
		// List is empty or has one item which is head
		return list.Head // returns nil if list is empty

	} else {
		currentNode := list.Head

		for currentNode.Next != nil {
			currentNode = currentNode.Next
		}

		// fmt.Println("Returning tail ", currentNode)
		return currentNode
	}
}

// LIFO
func (list *LinkedList[T]) Append(value T) {
	node := &Node[T]{Value: value, Next: nil}

	tail := list.GetTail() // pointer to the node object without next node item

	if tail == nil {
		// adding node as a head
		list.Head = node
		list.Length++

	} else {
		// Tagging node with its previous node
		node.Previous = tail

		//adding node as a new tail
		tail.Next = node // change nil pointer in tail.Next to new node item
		list.Length++
	}
}

// FIFO
func (list *LinkedList[T]) Push(value T) {
	node := &Node[T]{Value: value, Next: list.Head, Previous: nil}
	list.Head.Previous = node
	list.Head = node
}

func (list *LinkedList[T]) Pop() T {
	tail := list.GetTail()
	prevOfTail := list.GetTail().Previous
	prevOfTail.Next = nil
	tail.Previous = nil

	return tail.Value
}

func (list *LinkedList[T]) Print() {
	if list.Length == 0 {
		fmt.Println("[ ]")
	}

	if list.Length == 1 {
		list.GetTail().print(list.Length, "", true)
	}

	if list.Length > 1 {
		currentNode := *list.Head

		fmt.Print("[")
		currentNode.print(list.Length, "", true)

		for currentNode.Next != nil {
			currentNode = *currentNode.Next
			currentNode.print(list.Length, ", ", false)
		}
		fmt.Print("]")
	}
}

func (list *LinkedList[T]) PrintExplicitly() {
	if list.Length == 1 {
		list.GetTail().printExplicitly()
	}

	if list.Length > 1 {
		currentNode := *list.Head
		currentNode.printExplicitly()

		for currentNode.Next != nil {
			currentNode = *currentNode.Next
			currentNode.printExplicitly()
		}
	}
}
