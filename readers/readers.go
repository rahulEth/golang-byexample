package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	// crearting a reader from a string
	reader := strings.NewReader("hello world!")
	// create a buffer to hold the read data
	buffer := make([]byte, 5)

	for {
		// read data from reader into the buffer
		n, err := reader.Read(buffer)
		if err != nil {
			fmt.Println("Error reading: ", err)
			return
		}

		fmt.Printf("Read %d bytes: %s\n", n, buffer[:n])
		// when reader stream ends returns io.EOF error
		if err == io.EOF {
			break
		}
	}

}
