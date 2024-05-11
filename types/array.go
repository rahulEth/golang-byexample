package main

import "fmt"

func main() {
	var a [10]string
	a[0] = "name1"
	a[1] = "name2"
	fmt.Println(a, a[0], a[1])
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)
}
