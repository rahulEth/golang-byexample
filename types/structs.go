package main

import "fmt"

// type Str struct {
// 	num  int
// 	name string
// }

type Str1 struct {
	X int
	Y int
}

func main() {
	// fmt.Println(Str{2, "myname"})
	v := Str1{1, 2}
	v.Y = 10
	fmt.Println(v.Y)
	p := &v
	p.X = 22 // or  (*p).X = 22;
	fmt.Println(p.X, v)

	v0 := Str1{1, 2}

	v1 := Str1{X: 1}
	v2 := Str1{}
	q := &Str1{1, 2}
	fmt.Println(v0, v1, v2, q)
}
