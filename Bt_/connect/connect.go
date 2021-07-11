package connect

import (
	"log"

	"github.com/streadway/amqp"
)

type RMQ struct {
	ConnectionString string
	conn       *amqp.Connection
}
// /("amqp://guest:guest@174.138.40.239:5672/"

//tạo connect
func (rmq *RMQ) CreateConnect() {
	conn, err := amqp.Dial(rmq.ConnectionString)
	FailOnError(err, "Failed to connect to RabbitMQ")
	rmq.conn = conn
}

//tạo channel
func  (rmq *RMQ)GetChannel() *amqp.Channel {
	var channel, error = rmq.conn.Channel()
	FailOnError(error, "Failed to create channel")
	return channel
}

//đóng connect
func (rmq *RMQ) Close() {
	rmq.conn.Close()
}

//check lỗi
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
