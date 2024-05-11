package main

import (
	"fmt"
)

func main() {
	x, y := 3, 4
	var z uint32 = uint32(x)
	var a float32 = float32(x)
	fmt.Println(x, y, z, a)
}
