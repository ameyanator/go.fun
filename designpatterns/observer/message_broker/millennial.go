package main

import "fmt"

type Millennial struct {
	broker *Broker
	flag   bool
}

func (m *Millennial) asyncConsumeMessages(topic, message string) bool {
	m.consumeMessage(topic, message)
	return true
}

func (m *Millennial) registerTopicWithBroker(topic string) {
	m.broker.registerConsumer(topic, m)
}

func (m *Millennial) consumeMessage(topic, message string) {
	fmt.Println("For a millenial, On topic ", topic, " we have message ", message)
}

func (m *Millennial) rejoin() {
	m.broker.getNewMessages(m)
}

func (m *Millennial) setBroker(broker *Broker) {
	m.broker = broker
}

func (m *Millennial) goOnline() {
	m.flag = true
	m.rejoin()
}

func (m *Millennial) goOffline() {
	m.flag = false
}
