package main

// List represents a singly-linked list that holds any type
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {

}
