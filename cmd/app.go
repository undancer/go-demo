package main

import (
	"errors"
	"github.com/undancer/go-demo/mongodb"
	"github.com/undancer/go-demo/mysql"
	"github.com/undancer/go-demo/rabbitmq"
	"log"
)

func main() {
	switch "sign" {
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
	case "mysql":
		mysql.Main()
		break
	default:
		log.Println("未知内容")
		break
	}
}
