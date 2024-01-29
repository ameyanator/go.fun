package main

import (
	"errors"
	"sync"
)

type Topic struct {
	messages []string
	mu       sync.Mutex
}

func NewTopic() *Topic {
	return &Topic{
		messages: make([]string, 0),
	}
}

func (t *Topic) addMessage(message string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.messages = append(t.messages, message)
}

func (t *Topic) getMessage(offset int) (string, error) {
	if offset >= len(t.messages) {
		return "", errors.New("Message with specified offset does not exist")
	}
	return t.messages[offset], nil
}
