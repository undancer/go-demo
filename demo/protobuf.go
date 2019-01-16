package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/undancer/go-demo/demo/protobuf"
	"log"
)

func main() {
	ab := tutorial.AddressBook{
		People: []*tutorial.Person{
			{
				Id:    1,
				Name:  "undancer",
				Email: "1@c.cn",
				Phones: []*tutorial.Person_PhoneNumber{
					{Number: "555", Type: tutorial.Person_MOBILE},
				},
			},
		},
	}

	fmt.Println(ab)
	fmt.Println(len(ab.String()))

	buff, err := proto.Marshal(&ab);
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println(buff)
	fmt.Println(len(buff))
	fmt.Println(string(buff))
}
