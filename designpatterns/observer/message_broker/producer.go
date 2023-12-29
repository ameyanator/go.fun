package main

type Producer interface {
	produceMessage(topic, message string)
	registerBroker(*Broker)
}
