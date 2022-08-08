package rabbitmq

import (
	"context"
	"fmt"
	"github.com/alecthomas/kingpin"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type RabbitModule struct {
	ServerURL string
	User      string
	Password  string
}

func (m *RabbitModule) Configure() {
	kingpin.Flag("rabbit-server", "Rabbit MQ server").
		Default("localhost:5672").
		Envar("RABBITMQ_SERVER_URL").
		StringVar(&m.ServerURL)
	kingpin.Flag("rabbit-user", "Rabbit MQ user").
		Default("guest").
		Envar("RABBITMQ_USER").
		StringVar(&m.User)
	kingpin.Flag("rabbit-password", "Rabbit MQ password").
		Default("guest").
		Envar("RABBITMQ_PASSWORD").
		StringVar(&m.Password)
	kingpin.Parse()
}

func (m *RabbitModule) ProvideQueueModule() RabbitQueue {
	m.Configure()

	return NewRabbitQueue(fmt.Sprintf("amqp://%s:%s@%s/",
		m.User, m.Password, m.ServerURL))
}

func NewRabbitQueue(uri string) RabbitQueue {
	rabbitMq := &rabbitQueue{URI: uri}
	rabbitMq.openChannel()

	return rabbitMq
}

type RabbitQueue interface {
	CreateQueue(name string) amqp.Queue
	Publish(ctx context.Context, queue amqp.Queue, message string)
	PollMessages(queue amqp.Queue) <-chan amqp.Delivery
}

type rabbitQueue struct {
	URI     string
	Channel *amqp.Channel
}

func (r rabbitQueue) connect() *amqp.Connection {
	conn, err := amqp.Dial(r.URI)
	if err != nil {
		log.Panicf("%s: %s", "Failed to connect to RabbitMQ", err)
	}

	return conn
}

func (r *rabbitQueue) openChannel() {
	conn := r.connect()

	ch, err := conn.Channel()
	if err != nil {
		log.Panicf("%s: %s", "Failed to open a channel", err)
	}

	r.Channel = ch
}

func (r *rabbitQueue) CreateQueue(name string) amqp.Queue {
	q, err := r.Channel.QueueDeclare(
		name,  // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Panicf("%s: %s", "Cannot declare rabbit mq Queue", err)
	}

	return q
}

func (r *rabbitQueue) Publish(ctx context.Context, queue amqp.Queue, message string) {
	err := r.Channel.PublishWithContext(
		ctx,
		"",         // exchange
		queue.Name, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		log.Panicf("%s: %s", "Failed to publish a message", err)
	}

	log.Printf("Message published %s\n", message)
}

func (r *rabbitQueue) PollMessages(queue amqp.Queue) <-chan amqp.Delivery {
	messages, err := r.Channel.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		log.Panicf("%s: %s", "Failed to poll messages", err)
	}

	return messages
}
