package main

import (
	"fmt"
	"time"
)

func InfiniteLoop(interval time.Duration, stopChan chan struct{}) {
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-ticker.C:
			fmt.Println("Got something on the ticker")
		case <-stopChan:
			fmt.Println("Got something on stop channel closing goroutine")
			return
		}
	}
}

func main() {
	stopChannel := make(chan struct{})
	go InfiniteLoop(2*time.Second, stopChannel)
	time.Sleep(10 * time.Second)
	close(stopChannel)
	time.Sleep(5 * time.Second)
}
