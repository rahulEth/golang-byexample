package main

import "fmt"

type MyFloat float64

// method can be declared on non struct type
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(4)
	fmt.Println(f.Abs())
}
