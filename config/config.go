package config

func Config() {
	config := NewConfigRabbitMQ()
	ch := config.GetConsume("go-queue")
	app := newFlowApp()
	app.Run(ch)
}
