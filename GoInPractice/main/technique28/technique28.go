package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

type MyWriter struct {
}

func (m *MyWriter) Write([]byte) error {
	return nil
}

func Pad(s string, max uint) string {
	log.Printf("Testing len: %d, Str: %s\n", max, s)
	ln := uint(len(s))
	if ln > max {
		return s[:max]
	}
	s += strings.Repeat(" ", int(max-ln))
	return s
}

func main() {
	m := map[string]interface{}{
		"w": &MyWriter{},
	}

	doSomething(m)
}

func doSomething(m map[string]interface{}) {
	w := m["w"].(io.Writer)
	fmt.Println(w)
}
