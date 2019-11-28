package main

import (
	"bufio"
	"fmt"
	"os"

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

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter message: ")

	for {
		text, _ := reader.ReadString('\n')

		if text == "exit" {
			break
		}

		message := amqp.Publishing{
			Body: []byte(text),
		}

		err = channel.Publish("abc123", "random-key", false, false, message)
	}

	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	defer connection.Close()
}
