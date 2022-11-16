package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

/*
Idiomatic Go uses a method that’s both more efficient and simpler: create
errors as package-scoped variables and reference those variables.
*/

var ErrTimeout = errors.New("The request timed out")
var ErrRejected = errors.New("The request was rejected")

var random = rand.New(rand.NewSource(int64(time.Now().Second())))

func main() {
	response, err := SendRequest("Hello")
	for err == ErrTimeout {
		fmt.Println("Timeout, Retrying")
		response, err = SendRequest("Hello")
	}
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response)
	}
}

func SendRequest(req string) (string, error) {
	switch random.Int() % 3 {
	case 0:
		return "Success", nil
	case 1:
		return "", ErrRejected
	default:
		return "", ErrTimeout
	}
}
