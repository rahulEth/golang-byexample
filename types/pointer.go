package main

import "fmt"

func main() {
	// Declare a variable of type int
	var num int = 10

	// Declare a pointer variable of type *int (pointer to int)
	var ptr *int
	fmt.Println("ptr default value ", ptr) // nil
	// Assign the address of 'num' to the pointer variable 'ptr'
	ptr = &num

	// Print the value of 'num' and the value pointed to by 'ptr'
	fmt.Println("Value of num:", num)             // Output: Value of num: 10
	fmt.Println("Value pointed to by ptr:", *ptr) // Output: Value pointed to by ptr: 10

	// Change the value of 'num' using the pointer
	*ptr = 20 //This is known as "dereferencing" or "indirecting".

	// Print the updated value of 'num'
	fmt.Println("Updated value of num:", num) // Output: Updated value of num: 20
}
