package main

import (
	"bufio"
	"log"

	"rabbit/connect"
	"rabbit/rabbitMQ"
	"rabbit/readFile"
)

func main() {
	//create rmq connection
	rmq := connect.RMQ{
		ConnectionString: "amqp://guest:guest@174.138.40.239:5672/",
	}

	//tạo conection
	rmq.CreateConnect()
	defer rmq.Close()

	//tạo channel
	ch := rmq.GetChannel()

	// err := ch.ExchangeDeclare(
	// 	"hien",   // name
	// 	"direct", // type
	// 	true,     // durable
	// 	false,    // auto-deleted
	// 	false,    // internal
	// 	false,    // no-wait
	// 	nil,      // arguments
	// )
	// connect.FailOnError(err, "Failed to declare an exchange")

	//tạo queue trên rabbitMQ nếu chưa có (có rồi thì dùng luôn, không cần tạo)
	q, err := ch.QueueDeclare(
		"hien1", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	connect.FailOnError(err, "Failed to declare a queue")

	//bind 
	// err = ch.QueueBind(
	// 	q.Name,   // queue name
	// 	"btapFile",       // routing key
	// 	"hien", // exchange
	// 	false,
	// 	nil,
	// )
	// connect.FailOnError(err, "Failed to bind a queue")
	//mở file
	file := readFile.ReadFile()
	defer file.Close()
	//đọc file
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	data := make(map[string]int)
	for scanner.Scan() {
		body := scanner.Text()
		data[string(body)]++
		//publish message
		err = rabbitMQ.Publisher(ch, q, body)
		if err != nil {
			log.Println("Error in publisher")
		}
	}

forever := make(chan bool)

	var dataLine []string
	//consume message
	msgs := rabbitMQ.Consume(ch, q)
	go func() {
		i := 0
		for d := range msgs {
			i++
			log.Printf("Received a message:%d %s", i, d.Body)
			dataLine = append(dataLine, string(d.Body))
		}
		log.Println(1)
	}()
	<-forever


}
