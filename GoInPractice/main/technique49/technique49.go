package main

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
)

/*
Connection timeouts are a common problem and useful to detect. For example, if a
timeout error occurs, especially if it’s in the middle of a connection, retrying the oper-
ation might be worthwhile. On retry, the server you were connected to may be back
up, or you could be routed to another working server.
*/

/*
Connection timeouts are a common problem and useful to detect. For example, if a
timeout error occurs, especially if it’s in the middle of a connection, retrying the oper-
ation might be worthwhile. On retry, the server you were connected to may be back
up, or you could be routed to another working server.
*/

// When timeouts occur, a small variety of errors occurs. Check the error for each of these cases to see if it was a timeout.

func hasTimedOut(err error) bool {
	switch err := err.(type) {
	case *url.Error:
		//url.Error may be caused by an underlying net error that can be checked for a timeout
		if err, ok := err.Err.(net.Error); ok && err.Timeout() {
			return true
		}
	case net.Error:
		// Look for timeout detected by net package
		if err.Timeout() {
			return true
		}
	case *net.OpError:
		// Look for timeout detected by net package
		if err.Timeout() {
			return true
		}
	}
	errTxt := "use of closed network connection:"
	// check without custom Timeout set

	if err != nil && strings.Contains(err.Error(), errTxt) {
		return true
	}
	return false
}

func main() {
	res, err := http.Get("http://example.com/test.zip")
	if err != nil && hasTimedOut(err) {
		fmt.Println("Timeout occured")
		return
	}
	fmt.Println("Got response", res)
}
