package producer

import (
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"github.rakops.com/BNP/DisplayInvoiceGen/config"
	"github.rakops.com/BNP/DisplayInvoiceGen/log"
)

type RabbitProducer interface {
	Connect() error
	Close()
	Send(body string) error
}

func NewProducer(conf *config.Config) RabbitProducer {
	rabbit := RabbitWrapper{}
	rabbit.conf = conf

	return &rabbit
}

type RabbitWrapper struct {
	Connection *amqp.Connection
	conf       *config.Config
}

func (r *RabbitWrapper) Connect() error {
	conf := amqp.URI{
		Scheme:   "amqp",
		Host:     r.conf.Rabbit.Host,
		Port:     r.conf.Rabbit.Port,
		Username: r.conf.Rabbit.User,
		Password: r.conf.Rabbit.Pass,
		Vhost:    r.conf.Rabbit.VHost,
	}.String()

	log.Infof("amqp: %v", conf)

	conn, err := amqp.Dial(conf)
	r.Connection = conn

	return errors.Wrap(err, "can't connect to rabbitMQ")
}

func (r *RabbitWrapper) Close() {
	r.Connection.Close()
}

func (r *RabbitWrapper) Send(body string) error {
	ch, err := r.Connection.Channel()
	if err != nil {
		return errors.Wrap(err, "Failed to open a channel")
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		r.conf.Rabbit.ProducerQueueName, // name
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return errors.Wrap(err, "Failed to get a queue for publishing")
	}

	err = ch.Publish(
		"",
		q.Name, // routing key
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	return errors.Wrap(err, "Failed to publish msg")
}
