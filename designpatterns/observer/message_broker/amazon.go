package main

type Amazon struct {
	broker *Broker
}

func (a *Amazon) produceMessage(topic, message string) {
	a.broker.produceMessage(topic, message)
}

func (a *Amazon) registerBroker(broker *Broker) {
	a.broker = broker
}
