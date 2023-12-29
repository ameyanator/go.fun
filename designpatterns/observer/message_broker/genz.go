package main

import "fmt"

type GenZ struct {
	broker *Broker
	flag   bool
}

func (g *GenZ) asyncConsumeMessages(topic, message string) bool {
	if g.flag {
		go g.consumeMessage(topic, message)
		return true
	}
	return false
}

func (g *GenZ) registerTopicWithBroker(topic string) {
	g.broker.registerConsumer(topic, g)
}

func (g *GenZ) consumeMessage(topic, message string) {
	fmt.Println("For a genz, On topic ", topic, " we have message ", message)
}

func (g *GenZ) rejoin() {
	g.broker.getNewMessages(g)
}

func (g *GenZ) setBroker(broker *Broker) {
	g.broker = broker
}

func (g *GenZ) goOnline() {
	g.flag = true
	g.rejoin()
}

func (g *GenZ) goOffline() {
	g.flag = false
}
