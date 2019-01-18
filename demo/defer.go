package main

import (
	"fmt"
)

type MQ struct {
}

func (mq *MQ) gone(i int) {
	fmt.Println("g", i)
}

func (mq *MQ) Close() {
	fmt.Println("close")
}

func main() {
	ch := make(chan int)
	sessions := make(chan struct {})
	//mq := new(MQ)
	//defer mq.Close()

	c := func() {
		fmt.Println("close")
	}

	defer c()

	g := func(i int) {
		fmt.Println("g", i)
	}

	f := func(i int) {
		//defer mq.gone(i)
		defer g(i)
		fmt.Println("c", i)
	}

	go func() {
		for i := 0; i < 10; i ++ {
			r := i //int(rand.Int63n(time.Now().UnixNano()))
			fmt.Println("p", r)
			ch <- r
		}
		close(sessions)
	}()

	go func() {

		for i := range ch {
			f(i)
			//fmt.Println("c", i)
		}
	}()


	for {
		select {
		case i:= <-ch:
			//println(reflect.TypeOf(ok))
			f(i)
		case <-sessions:
			fmt.Println("done")
			return
		}
	}
}
