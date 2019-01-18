package main

import (
	"errors"
	"fmt"
)

func doSomeThing(data interface{}) error {
	fmt.Println(data)

	return errors.New("err")
}

func main() {
	closechan := make(chan int, 0)
	dchan := make(chan int, 2)

	go func() {
		for {
			select {
			case data := <-dchan:        //2
				err := doSomeThing(data) //3
				if err != nil /* some thing wrong*/ { //4

					select {
					case <-closechan:
						return
					default:
						close(closechan) //5
					}
				}
			}
		}
	}()

	for i := 0; i < 10; i ++ {

		select {
		case <-closechan:
			fmt.Println("channel already closed.")
			return
		default:
			fmt.Println("channel not closed, do your things")
			dchan <- 1 //1
		}
	}

}
