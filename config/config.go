package config

import (
	"github.com/trustmaster/goflow"
	"github.com/undancer/go-demo/flow"
)

func newGraphApp() *goflow.Graph {
	graph := goflow.NewGraph()

	graph.Add("consumer", new(flow.Consumer))
	graph.Add("greeter", new(flow.Greeter))
	graph.Add("printer", new(flow.Printer))
	//graph.Add("logger", new(flow.Logger))

	graph.Connect("consumer", "Out", "greeter", "Name")
	graph.Connect("greeter", "Res", "printer", "Line")
	//graph.Connect("consumer", "Out", "logger", "In")
	//graph.Connect("greeter", "Res", "logger", "In")

	//graph.MapInPort("In", "greeter", "Name")
	graph.MapInPort("In", "consumer", "In")
	return graph
}

func Config() {
	config := NewConfigRabbitMQ()
	ch := config.GetConsume("go-queue")
	app := newGraphApp()
	app.SetInPort("In", ch)
	wait := goflow.Run(app)
	<-wait
}
