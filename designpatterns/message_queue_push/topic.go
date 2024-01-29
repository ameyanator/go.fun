package main

import "sync"

type Topic struct {
	messages               []string
	consumers              []*Consumer
	lastMessagePerConsumer map[*Consumer]int
	messageMu              sync.Mutex
	consumerSendMu         map[*Consumer]sync.Mutex
	consumerMu             sync.Mutex
}

func NewTopic() *Topic {
	return &Topic{
		messages:               make([]string, 0),
		consumers:              make([]*Consumer, 0),
		lastMessagePerConsumer: make(map[*Consumer]int),
		consumerSendMu:         make(map[*Consumer]sync.Mutex),
	}
}

func (t *Topic) addConsumer(consumer *Consumer) {
	t.consumerMu.Lock()
	defer t.consumerMu.Unlock()

	t.consumers = append(t.consumers, consumer)
}

func (t *Topic) addMessage(message string) {
	t.messageMu.Lock()
	t.messages = append(t.messages, message)
	t.messageMu.Unlock()

	for _, consumer := range t.consumers {
		go t.sendNewMessagesToConsumer(consumer)
	}
}

func (t *Topic) sendNewMessagesToConsumer(consumer *Consumer) {
	mu := t.consumerSendMu[consumer]
	mu.Lock()
	defer mu.Unlock()

	for i := t.lastMessagePerConsumer[consumer]; i < len(t.messages); i++ {
		messageSent := consumer.receiveMessage(t.messages[i])
		if messageSent {
			t.lastMessagePerConsumer[consumer] = i + 1
		} else {
			break
		}
	}
}
