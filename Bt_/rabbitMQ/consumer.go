package rabbitMQ

import (
	"rabbit/connect"
	"github.com/streadway/amqp"
)

func Consume(ch *amqp.Channel, q amqp.Queue) <-chan amqp.Delivery {
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	connect.FailOnError(err, "Failed to register a consumer")
	return msgs
}
