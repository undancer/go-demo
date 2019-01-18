package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	//context.Background()
	//
	//context2.Background()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println(ctx.Err())

	<-ctx.Done()

	fmt.Println(ctx.Err())

	//context.WithTimeout()

	//in := gen(2, 3)

	c1 := sq(gen(2, 3))
	c2 := sq(gen(4, 5))

	for n := range merge(c1, c2) {
		fmt.Println(n)
	}

}

func gen(nums ...int) <-chan int {
	out := make(chan int, len(nums))
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int, 1)

	// 为每一个输入channel cs 创建一个 goroutine output
	// output 将数据从 c 拷贝到 out，直到 c 关闭，然后 调用 wg.Done
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// 启动一个 goroutine，用于所有 output goroutine结束时，关闭 out
	// 该goroutine 必须在 wg.Add 之后启动
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
