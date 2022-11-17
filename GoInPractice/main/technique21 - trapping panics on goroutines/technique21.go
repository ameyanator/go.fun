package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"
)

func response(data []byte, conn net.Conn) {
	// defer func() {
	// 	conn.Close()
	// }()
	conn.Write(data)
	panic(errors.New("Failure in response!"))
}

func handle(conn net.Conn) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Fatal Error: %s\n", err)
		}
		conn.Close()
	}()
	reader := bufio.NewReader(conn)
	data, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Println("Failed to read from socket")
		conn.Close()
	}
	response(data, conn)
}

func listen() {
	listner, err := net.Listen("tcp", ":1026")
	if err != nil {
		fmt.Println("Failed to open port on 1026")
		return
	}

	for {
		conn, err := listner.Accept()
		if err != nil {
			fmt.Println("Error accepting connection")
			continue
		}
		go handle(conn)
	}
}

func technique1() {
	listen()
}

func main() {
	technique1()
}
