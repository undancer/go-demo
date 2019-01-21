package flow

import "fmt"

type Printer struct {
	Line <-chan string
}

func (p *Printer) Process() {

	for n := range p.Line {
		fmt.Println("printer", n)
	}

}
