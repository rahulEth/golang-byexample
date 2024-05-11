package main

import "fmt"

type Vertex struct {
	x, y float64
}

func (v *Vertex) Scalee(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.x = v.x * f
	v.x = v.x * f
}
func main() {
	v := Vertex{2, 3}
	v.Scalee(10)
	ScaleFunc(&v, 10)
	p := &Vertex{2, 3}
	p.Scalee(10)
	ScaleFunc(p, 10)

	fmt.Println(v, p)
}
