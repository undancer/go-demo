package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/undancer/go-demo/demo/protobuf"
	"log"
)

func main() {
	var p = tutorial.Person{
		Id:    1,
		Name:  "undancer",
		Email: "1@c.cn",
		Phones: []*tutorial.Person_PhoneNumber{
			{Number: "555", Type: tutorial.Person_MOBILE},
		},
	}

	fmt.Println(p)

	buff, err := proto.Marshal(&p);
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println(buff)
	fmt.Println(string(buff))
}
