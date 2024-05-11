package main

import "fmt"

func add(a, b int) int {
	return a + b
}

// functiion that take a function as an argument

func applyOperation(a, b int, operation func(a, b int) int) int {
	return operation(a, b)
}

func main() {
	// Assigning the add function to a variable
	var addFunc func(a, b int) int
	addFunc = add

	// using the addFunc variable to call the add function
	sum := addFunc(2, 3)

	fmt.Println("sum ", sum)

	// passing add function as an argument to applyOperation function
	result := applyOperation(2, 3, add)
	fmt.Println("result: ", result)
}
