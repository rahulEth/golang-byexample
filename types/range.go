package main

import "fmt"

func main() {
	pow := make([]int, 10)

	for i := range pow {
		pow[i] = i << uint(i) // 2**i
	}
	for _, val := range pow {
		fmt.Println("valu...... ", val)
	}

}
