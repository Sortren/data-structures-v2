package models

import (
	"fmt"
)

type Queue interface {
	Push(value int)
}

type Stack interface {
	Pop() int
	Append(value int)
}
type Node struct {
	Value    int
	Next     *Node
	Previous *Node
}

func (node *Node) printExplicitly() {
	// print the node in the explicit way with each its properties (Next and Previous)

	fmt.Println("[", node.Value, " next->", node.Next, " previous->", node.Previous, "]")
}

func (node *Node) print(listLength int, delimeter string, isFirst bool) {
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
type LinkedList struct {
	Head   *Node
	Length int
}

func (list *LinkedList) GetTail() *Node {
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
func (list *LinkedList) Append(value int) {
	node := &Node{Value: value, Next: nil}

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
func (list *LinkedList) Push(value int) {
	node := &Node{Value: value, Next: list.Head, Previous: nil}
	list.Head.Previous = node
	list.Head = node
}

func (list *LinkedList) Pop() int {
	tail := list.GetTail()
	prevOfTail := list.GetTail().Previous
	prevOfTail.Next = nil
	tail.Previous = nil

	return tail.Value
}

func (list *LinkedList) Print() {
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

func (list *LinkedList) PrintExplicitly() {
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
