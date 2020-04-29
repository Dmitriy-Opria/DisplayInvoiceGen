package consumer

import (
	"time"

	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"github.rakops.com/BNP/DisplayInvoiceGen/config"
)

const (
	exchange     = "test-exchange"   // "Durable, non-auto-deleted AMQP exchange name")
	exchangeType = "direct"          // "Exchange type - direct|fanout|topic|x-custom")
	queue        = "test-queue"      // "Ephemeral AMQP queue name")
	bindingKey   = "test-key"        // "AMQP binding key")
	consumerTag  = "simple-consumer" // "AMQP consumer tag (should not be blank)")
	lifetime     = 5 * time.Second   // "lifetime of process before shutdown (0s=infinite)")
)

type RabbitConsumer interface {
	Connect() error
	Shutdown() error
	GetChannel() (<-chan amqp.Delivery, error)
	Done(err error)
}

func NewConsumer(conf *config.Config) RabbitConsumer {
	rabbit := RabbitWrapper{}
	rabbit.conf = conf
	rabbit.consumerName = conf.Rabbit.ConsumerExchangeName
	rabbit.done = make(chan error)
	return &rabbit
}

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

func (r *RabbitWrapper) GetChannel() (<-chan amqp.Delivery, error) {
	ch, err := r.Connection.Channel()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to open a channel")
	}

	r.Channel = ch

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
		return nil, errors.Wrap(err, "Failed to setup exchange for consuming")
	}

	q, err := ch.QueueDeclare(
		r.conf.Rabbit.ConsumerQueueName, // name
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, errors.Wrap(err, "Failed to declare a queue for consuming")
	}

	err = ch.QueueBind(
		q.Name, // name
		r.conf.Rabbit.ConsumerRouteKey,
		r.conf.Rabbit.ConsumerExchangeName,
		false,
		nil,
	)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to bind a queue for consuming")
	}
	msgs, err := ch.Consume(
		r.conf.Rabbit.ConsumerQueueName, // queue
		r.consumerName,
		true,
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
