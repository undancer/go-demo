package flow

import "fmt"

type Greeter struct {
	Name <-chan string // input port
	Res  chan<- string // output port
}

func (g *Greeter) Process() {

	for n := range g.Name {
		fmt.Println("greeter", n)
		g.Res <- n
	}

}
