package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"github.com/undancer/go-demo/utils"
	"log"
	"os"
)

const (
	url          = "amqp://localhost"
	username     = "spark"
	password     = "spark"
	vhost        = "/go-vhost"
	exchangeName = "go-exchange"
	queueName    = "go-queue"
)

func Main() {
	fmt.Println("RabbitMQ DEMO")
	config := amqp.Config{
		SASL: []amqp.Authentication{
			&amqp.PlainAuth{
				Username: username,
				Password: password,
			},
		},
		Vhost: vhost,
	}

	var (
		conn    *amqp.Connection
		channel *amqp.Channel
		err     error
	)

	if conn, err = amqp.DialConfig(url, config); err != nil { // 连接RabbitMQ
		log.Println("连接失败")
		os.Exit(100)
	}

	defer conn.Close()

	if channel, err = conn.Channel(); err != nil { //创建频道
		log.Println("频道创建失败")
		os.Exit(101)
	}

	defer channel.Close()

	if channel.ExchangeDeclare(
		exchangeName,
		amqp.ExchangeDirect,
		true,
		false,
		false,
		false,
		nil); err != nil {
		log.Println(err.Error())
	}

	if _, err := channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil); err != nil {
		log.Println(err.Error())
	}

	if err := channel.QueueBind(
		queueName,
		"go-test",
		exchangeName,
		false,
		nil); err != nil {
		log.Println(err.Error())
	}

	var deliveries, _ = channel.Consume(
		queueName,
		"",
		false,
		false,
		false,
		false,
		nil)

	sessions := make(chan bool)

	go func() {
		for delivery := range deliveries {
			var str = utils.BytesToString(&(delivery.Body))
			fmt.Printf("收到消息 %s \n", *str)

		}
	}()

	<-sessions

}
