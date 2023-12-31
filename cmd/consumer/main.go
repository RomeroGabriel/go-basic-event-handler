package main

import (
	"fmt"

	"github.com/RomeroGabriel/go-basic-event-handler/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch := rabbitmq.OpenChannel()
	defer ch.Close()

	msgs := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, msgs, "my_queue")
	for msg := range msgs {
		fmt.Println(string(msg.Body))
		msg.Ack(false)
	}
}
