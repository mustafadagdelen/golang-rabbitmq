package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	var url string = "amqp://url"

	connection, err := amqp.Dial(url)

	if err != nil {
		panic("could not establish connection with RabbitMQ:" + err.Error())
	}

	channel, err := connection.Channel()

	if err != nil {
		panic("could not open RabbitMQ channel:" + err.Error())
	}

	err = channel.ExchangeDeclare("abc123", "topic", true, false, false, false, nil)

	if err != nil {
		panic(err)
	}

	_, err = channel.QueueDeclare("abc12queue", true, false, false, false, nil)

	if err != nil {
		panic("error declaring the queue: " + err.Error())
	}

	err = channel.QueueBind("abc12queue", "#", "abc123", false, nil)

	if err != nil {
		panic("error binding to the queue: " + err.Error())
	}

	msgs, err := channel.Consume("abc12queue", "", false, false, false, false, nil)

	if err != nil {
		panic("error consuming the queue: " + err.Error())
	}

	for msg := range msgs {
		fmt.Println("message received: " + string(msg.Body))
		msg.Ack(false)
	}

	defer connection.Close()
}
