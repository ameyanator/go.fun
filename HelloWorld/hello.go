package main

import "fmt"

func Hello(name string) string {
	return "Namaste " + name + "!"
}

func main() {
	fmt.Println(Hello("World"))
}