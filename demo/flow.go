package main

import (
	"fmt"
	"github.com/trustmaster/goflow"
)

// Greeter sends greetings
type Greeter struct {
	Name <-chan string // input port
	Res  chan<- string // output port
}

// Process incoming data
func (c *Greeter) Process() {
	// Keep reading incoming packets
	for name := range c.Name {
		greeting := fmt.Sprintf("Hello, %s!", name)
		// Send the greeting to the output port
		c.Res <- greeting
	}
}

// Printer prints its input on screen
type Printer struct {
	Line <-chan string // inport
}

// Process prints a line when it gets it
func (c *Printer) Process() {
	for line := range c.Line {
		fmt.Println(line)
	}
}

// NewGreetingApp defines the app graph
func NewGreetingApp() *goflow.Graph {
	n := goflow.NewGraph()

	// Add processes to the network
	n.Add("greeter", new(Greeter))

	n.Add("printer", new(Printer))
	// Connect them with a channel
	n.Connect("greeter", "Res", "printer", "Line")
	// Our net has 1 inport mapped to greeter.Name
	n.MapInPort("In", "greeter", "Name")
	return n
}

func main() {
	// Create the network
	net := NewGreetingApp()
	// We need a channel to talk to it
	in := make(chan string)
	net.SetInPort("In", in)
	// Run the net
	wait := goflow.Run(net)
	// Now we can send some names and see what happens
	in <- "John"
	in <- "Boris"
	in <- "Hanna"
	// Send end of input
	close(in)
	// Wait until the net has completed its job
	<-wait
}
