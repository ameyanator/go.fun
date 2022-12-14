package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

/*
here are a few useful guidelines for working with deferred functions:
 Put deferred functions as close to the top of a function declaration as possible.
 Simple declarations such as foo := 1 are often placed before deferred functions.
 More-complex variables are declared before deferred functions (var myFile
io.Reader), but not initialized until after.
 Although it’s possible to declare multiple deferred functions inside a function,
this practice is generally frowned upon.
 Best practices suggest closing files, network connections, and other similar
resources inside a defer clause. This ensures that even when errors or panics
occur, system resources will be freed.
*/

func technique1() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Trapped panic: %s (%T)\n", err, err)
		}
	}()
	yikes()
}

func technique2() {
	var file io.ReadCloser
	file, err := OpenCSV("data.csv")
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	defer file.Close()
	// Do something with the file
}

func OpenCSV(filename string) (file *os.File, err error) {
	defer func() {
		if r := recover(); r != nil {
			file.Close()
			err = r.(error)
		}
	}()
	file, err = os.Open(filename)
	if err != nil {
		fmt.Printf("Failed to open file\n")
		return file, err
	}
	RemoveEmptyLines(file)
	return file, err
}

func RemoveEmptyLines(file *os.File) {
	panic(errors.New("Failed parse\n"))
}

func main() {
	technique2()
}

func yikes() {
	panic(errors.New("Something Bad Happened"))
}
