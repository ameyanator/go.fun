package main

type ConsumerFactory struct {
	genZ       Consumer
	millennial Consumer
}

func NewConsumerFactory() *ConsumerFactory {
	return &ConsumerFactory{
		genZ:       &GenZ{},
		millennial: &Millennial{},
	}
}

func (c *ConsumerFactory) getConsumer(birthYear int) Consumer {
	if birthYear < 2000 {
		return c.millennial
	} else {
		return c.genZ
	}
}
