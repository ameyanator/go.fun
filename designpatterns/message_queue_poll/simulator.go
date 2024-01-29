package main

import (
	"fmt"
	"time"
)

func main() {
	topic1 := NewTopic()

	consumer1 := NewConsumer(topic1, 1)
	consumer2 := NewConsumer(topic1, 2)

	go consumer1.ConsumeFromTopic()
	go consumer2.ConsumeFromTopic()

	fmt.Println("Sending first message")
	topic1.addMessage("First message")
	// fmt.Println("Main sleeping for 2 secs")
	time.Sleep(2 * time.Second)

	fmt.Println("Sending Second message")
	topic1.addMessage("Second message")
	// fmt.Println("Main sleeping for 2 secs")
	time.Sleep(2 * time.Second)

	time.Sleep(5 * time.Second)
	consumer1.stopChan <- struct{}{}
	time.Sleep(5 * time.Second)
}
