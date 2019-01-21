package config

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type Consumer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	tag     string
	done    chan error
}

type RabbitMQConfig struct {
	Consumer

	c *Consumer
}

func NewConsumer(amqpURI, exchange, exchangeType, queueName, key, ctag string) (*Consumer, error) {
	c := &Consumer{
		conn:    nil,
		channel: nil,
		tag:     ctag,
		done:    make(chan error),
	}

	var err error

	log.Printf("dialing %q", amqpURI)
	c.conn, err = amqp.Dial(amqpURI)
	if err != nil {
		return nil, fmt.Errorf("err: Dial: %s", err)
	}

	go func() {
		fmt.Printf("closing %s", <-c.conn.NotifyClose(make(chan *amqp.Error)))
	}()

	log.Printf("got Connection, getting Channel")

	c.channel, err = c.conn.Channel()

	if err != nil {
		return nil, fmt.Errorf("err: Channel: %s", err)
	}

	log.Printf("got Channel, declaring Exchange (%q)", exchange)
	if err = c.channel.ExchangeDeclare(
		exchange,     // name of the exchange
		exchangeType, // type
		true,         // durable
		false,        // delete when complete
		false,        // internal
		false,        // noWait
		nil,          // arguments
	); err != nil {
		return nil, fmt.Errorf("err: Exchange Declare: %s", err)
	}

	log.Printf("declared Exchange, declaring Queue %q", queueName)
	queue, err := c.channel.QueueDeclare(
		queueName, // name of the queue
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // noWait
		nil,       // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("err: Queue Declare: %s", err)
	}

	log.Printf("declared Queue (%q %d messages, %d consumers), binding to Exchange (key %q)",
		queue.Name, queue.Messages, queue.Consumers, key)

	if err = c.channel.QueueBind(
		queue.Name, // name of the queue
		key,        // bindingKey
		exchange,   // sourceExchange
		false,      // noWait
		nil,        // arguments
	); err != nil {
		return nil, fmt.Errorf("err: Queue Bind: %s", err)
	}

	c.channel.Qos(1, 0, false)

	//log.Printf("Queue bound to Exchange, starting Consume (consumer tag %q)", c.tag)
	//deliveries, err := c.channel.Consume(
	//	queue.Name, // name
	//	c.tag,      // consumerTag,
	//	false,      // noAck
	//	false,      // exclusive
	//	false,      // noLocal
	//	false,      // noWait
	//	nil,        // arguments
	//)
	//if err != nil {
	//	return nil, fmt.Errorf("err: Queue Consume: %s", err)
	//}
	//
	//go handle(deliveries, c.done)

	return c, nil
}

func NewConfigRabbitMQ() *RabbitMQConfig {

	var c, err = NewConsumer(
		"amqp://go:go@localhost//go-vhost",
		"go-exchange",
		amqp.ExchangeDirect,
		"go-queue",
		"go",
		"go")

	if err != nil {
		log.Println(err.Error())
	}

	return &RabbitMQConfig{
		c: c,
	}

}

func handle(deliveries <-chan amqp.Delivery, done chan error) {
	for d := range deliveries {
		log.Printf(
			"got %dB delivery: [%v] %q",
			len(d.Body),
			d.DeliveryTag,
			d.Body,
		)
		d.Ack(false)
	}
	log.Printf("handle: deliveries channel closed")
	done <- nil
}

func (c *RabbitMQConfig) GetConsume(queueName string) <-chan amqp.Delivery {

	var deliveries <-chan amqp.Delivery
	var err error

	deliveries, err = c.c.channel.Consume(queueName,
		"",
		true,
		false,
		false,
		false,
		nil)

	if err != nil {
		log.Println(err.Error())
	}

	return deliveries
}

func (c *Consumer) Shutdown() error {
	// will close() the deliveries channel
	if err := c.channel.Cancel(c.tag, true); err != nil {
		return fmt.Errorf("err: Consumer cancel failed: %s", err)
	}

	if err := c.conn.Close(); err != nil {
		return fmt.Errorf("AMQP connection close error: %s", err)
	}

	defer log.Printf("AMQP shutdown OK")

	// wait for handle() to exit
	return <-c.done
}
