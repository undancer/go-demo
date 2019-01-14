package main

import "fmt"

type name struct {
	username string
	age      int
}

func (*name) print() {
	fmt.Println("println !!")
}

func main() {
	name := name{
		username: "undancer",
		age:      18,
	}
	fmt.Println(name)
	fmt.Println(name.username)
	fmt.Println(name.age)
	name.print()
}
