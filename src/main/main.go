package main

import (
	"fmt"
	"src/pkg/models"
)

func main() {
	list := models.LinkedList{}

	for i := 1; i <= 5; i++ {
		list.Append(i)
	}

	for i := 10; i >= 5; i-- {
		list.Push(i)
	}

	list.Print()

	fmt.Println()
	fmt.Println()

	fmt.Println("Popping value")
	val := list.Pop()
	fmt.Println(val)

	fmt.Println()

	list.PrintExplicitly()
	fmt.Println(list.GetTail())
}
