package main

// func add(x, y int) int {
// 	return x + y
// }

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - 4
	return
}
func main() {
	a, b := swap("hello", "world")
	println(a, b)
	x, y := split(16)
	fmt.Println(x, y)
}
