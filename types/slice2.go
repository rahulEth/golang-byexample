package main

import "fmt"

func main() {
	a := make([]int, 5)
	fmt.Println(a, len(a), cap(a))
	b := make([]int, 2, 5)
	fmt.Println(b, len(b), cap(b))
	c := b[:2]
	fmt.Println(c, len(c), cap(c))

	d := c[2:5]

	fmt.Println(d, len(d), cap(d))

	var s []int
	fmt.Println(s, len(s), cap(s))

	if s == nil {
		fmt.Println("nil")
	}

}
