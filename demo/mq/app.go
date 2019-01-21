package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type MQ struct {
	ch <-chan interface{}
}

func (q *MQ) config() {

}

func main() {
	var wg sync.WaitGroup

	ch := make(chan string)
	ch2 := make(chan string)

	defer func() {
		fmt.Println("close")
	}()

	go func() {
		for n := 0; n < 10; n++ {
			wg.Add(1)
			ch <- strconv.FormatInt(int64(n), 10)
		}
	}()

	go func() {
		for {
			select {
			case i := <-ch2:
				fmt.Println("ch2", i)
			}
		}
	}()

	go func() {
		for n := range ch {
			//获取消息
			func(i string) {
				//defer func() {
				//	wg.Done()
				//	fmt.Println("d", i)
				//}()

				//done := make(chan string)

				//异步处理
				go func(i string) {
					time.Sleep(time.Second)
					//done <- i
					fmt.Println("c", i)
					ch2 <- i
				}(i)

				//消费
				//fmt.Println("c", <-done)
			}(n)
		}
	}()

	time.Sleep(500 * time.Microsecond)

	wg.Wait()
	fmt.Println("wait")
}
