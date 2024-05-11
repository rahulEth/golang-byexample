package main

import (
	"fmt"
	"time"
)

func sayHello(s string) {
	for i := 0; i < 5; i++ {
		// sleep for a short duration to allow the goroutie to execute
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// call sayHello as go routine
	go sayHello("hello")

	sayHello("world")

}
