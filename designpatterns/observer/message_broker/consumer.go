package main

type Consumer interface {
	asyncConsumeMessages(topic, message string) bool
	registerTopicWithBroker(topic string)
	setBroker(*Broker)
	goOnline()
	goOffline()
}
