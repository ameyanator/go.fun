package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Consumer struct {
	topic    *Topic
	offset   int
	id       int
	stopChan chan struct{}
}

func NewConsumer(topic *Topic, id int) *Consumer {
	return &Consumer{
		topic:    topic,
		offset:   0,
		id:       id,
		stopChan: make(chan struct{}),
	}
}

func (c *Consumer) ConsumeFromTopic() {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			message, err := c.topic.getMessage(c.offset)
			if err != nil {
				fmt.Println("Consumer ", c.id, " Got error ", err.Error())
				continue
			}
			fmt.Println("Consumer ", c.id, " Got message from topic - ", message)
			randVal := rand.Intn(5)
			fmt.Println("Consumer ", c.id, " sleeping for ", randVal, " secs")
			time.Sleep(time.Duration(randVal) * time.Second)
			fmt.Println("Consumer ", c.id, " processed message from topic - ", message)
			c.offset++
		case <-c.stopChan:
			fmt.Println("Stopping consumer ", c.id)
			return
		}
	}
}
