package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y float64
}

func (v Vertex) abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}
func (v *Vertex) Scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}
func main() {
	v := Vertex{2, 3}
	v.Scale(10)
	fmt.Println(v.abs())

}
