package main

import "fmt"

func main() {
	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)
	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)
	j := i.(float64) // panic
	fmt.Println(j)
}
