package rabbitmq

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

func OpenChannel() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}

func Consume(ch *amqp.Channel, msgsOut chan<- amqp.Delivery, queue string) error {
	msgs, err := ch.Consume(
		queue,
		"go-consumer",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil
	}
	for msg := range msgs {
		msgsOut <- msg
	}
	return nil
}

func Publish(ch *amqp.Channel, body string, exName string) error {
	// err := ch.Publish(
	err := ch.PublishWithContext(
		context.Background(),
		exName,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	return err
}
