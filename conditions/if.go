package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	if x < 0 {
		return math.Sqrt(-x)
	}

	return math.Sqrt(x)
}

func pow(x, y, lim float64) float64 {
	if v := math.Pow(x, y); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(pow(2, 3, 10), pow(3, 3, 20))
}
