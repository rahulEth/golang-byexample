package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["answer"] = 32
	fmt.Println("the value: ", m["answer"])
	m["answer"] = 42
	fmt.Println("the value: ", m["answer"])
	delete(m, "answer")
	fmt.Println("the value: ", m["answer"])
	v, ok := m["answer"]
	fmt.Println("The value: ", v, "exist? ", ok)
}
