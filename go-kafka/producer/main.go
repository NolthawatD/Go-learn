package main

import (
	"fmt"

	"gopkg.in/Shopify/sarama.v1"
)

func main() {

	servers := []string{"localhost:9092"}

	producer, err := sarama.NewSyncProducer(servers, nil)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	msg := sarama.ProducerMessage{
		Topic: "nolhello",
		Value: sarama.StringEncoder("Hello World"),
	}

	p, o, err := producer.SendMessage(&msg)
	if err != nil {
		panic(err)
	}
	fmt.Println("partition=%v, offset=%v", p, o)
}
