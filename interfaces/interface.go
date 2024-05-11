package main

import (
	"fmt"
	"math"
)

type Absr interface {
	Abs() float64
}
type Vertex struct {
	x, y float64
}

type myFloat float64

func main() {
	var d Absr
	f := myFloat(-2)
	v := Vertex{2, 3}

	d = f  // a myFloat implentes interface Absr
	d = &v // a *Vertex implement interface Absr
	// v is vertex not *vertex and does not implement Absr
	// a = v; // a is vertex and dont implement Absr
	fmt.Println(d.Abs())
}

func (a myFloat) Abs() float64 {
	if a < 0 {
		return float64(-a)
	}
	return float64(a)
}

func (a *Vertex) Abs() float64 {
	return math.Sqrt(a.x*a.x + a.y*a.y)
}
