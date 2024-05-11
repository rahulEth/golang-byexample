package main

import (
	"fmt"
	"time"
)

func main() {

	// switch with condition
	// switch os := runtime.GOOS; os {
	// case "windows":
	// 	fmt.Println("windows os")

	// case "linux":
	// 	fmt.Println("linux")
	// default:
	// 	fmt.Println("kuch nhi")
	// }
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("good morning")
	case t.Hour() < 17:
		fmt.Println("good evening")

	default:
		fmt.Println(".....")
	}
}
