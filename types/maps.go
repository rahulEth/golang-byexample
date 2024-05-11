package main

import "fmt"

type Vertex struct {
	lat, long float64
}

var m map[string]Vertex // nil map

func main() {
	m = make(map[string]Vertex)
	m["bela vita"] = Vertex{
		40.0344, -43.0877,
	}
	fmt.Println(m["bela vita"])
}
