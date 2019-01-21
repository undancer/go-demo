package flow

import (
	"github.com/streadway/amqp"
	"github.com/undancer/go-demo/utils"
)

type Consumer struct {
	In  <-chan amqp.Delivery // input port
	Out chan<- string        // output port
}

func (c *Consumer) Process() {

	for delivery := range c.In {
		var str = utils.BytesToString(&(delivery.Body))
		c.Out <- *str
	}

}
