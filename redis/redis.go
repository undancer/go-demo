package redis

import (
	"fmt"
	"github.com/mediocregopher/radix/v3"
	"log"
)

func Main() {
	pool, err := radix.NewPool("tcp", "localhost:6379", 10)

	if err != nil {
		log.Println(err.Error())
	}
	err = pool.Do(radix.Cmd(nil, "SET", "name", "undancer"))
	fmt.Println(err)

	var name string
	err = pool.Do(radix.Cmd(&name, "GET", "name"))
	fmt.Println(name)

	fmt.Println(&pool)
	fmt.Println(err)

	pool.Close()

}
