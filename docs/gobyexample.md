coinGo CLI: 
Go is strictly call by value language. Even when you pass a pinter inside the function a copy of the pointer is created.

```package main

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
}```


But maps and slices are curious case in GO. map and slice implemented as data stricture that hold memory addresses.(they are reference type )
Go  build: to compile the one or more file in same directory
Go run : to compile and run one or more fle in same directory
Go install: compile and install package
Go get: to download row code of another package
Go fmt: to format the coe
Go test:  to run the test

Import statement: to gain access of another packages inside of one we are creating

Basic Types in GO:

```String
bool
Int
float64``` 
