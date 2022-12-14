package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

func main() {
	// technique1()
	technique2()
	/*
		you must supply a presized buffer. But there’s no convenient way
		to determine how big the buffer needs to be to capture all of the output. (And in
		some cases, the output is so big that you might not want to capture it all.) You need to
		decide ahead of time how much space you’d like to allocate.
	*/
	/*
		Stack takes two arguments. The second is a Boolean flag, which is set to
		false in this example. Setting it to true will cause Stack to also print out stacks for all
		running goroutines. This can be tremendously useful when debugging concurrency
		problems, but it substantially increases the amount of output.
	*/
}

func technique2() {
	bar2()
}

func technique1() {
	bar()
}

func bar() {
	debug.PrintStack()
}

func bar2() {
	buf := make([]byte, 1024)
	runtime.Stack(buf, false)
	fmt.Printf("Trace: \n %s \n", buf)
}

/*
If all of this isn’t sufficient, you can use the runtime package’s Caller and Callers
functions to get programmatic access to the details of the call stack. Although it’s
quite a bit of work to retrieve and format the data, these functions give you the flexi-
bility to discover the details of a particular call stack. Both the runtime and the
runtime/debug packages contain numerous other functions for analyzing memory
usage, goroutines, threading, and other aspects of your program’s resource usage.
*/
