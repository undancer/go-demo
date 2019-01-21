package flow

import "log"

type Logger struct {
	In <-chan string
}

func (l *Logger) Process() {

	for msg := range l.In {
		log.Println("logger", msg)
	}

}
