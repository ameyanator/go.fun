package main

type Broker struct {
	topics []*Topic
}

func NewBroker() *Broker {
	return &Broker{
		topics : make([]*Topic, 0),
	}
}

func (b *Broker) RegisterConsumerWithTopic()