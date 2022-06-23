package main

import (
	"fmt"
	"src/pkg/models"
)

func main() {
	integerList := models.LinkedList[int]{}

	for i := 1; i <= 5; i++ {
		integerList.Append(i)
	}

	for i := 10; i >= 5; i-- {
		integerList.Push(i)
	}

	integerList.Print()

	fmt.Println()
	fmt.Println()

	fmt.Println("Popping value")
	intVal := integerList.Pop()
	fmt.Println(intVal)

	fmt.Println()

	integerList.PrintExplicitly()
	fmt.Println(integerList.GetTail())

	fmt.Println()
	fmt.Println()

	stringList := models.LinkedList[string]{}

	for i := 1; i <= 5; i++ {
		stringList.Append(fmt.Sprintf("%s%d", "a", i))
	}

	for i := 10; i >= 5; i-- {
		stringList.Push(fmt.Sprintf("%s%d", "b", i))
	}

	stringList.Print()

	fmt.Println()
	fmt.Println()

	fmt.Println("Popping value")
	strVal := stringList.Pop()
	fmt.Println(strVal)

	fmt.Println()

	stringList.PrintExplicitly()
	fmt.Println(stringList.GetTail())

	fmt.Println()
	fmt.Println()

}
