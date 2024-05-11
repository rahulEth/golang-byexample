package main

import "fmt"

func main() {
	//// Using empty interface to hold values of different types
	var i interface{}

	i = "hello"
	describeee(i)
	i = 42
	describeee(i)
}

func describeee(i interface{}) {
	fmt.Printf("%v, %T \n", i, i)
}
