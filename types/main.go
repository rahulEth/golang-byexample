package main

import "fmt"

func main() {
	a := 10
	aPtr := &a
	increment(aPtr)
	fmt.Println(aPtr) // becaue go is strictly call by value thats why it stil hold the addreess
}

// a copy of pointer of a will be created in increment function
func increment(x *int) {
	x = nil
}
