package main

import "fmt"

// outer function retrun innner function that captures the varibale x
func outerFunc() func() int {
	x := 0
	// innerFunction capute the variable x from outer Function
	innerFunc := func() int {
		x++
		return x
	}
	return innerFunc
}

func main() {

	// creating a clouser by callig outerFunction
	closuer := outerFunc()
	// calling the 	closure multiple times

	fmt.Println(closuer())
	fmt.Println(closuer())
	fmt.Println(closuer())
}
