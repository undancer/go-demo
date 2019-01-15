package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	var _map = make(map[string]interface{})
	_map["username"] = "undancer"
	_map["age"] = 18

	fmt.Println(_map)
	fmt.Println()

	var (
		bs  []byte
		err error
	)

	bs, err = json.Marshal(_map)
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println(bs)
	fmt.Println(string(bs))
	fmt.Println()

	bs, err = json.MarshalIndent(_map, "", "  ")
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println(bs)
	fmt.Println(string(bs))

	fmt.Println()
}
