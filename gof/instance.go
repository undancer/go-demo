package main

import (
	"fmt"
	i "github.com/undancer/go-demo/gof/instance"
)

func main() {
	a := i.GetInstance()
	a.Name = "a"
	b := i.GetInstance()
	b.Name = "b"
	fmt.Println(&a.Name, a)
	fmt.Println(&b.Name, b)
	fmt.Printf("%p %T\n", a, a)
	fmt.Printf("%p %T\n", b, b)
}
