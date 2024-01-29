package main

import "fmt"

type Consumer struct {
	consumerUp bool
	id         int
	topic      *Topic
}

func NewConsumer(id int, topic *Topic) *Consumer {
	consumer := Consumer{
		consumerUp: true,
		id:         id,
		topic:      topic,
	}
	topic.addConsumer(&consumer)
	return &consumer
}

func (c *Consumer) receiveMessage(message string) bool {
	if c.consumerUp {
		fmt.Println("Consumer ", c.id, "received message - ", message)
		return true
	}
	return false
}

func (c *Consumer) getNewMessagesInTopic() {
	c.topic.sendNewMessagesToConsumer(c)
}
