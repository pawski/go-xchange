package rabbitmq

import (
	"sync"

	"github.com/streadway/amqp"

	"github.com/pawski/go-xchange/configuration"
	"github.com/pawski/go-xchange/logger"
)

var connection = amqp.Connection{}
var once sync.Once

func Connect() amqp.Connection {
	once.Do(func() {
		conn, err := amqp.Dial(configuration.Get().RabbitMqUrl)
		failOnError(err, "Failed to connect to RabbitMQ")

		connection = *conn
	})

	return connection
}

func ConsumeFromQueue(consume func(<-chan amqp.Delivery)) {
	Connect()

	channel, err := connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer channel.Close()

	q, err := channel.QueueDeclare(
		"rates_queue", // name
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = channel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	messages, err := channel.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	consume(messages)
}

func PublishToQueue(message []byte) {

	Connect()

	channel, err := connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer channel.Close()

	q, err := channel.QueueDeclare(
		"rates_queue", // name
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := message
	err = channel.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	logger.Get().Infof(" [x] Sent %s", message)
	failOnError(err, "Failed to publish a message")
}

func failOnError(err error, msg string) {
	if err != nil {
		logger.Get().Fatalf("%s: %s", msg, err)
	}
}
