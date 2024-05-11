package main

import "fmt"

type insuffError struct {
	val int
}

func (e insuffError) Error() string {
	return fmt.Sprintf("Insufficient funds you need at least %d funds\n", e.val)
}

func withdraw(balance, amout int) error {
	if balance < amout {
		return insuffError{val: amout}
	}
	return nil
}

func main() {
	balanace := 100
	withdrawl := 150
	err := withdraw(balanace, withdrawl)
	if err != nil {
		fmt.Println("Error ", err)
	} else {
		fmt.Println("withdrawl successful")
	}
}
