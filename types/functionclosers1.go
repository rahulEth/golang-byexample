package main

import "fmt"

// outer function calls the innerFuc that add sum to input x
func outerFun() func(int) int {
	sum := 0
	innerFunc := func(x int) int {
		sum += x
		return sum
	}
	return innerFunc

}

func main() {
	// create the two instance of outerFunc each will contain surrouding lexical scope indepently
	pos, neg := outerFun(), outerFun()

	for i := 0; i < 10; i++ {
		// Call the pos closure with the loop index.
		// This adds the loop index to the positive sum.
		// Also, call the neg closure with a negative multiple of the loop index.
		// This subtracts twice the loop index from the negative sum.
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
