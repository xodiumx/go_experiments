package main

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

const rabbitDSN = "amqp://admin:admin@localhost:5672/"

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {

	// Connect to rabbit
	conn, err := amqp091.Dial(rabbitDSN)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer func(conn *amqp091.Connection) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Failed to close RabbitMQ connection")
		}
	}(conn)

	// Open rabbit chanel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer func(ch *amqp091.Channel) {
		err := ch.Close()
		if err != nil {
			log.Fatalf("Failed to close RabbitMQ channel")
		}
	}(ch)

	// Create queue
	q, err := ch.QueueDeclare(
		"hello", // queue name
		false,   // durable - save after broker close
		false,   // delete when unused
		false,   // exclusive - queue only for current connection
		false,   // no-wait - no wait response from server
		nil,     // arguments - TTL, length, etc.
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer - id
		true,   // auto-ack - no wait confirm of deliv msg
		false,  // exclusive - only one consumer can read queue
		false,  // no-local
		false,  // no-wait - no wait response from server
		nil,    // args - x-priority, etc.
	)
	failOnError(err, "Failed to register a consumer")

	log.Println(" [*] Waiting for messages. To exit press CTRL+C")
	for msg := range msgs {
		log.Printf(" [x] Received: %s", msg.Body)
	}
}
