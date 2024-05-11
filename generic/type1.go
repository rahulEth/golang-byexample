package main

import "fmt"

func Index[T comparable](s []T, x T) int {
	// v and x are type T which has comparable
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	si := []int{1, 2, 3, 4}
	fmt.Println(Index(si, 3))

	ss := []string{"foo", "bar", "buz"}
	fmt.Println(Index(ss, "hello"))

}
