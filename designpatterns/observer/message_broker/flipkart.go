package main

type Flipkart struct {
	broker *Broker
}

func (f *Flipkart) produceMessage(topic, message string) {
	f.broker.produceMessage(topic, message)
}

func (f *Flipkart) registerBroker(broker *Broker) {
	f.broker = broker
}
