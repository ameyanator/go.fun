package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1902")
	if err != nil {
		panic("Failed to connect to localhost")
	}
	defer conn.Close()
	/*
		It’s always recommended to close a network connection in a defer block. If nothing else, when a panic
		occurs (as it will in this demo code), the network buffer will be flushed on close, and
		you’re less likely to lose critical log messages telling you why the code panicked.
	*/

	f := log.Ldate | log.Lshortfile
	logger := log.New(conn, "example-", f)

	logger.Println("This is a regular message")
	/*
		the log.Fatal* functions have an unfortunate
		side effect: the deferred function isn’t called. Why not? Because log.Fatal* calls
		os.Exit, which immediately terminates the program without unwinding the function
		stack. We covered this topic in the preceding chapter. Because the deferred function
		is skipped, your network connection is never properly flushed and closed. Panics, on
		the other hand, are easier to capture. In reality, production code for anything but sim-
		ple command-line clients should avoid using fatal errors.
	*/
	logger.Panicln("This is a panic")
}
