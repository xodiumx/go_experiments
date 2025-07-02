package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

const rabbitDSN = "amqp://admin:admin@localhost:5672/"

type Message struct {
	Event string `json:"event"`
	ID    int    `json:"id"`
}

func failOnError(err error, msg string) {
	// Error callback
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

	i := 0

	for {

		// gen message
		msg := Message{
			Event: "user_created",
		}
		msg.ID = i
		body, err := json.Marshal(msg)
		failOnError(err, "Failed to marshal JSON")

		// Send to queue
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key (queue name)
			false,  // mandatory - false bool mandatory If true and the message cannot be delivered to the queue - it is returned to the sender (in Go via Return callback).
			false,  // immediate - Not used in RabbitMQ (deprecated parameter). In other implementations: if there is no active consumer, the message is not sent.
			amqp091.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		failOnError(err, "Failed to publish a message")
		log.Printf(" [x] Sent %s", body)
		time.Sleep(1 * time.Second)

		i++
	}
}
