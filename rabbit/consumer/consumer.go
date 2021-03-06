package consumer

import (
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"github.rakops.com/BNP/DisplayInvoiceGen/config"
)

type RabbitConsumer interface {
	Connect() error
	Shutdown() error
	GetChannel(to Source) (<-chan amqp.Delivery, error)
	Done(err error)
}

func NewConsumer(conf *config.Config) RabbitConsumer {
	rabbit := RabbitWrapper{}
	rabbit.conf = conf
	rabbit.consumerName = conf.Rabbit.ExchangeName
	rabbit.done = make(chan error)
	return &rabbit
}

type Source int64

const (
	Invoice    = 1
	SalesForce = 2
)

type RabbitWrapper struct {
	Connection   *amqp.Connection
	Channel      *amqp.Channel
	conf         *config.Config
	consumerName string
	done         chan error
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

	conn, err := amqp.Dial(conf)
	r.Connection = conn

	return errors.Wrap(err, "can't connect to rabbitMQ")
}

func (r *RabbitWrapper) GetChannel(from Source) (<-chan amqp.Delivery, error) {
	ch, err := r.Connection.Channel()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to open a channel")
	}

	r.Channel = ch

	queueName := ""

	switch from {
	case Invoice:
		queueName = r.conf.Rabbit.ConsumerInvoiceQueueName
	case SalesForce:
		queueName = r.conf.Rabbit.ConsumerSFQueueName
	}

	_, err = ch.QueueDeclare(
		queueName, // name
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, errors.Wrap(err, "Failed to declare a queue for consuming")
	}

	msgs, err := ch.Consume(
		queueName, // queue
		r.consumerName,
		false,
		false,
		false,
		false,
		nil,
	)

	return msgs, errors.Wrap(err, "Failed to get consumer channel")
}

func (r *RabbitWrapper) Shutdown() error {
	// will close() the deliveries channel
	if err := r.Channel.Cancel(r.consumerName, true); err != nil {
		return errors.Wrap(err, "Consumer cancel failed")
	}

	if err := r.Connection.Close(); err != nil {
		return errors.Wrap(err, "AMQP connection close error: %s")
	}

	// wait for handle() to exit
	return <-r.done
}

func (r *RabbitWrapper) Done(err error) {
	r.done <- err
}
