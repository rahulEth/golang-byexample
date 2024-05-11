package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Printf("%T, %v\n", v, v)
	case int:
		fmt.Printf("%T, %v\n", v, v)
	default:
		fmt.Printf("i do not know about type %T\n", v)

	}
}

func main() {
	do("hello")
	do(21)
	do(true)
}
