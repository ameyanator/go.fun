package main

type Broker struct {
	consumers           map[string][]Consumer
	messages            map[string][]string
	lastConsumedMessage map[string]map[Consumer]int //for particular topic and consumer, get last consumed message
}

func NewBroker() *Broker {
	return &Broker{
		consumers:           make(map[string][]Consumer),
		messages:            make(map[string][]string),
		lastConsumedMessage: make(map[string]map[Consumer]int),
	}
}

func (b *Broker) registerConsumer(topic string, consumer Consumer) {
	b.consumers[topic] = append(b.consumers[topic], consumer)
	if _, present := b.lastConsumedMessage[topic]; present == false {
		b.lastConsumedMessage[topic] = make(map[Consumer]int)
		b.lastConsumedMessage[topic][consumer] = 0
	} else if _, present := b.lastConsumedMessage[topic][consumer]; !present {
		b.lastConsumedMessage[topic][consumer] = 0
	}
}

func (b *Broker) unregisterConsumer(topic string, consumer Consumer) {
	for i, c := range b.consumers[topic] {
		if c == consumer {
			b.consumers[topic] = append(b.consumers[topic][:i], b.consumers[topic][i+1:]...)
			break
		}
	}
}

func (b *Broker) produceMessage(topic, message string) {
	b.messages[topic] = append(b.messages[topic], message)
	for _, consumer := range b.consumers[topic] {
		go b.sendMessage(topic, consumer)
	}
}

func (b *Broker) getNewMessages(consumer Consumer) {
	for topic, consumers := range b.consumers {
		if Contains(consumers, consumer) {
			b.sendMessage(topic, consumer)
		}
	}
}

func (b *Broker) sendMessage(topic string, consumer Consumer) {
	idx := b.lastConsumedMessage[topic][consumer]
	// fmt.Println("For consumer ", consumer, "last consumed message ", idx)
	for i := idx; i < len(b.messages[topic]); i++ {
		consumed := consumer.asyncConsumeMessages(topic, b.messages[topic][i])
		if consumed {
			b.lastConsumedMessage[topic][consumer] = i + 1
		}
	}
}
