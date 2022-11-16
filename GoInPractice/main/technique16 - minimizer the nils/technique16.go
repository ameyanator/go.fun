package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if result, err := Concat(args...); err != nil {
		fmt.Printf("Error %s\n", err)
	} else {
		fmt.Printf("Concatenated string %s\n", result)
	}
}

func Concat(parts ...string) (string, error) {
	if len(parts) == 0 {
		return "", errors.New("No String Supplied") // Its a good idea to not always return nil if possible and return a 0 value so users who dont do defensive programming can also use it
	}
	return strings.Join(parts, "/"), nil
}
