package main

import (
	"errors"
	"github.com/undancer/go-demo/mongodb"
	"github.com/undancer/go-demo/rabbitmq"
	"log"
)

func main() {
	switch "neo4j" {
	case "rabbitMQ":
		rabbitmq.Main()
		break
	case "MongoDB":
		mongodb.Main()
		break
	case "neo4j":
		//neo4j.Main()
		log.Println(errors.New("未实现").Error())
		break
	default:
		log.Println("未知内容")
		break
	}
}
