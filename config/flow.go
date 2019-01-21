package config

import (
	"github.com/trustmaster/goflow"
	"github.com/undancer/go-demo/flow"
)

type FlowApp struct {
	app *goflow.Graph
}

func newFlowApp() *FlowApp {
	app := &FlowApp{}

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

	app.app = graph

	return app
}

func (a *FlowApp) Run(ch interface{}) {
	a.app.SetInPort("In", ch)
	wait := goflow.Run(a.app)
	<-wait
}
