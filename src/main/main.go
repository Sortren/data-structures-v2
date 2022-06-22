package main

import (
	"fmt"
	"src/pkg/models"
)

func main() {
	list := models.LinkedList{}

	for i := 1; i <= 10; i++ {
		list.Push(i)
	}

	list.Print()

	fmt.Println()
	fmt.Println()

	list.PrintExplicitly()
}
