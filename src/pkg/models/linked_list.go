package models

import (
	"fmt"
	"strconv"
)

type Queue interface {
	Append(value int)
}

type Stack interface {
	Pop() Node
	Push(value int)
}
type Node struct {
	Value int
	Next  *Node
	// Previous *Node
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
		currentNode := *list.Head

		for currentNode.Next != nil {
			currentNode = *currentNode.Next
		}

		fmt.Println("Returning tail")
		return &currentNode
	}
}

func (list *LinkedList) Push(value int) {
	node := &Node{Value: value, Next: nil}

	tail := list.GetTail() // pointer to the node object without next node item

	if tail == nil {
		// adding node as a head
		list.Head = node
		list.Length++

	} else {
		// Tagging node with its previous node

		//adding node as a new tail
		tail.Next = node // change nil pointer in tail.Next to new node item
		list.Length++
	}
}

func (list *LinkedList) Print() {
	values := ""

	if list.Length == 1 {
		values = strconv.Itoa(list.GetTail().Value)
	}

	if list.Length > 1 {
		currentNode := *list.Head
		values += strconv.Itoa(currentNode.Value)

		for currentNode.Next != nil {
			currentNode = *currentNode.Next
			values += ", " + strconv.Itoa(currentNode.Value)
		}
	}

	fmt.Printf("[%s]\n", values)
}
