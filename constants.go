package main

import "fmt"

const m string = "constant"

// func main() {
// 	const e = 'a'
// 	fmt.Printf("Tpye: %T value: %v\n", e, e)
// 	var f1 float32
// 	var f2 float64
// 	f1 = math.Pi
// 	f2 = math.Pi
// 	fmt.Printf("Type: %T value: %v\n", math.Pi, math.Pi)
// 	fmt.Printf("Type: %T value: %v\n", f1, f1)
// 	fmt.Printf("Tpye: %T value: %v\n", f2, f2)

// }

func main() {
	type myChar int32
	const aa int32 = 'a'
	var b myChar = aa
	var c = aa
	fmt.Println(b, c)

	const m = map[string]int{
		"name": 2,
	}

}
