package main

// https://leetcode.com/problems/print-zero-even-odd/description/

import (
	"fmt"
	"sync"
)

type Message struct {
	terminate bool
}

type ZeroEvenOdd struct {
	n        int
	zeroChan chan Message
	evenChan chan Message
	oddChan  chan Message
	i        int
	wg       sync.WaitGroup
}

func NewZeroEvenOdd(n int) *ZeroEvenOdd {
	st := ZeroEvenOdd{
		n:        n,
		zeroChan: make(chan Message),
		evenChan: make(chan Message),
		oddChan:  make(chan Message),
		i:        1,
	}
	f := func(x int) {
		fmt.Printf("%d ", x)
	}
	go st.even(f)
	go st.odd(f)
	go st.zero(f)
	st.zeroChan <- Message{false}
	st.wg.Add(3)
	return &st
}

func (z *ZeroEvenOdd) zero(f func(int)) {
	for {
		select {
		case message := <-z.zeroChan:
			if message.terminate {
				// fmt.Println("received terminate on zero channel returning")
				z.wg.Done()
				return
			}
			// fmt.Println("Received message in zero chan")
			if z.i > z.n {
				// fmt.Println("i > n in zero chan returning")
				z.oddChan <- Message{true}
				z.evenChan <- Message{true}
				z.wg.Done()
				return
			}
			f(0)
			if z.i%2 == 0 {
				// fmt.Println("Pushing into even chan")
				z.evenChan <- Message{false}
			} else {
				// fmt.Println("Pushing into odd chan")
				z.oddChan <- Message{false}
			}
		}
	}
}

func (z *ZeroEvenOdd) even(f func(int)) {
	for {
		select {
		case message := <-z.evenChan:
			if message.terminate {
				// fmt.Println("received terminate on even channel returning")
				z.wg.Done()
				return
			}
			// fmt.Println("Received message in even chan")
			if z.i > z.n {
				// fmt.Println("i > n in even chan returning")
				z.zeroChan <- Message{true}
				z.oddChan <- Message{true}
				z.wg.Done()
				return
			}
			f(z.i)
			z.i++
			z.zeroChan <- Message{false}
		}
	}
}

func (z *ZeroEvenOdd) odd(f func(int)) {
	for {
		select {
		case message := <-z.oddChan:
			if message.terminate {
				// fmt.Println("received terminate on odd channel returning")
				z.wg.Done()
				return
			}
			// fmt.Println("Received message in odd chan")
			if z.i > z.n {
				// fmt.Println("i > n in odd chan returning")
				z.zeroChan <- Message{true}
				z.evenChan <- Message{true}
				z.wg.Done()
				return
			}
			f(z.i)
			z.i++
			z.zeroChan <- Message{false}
		}
	}
}

func main() {
	st := NewZeroEvenOdd(15)
	st.wg.Wait()
}
