package producer

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"github.rakops.com/BNP/DisplayInvoiceGen/config"
	"github.rakops.com/BNP/DisplayInvoiceGen/log"
)

type RabbitProducer interface {
	Connect() error
	Close()
	Send(body string) error
	SendJSON(request interface{}) error
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

	err = ch.ExchangeDeclare(
		r.conf.Rabbit.ConsumerExchangeName, // name
		"topic",                            // type
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return errors.Wrap(err, "Failed to setup exchange for publishing")
	}
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

	err = ch.QueueBind(
		q.Name, // name
		r.conf.Rabbit.ProducerRouteKey,
		r.conf.Rabbit.ConsumerExchangeName,
		false,
		nil,
	)
	if err != nil {
		return errors.Wrap(err, "Failed to bind a queue for publishing")
	}

	err = ch.Publish(
		r.conf.Rabbit.ConsumerExchangeName,
		r.conf.Rabbit.ProducerRouteKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	return errors.Wrap(err, "Failed to publish msg")
}

func (r *RabbitWrapper) SendJSON(request interface{}) error {
	ch, err := r.Connection.Channel()
	if err != nil {
		return errors.Wrap(err, "Failed to open a channel")
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		r.conf.Rabbit.ConsumerExchangeName, // name
		"topic",                            // type
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return errors.Wrap(err, "Failed to setup exchange for publishing")
	}
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

	err = ch.QueueBind(
		q.Name, // name
		r.conf.Rabbit.ProducerRouteKey,
		r.conf.Rabbit.ConsumerExchangeName,
		false,
		nil,
	)
	if err != nil {
		return errors.Wrap(err, "Failed to bind a queue for publishing")
	}

	body, err := json.Marshal(request)
	if err != nil {
		return errors.Wrap(err, "Failed to marshal message for publishing")
	}

	err = ch.Publish(
		r.conf.Rabbit.ConsumerExchangeName,
		r.conf.Rabbit.ProducerRouteKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})

	return errors.Wrap(err, "Failed to publish msg")
}
