package main

type ProducerFactory struct {
	producer map[string]Producer
}

func NewProducerFactory() *ProducerFactory {
	producer := make(map[string]Producer)
	producer["Amazon"] = &Amazon{}
	producer["Flipkart"] = &Flipkart{}

	return &ProducerFactory{
		producer: producer,
	}
}

func (p *ProducerFactory) getProducer(name string) Producer {
	return p.producer[name]
}
