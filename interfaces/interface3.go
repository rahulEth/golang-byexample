package main

import "fmt"

type I interface {
	M()
}

func M() {
}
func main() {
	var i I
	describee(i)
	i.M()
}

func describee(i I) {
	fmt.Println("(%v, %T)\n ", i, i)
}
