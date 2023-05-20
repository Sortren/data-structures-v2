package main

import (
	"fmt"
)

func main() {
	integerList := LinkedList[int]{}

	for i := 1; i <= 5; i++ {
		integerList.Append(i)
	}

	for i := 10; i >= 5; i-- {
		integerList.Push(i)
	}

	fmt.Println("Initial integer list")
	integerList.Show()

	fmt.Print("\n\n")

	fmt.Println("Adding at index 2")
	integerList.AddAt(99, 2)
	integerList.Show()

	fmt.Print("\n\n")

	fmt.Println("Popping value")
	intVal := integerList.Pop()
	fmt.Println("Popped value ->", intVal)
	integerList.Show()

	fmt.Print("\n\n")

	stringList := LinkedList[string]{}

	for i := 1; i <= 5; i++ {
		stringList.Append(fmt.Sprintf("%s%d", "a", i))
	}

	for i := 10; i >= 5; i-- {
		stringList.Push(fmt.Sprintf("%s%d", "b", i))
	}

	fmt.Println("Initial string list")
	stringList.Show()

	fmt.Print("\n\n")

	fmt.Println("Popping value")
	strVal := stringList.Pop()
	fmt.Println("Popped value ->", strVal)
	stringList.Show()

	fmt.Print("\n\n")

	strVal2 := stringList.DeleteAt(2)
	fmt.Println("DeletedAt value ->", strVal2)
	stringList.Show()

	fmt.Print("\n\n")

	fmt.Println(stringList.Length)
	stringList.ShowAt(5)
}

type Stack[T any] interface {
	Push(value T)
	Pop() T
}

type Queue[T any] interface {
	Append(value T)
	Pop() T
}

type Node[T any] struct {
	Value    T
	Next     *Node[T]
	Previous *Node[T]
}

type LinkedList[T any] struct {
	Head   *Node[T]
	Length int
}

// Push represents LIFO queue
func (list *LinkedList[T]) Push(value T) {
	node := &Node[T]{Value: value, Next: list.Head, Previous: nil}
	list.Head.Previous = node
	list.Head = node
	list.Length++
}

// Append represents FIFO stack
func (list *LinkedList[T]) Append(value T) {
	node := &Node[T]{Value: value, Next: nil}

	tail := list.getTail() // pointer to the node object without next node item

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

func (list *LinkedList[T]) Pop() T {
	tail := list.getTail()
	prevOfTail := list.getTail().Previous
	prevOfTail.Next = nil
	tail.Previous = nil

	return tail.Value
}

func (list *LinkedList[T]) AddAt(value T, index int) {
	if index >= list.Length || index < 0 {
		fmt.Println("can't add value at position out of list's range")
		return
	}

	prev := &Node[T]{} // empty previous node to keep track of it
	currentNode := list.Head

	if index == 0 {
		list.Head = &Node[T]{
			Value:    value,
			Next:     currentNode,
			Previous: nil,
		}
		return
	}

	currentPosition := 0
	for {
		if currentPosition == index {
			newNode := &Node[T]{
				Value: value,
				Next:  currentNode.Next,
			}

			prev.Next = newNode
			currentNode = newNode
			list.Length++
			return
		} else {
			prev = currentNode
			currentNode = currentNode.Next
			currentPosition++
			continue
		}
	}

}

func (list *LinkedList[T]) DeleteAt(index int) T {
	var deletedValue T
	if index >= list.Length || index < 0 {
		fmt.Println("can't delete element at position out of list's range")
		return deletedValue
	}

	prev := &Node[T]{} // empty previous node to keep track of it
	currentNode := list.Head

	if index == 0 {
		prev = list.Head
		list.Head = currentNode.Next
		return prev.Value
	}

	currentPosition := 0
	for {
		if currentPosition == index {
			prev.Next = currentNode.Next
			list.Length--
			return currentNode.Value
		} else {
			prev = currentNode
			currentNode = currentNode.Next
			currentPosition++
			continue
		}
	}
}

func (list *LinkedList[T]) Show() {
	if list.Length == 0 {
		fmt.Println("[ ]")
	}

	if list.Length == 1 {
		list.getTail().print(list.Length, "", true)
	}

	if list.Length > 1 {
		currentNode := *list.Head

		fmt.Print("[")
		currentNode.print(list.Length, "", true)

		for currentNode.Next != nil {
			currentNode = *currentNode.Next
			currentNode.print(list.Length, ", ", false)
		}
		fmt.Print("]\n")
	}
}

func (list *LinkedList[T]) ShowAt(index int) {
	if index >= list.Length || index < 0 {
		fmt.Println("can't show value at position out of list's range")
		return
	}

	if list.Length == 0 {
		fmt.Println("[ ]")
	}

	currentPosition := 0

	currentNode := *list.Head

	fmt.Print("[")

	for currentNode.Next != nil {
		if currentPosition == index {
			currentNode.print(list.Length, "", true)
			break
		}

		currentNode = *currentNode.Next
		currentPosition++
	}
	fmt.Print("]\n")
}

func (node *Node[T]) print(listLength int, delimiter string, isFirst bool) {
	// print the node in the regular list way [val1, val2, val3, ...]

	if listLength > 1 {
		if isFirst {
			fmt.Print(node.Value)
		} else {
			fmt.Print(delimiter, node.Value)
		}
	} else {
		fmt.Println("[", node.Value, "]")
	}
}

func (list *LinkedList[T]) getTail() *Node[T] {
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
