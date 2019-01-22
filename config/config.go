package config

import "fmt"

func Config() {
	fmt.Println(RedisPool)
	config := NewConfigRabbitMQ()
	ch := config.GetConsume("go-queue")
	app := newFlowApp()
	app.Run(ch)
}
