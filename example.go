package main

import (
	"fmt"
	"time"
)

func sayHello(s string, ch chan bool) {
	for i := 0; i < 5; i++ {
		fmt.Println("\n", s)
		time.Sleep(100 * time.Millisecond)
	}

	ch <- true // signals that the goroutine is done
}
func sayGoodbye(ch chan bool) {
	for i := 0; i < 5; i++ {
		fmt.Println("Goodbye")
		time.Sleep(100 * time.Millisecond)
	}
	ch <- true // singals that the goroutine is done
}
func main() {

	done := make(chan bool, 2) // creating the channel to signal completetion
	go sayHello("hello", done)

	go sayGoodbye(done)

	// Wait for both goroutines to complete

	<-done
	<-done

	fmt.Println("main function ends")
}
