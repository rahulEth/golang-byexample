package main

import "fmt"

type Vertex struct {
	lat, long float64
}

var m = map[string]Vertex{
	"bell labs": {23.4566, -34.000},
	"google": {
		21.00, -24.000,
	},
}

func main() {
	fmt.Println(m)
	fmt.Println(m["bell labs"])
}
