package config

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
)

const (
	amqpUrl      = "amqp://go:go@localhost//go-vhost"
	exchangeName = "go-exchange"
	routingKey   = "go"
	//queueName    = "go-queue"
)

var (
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
	err     error
)

type RabbitMQConfig struct {
}

func NewConfigRabbitMQ() RabbitMQConfig {

	fmt.Println("RabbitMQ DEMO")

	if conn, err = amqp.Dial(amqpUrl); err != nil { // 连接RabbitMQ
		log.Println("连接失败")
		os.Exit(100)
	}

	//defer func() {
	//	fmt.Println("conn close")
	//	conn.Close()
	//}()

	if channel, err = conn.Channel(); err != nil { //创建频道
		log.Println("频道创建失败")
		os.Exit(101)
	}

	//defer func() {
	//	fmt.Println("chan close")
	//	channel.Close()
	//}()

	err = channel.ExchangeDeclare(
		exchangeName,
		amqp.ExchangeDirect,
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		log.Println(err.Error())
	}

	var queueName = "go-queue"

	queue, err = channel.QueueDeclare(
		queueName, //设置为空，让程序自动生成
		true,
		true,
		false,
		false,
		nil)
	if err != nil {
		log.Println(err.Error())
	}

	queueName = queue.Name

	err = channel.QueueBind(
		queueName,
		routingKey,
		exchangeName,
		false,
		nil)
	if err != nil {
		log.Println(err.Error())
	}

	err = channel.Qos(1, 0, false)
	if err != nil {
		log.Println(err.Error())
	}
	//
	//var deliveries <-chan amqp.Delivery
	//
	//deliveries, err = channel.Consume(queueName,
	//	"",
	//	true,
	//	false,
	//	false,
	//	false,
	//	nil)
	//
	//fmt.Println("MQ", &channel)
	//
	//sessions := make(chan bool)
	//
	//go func() {
	//	fmt.Println("go init")
	//	for delivery := range deliveries {
	//		var str = utils.BytesToString(&(delivery.Body))
	//		fmt.Printf("收到消息 %s \n", *str)
	//	}
	//	fmt.Println("gone")
	//}()
	//
	//<-sessions

	//var deliveries, _ = channel.Consume(
	//	queueName,
	//	"",
	//	false,
	//	false,
	//	false,
	//	false,
	//	nil)
	//
	//sessions := make(chan bool)
	//
	//go func() {
	//	for delivery := range deliveries {
	//		var str = utils.BytesToString(&(delivery.Body))
	//		fmt.Printf("收到消息 %s \n", *str)
	//
	//	}
	//}()
	//
	//<-sessions

	return RabbitMQConfig{}

}

func (c *RabbitMQConfig) GetConsume(queueName string) <-chan amqp.Delivery {

	var deliveries <-chan amqp.Delivery

	deliveries, err = channel.Consume(queueName,
		"",
		true,
		false,
		false,
		false,
		nil)

	return deliveries
}

func (c *RabbitMQConfig) Close() {
	conn.Close()
	channel.Close()
}
