package main

import (
	"src/pkg/models"
)

func main() {
	list := models.LinkedList{}

	list.Push(20)
	list.Push(30) // next -> nil
	list.Push(40)

	list.Print()
}
