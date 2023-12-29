package main

import (
	"time"
)

func main() {
	broker := NewBroker()
	producerFactory := NewProducerFactory()
	consumerFactory := NewConsumerFactory()

	amazon := producerFactory.getProducer("Amazon")
	flipkart := producerFactory.getProducer("Flipkart")
	amazon.registerBroker(broker)
	flipkart.registerBroker(broker)

	ameya := consumerFactory.getConsumer(1997)
	arya := consumerFactory.getConsumer(2005)
	ameya.setBroker(broker)
	arya.setBroker(broker)

	// fmt.Println(ameya, arya)

	ameya.registerTopicWithBroker("f1")
	arya.registerTopicWithBroker("f1")
	ameya.registerTopicWithBroker("chess")

	arya.goOffline()

	amazon.produceMessage("f1", "red bull")
	time.Sleep(2 * time.Second)

	flipkart.produceMessage("f1", "aston martin")
	amazon.produceMessage("chess", "magnus carlsen")
	time.Sleep(2 * time.Second)

	arya.goOnline()
	time.Sleep(10 * time.Second)

	amazon.produceMessage("f1", "haas")
	time.Sleep(2 * time.Second)
}
