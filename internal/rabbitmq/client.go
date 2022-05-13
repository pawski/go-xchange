package rabbitmq

import (
	"sync"

	"github.com/streadway/amqp"

	"github.com/pawski/go-xchange/internal/configuration"
	"github.com/pawski/go-xchange/internal/logger"
)

var connection = &amqp.Connection{}
var once sync.Once

func Connect() *amqp.Connection {
	once.Do(func() {
		var err error
		connection, err = amqp.Dial(configuration.Get().RabbitMqUrl)
		failOnError(err, "Failed to connect to RabbitMQ")
	})

	return connection
}

func NewConnection(url string) (*amqp.Connection, error) {
	return amqp.Dial(url)
}

func ConsumeFromQueue(channel *amqp.Channel, queueName string, consume func(<-chan amqp.Delivery)) {
	err := channel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	messages, err := channel.Consume(
		queueName, // queue
		"",        // consumer
		false,     // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	failOnError(err, "Failed to register a consumer")

	consume(messages)
}

func Publish(channel *amqp.Channel, queueName string, message []byte) error {
	return channel.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		})
}

func failOnError(err error, msg string) {
	if err != nil {
		logger.Get().Fatalf("%s: %s", msg, err)
	}
}
