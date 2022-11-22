package main

import (
	"testing"
	"testing/quick"
)

func TestWriter(t *testing.T) {
	// var _ io.Writer = &MyWriter{}
	// This is a simple test. You don’t even have to run the test to cause it to fail. The compiler will fail before the binary can ever be built:
}

func TestPadGenerative(t *testing.T) {
	fn := func(s string, max uint8) bool {
		p := Pad(s, uint(max))
		return len(p) == int(max)
	}
	if err := quick.Check(fn, &quick.Config{MaxCount: 200}); err != nil {
		t.Error(err)
	}
}
