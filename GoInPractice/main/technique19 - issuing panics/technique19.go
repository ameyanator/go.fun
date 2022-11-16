package main

import (
	"errors"
	"fmt"
)

/*
Panics, on the other hand, are unexpected. They occur when a constraint or limitation
is unpredictably surpassed. When it comes to declaring a panic in your code,
the general rule of thumb is don’t panic unless there’s no clear way to handle the 
condition within the present context. When possible, return errors instead.
*/

var ErrDivideByZero = errors.New("Can't divide by zero")

func main() {
	fmt.Println("Divide 1 by 0")
	_, err := precheckDivide(1, 0)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	fmt.Println("Divide 2 by 0")
	divide(2, 0)
}

func precheckDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return divide(a, b), nil
}

func divide(a, b int) int {
	return a / b
}
