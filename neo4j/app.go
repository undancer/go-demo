package neo4j

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"log"
)

const (
	url      = "bolt://localhost:7687"
	username = ""
	password = ""
)

func Main() {

	var (
		driver neo4j.Driver
		err    error
	)
	fmt.Println("neo4j")
	if driver, err = neo4j.NewDriver(url, neo4j.BasicAuth(username, password, "")); err != nil {
		log.Println(err.Error())
	}

	fmt.Println(driver)

}
