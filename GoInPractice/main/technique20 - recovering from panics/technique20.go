package main

import "fmt"

func main() {
	defer goodbye()
	fmt.Println("Hello World")
}

func goodbye() {
	fmt.Println("Goodbye")
}
