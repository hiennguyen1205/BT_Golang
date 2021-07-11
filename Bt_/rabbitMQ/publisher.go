package rabbitMQ

import (
	"rabbit/connect"
	"github.com/streadway/amqp"
)

type producer struct{

}

func Publisher(ch *amqp.Channel, q amqp.Queue, body string) (err error){
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	connect.FailOnError(err, "Failed to publish a message")
	return err
}