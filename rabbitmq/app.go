package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
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

	if channel, err = conn.Channel(); err != nil { //创建频道
		log.Println("频道创建失败")
		os.Exit(101)
	}

	fmt.Println(channel)

}
