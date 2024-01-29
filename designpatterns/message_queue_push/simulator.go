package main

import (
	"fmt"
	"time"
)

func main() {
	topic := NewTopic()

	NewConsumer(1, topic)
	consumer2 := NewConsumer(2, topic)

	// consumer2.consumerUp = false

	topic.addMessage("First Message")
	fmt.Println("Sleeping for 2 secs")
	time.Sleep(2 * time.Second)

	topic.addMessage("Second Message")
	fmt.Println("Sleeping for 2 secs")
	time.Sleep(2 * time.Second)

	consumer2.consumerUp = true

	consumer2.getNewMessagesInTopic()
	time.Sleep(2 * time.Second)
}
