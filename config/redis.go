package config

import (
	"fmt"
	"github.com/mediocregopher/radix/v3"
	"log"
)

const (
	redisUrl = "localhost:6379"
)

var (
	RedisPool *radix.Pool
)

func init() {
	pool, err := configRedis()
	if err != nil {
		log.Println(err)
	}
	RedisPool = pool
}

func configRedis() (*radix.Pool, error) {
	pool, err := radix.NewPool("tcp", redisUrl, 10)
	if err != nil {
		fmt.Println(err.Error())
		return nil, fmt.Errorf("err: Redis %s", err)
	}

	log.Println("redis 初始化完成")

	return pool, nil

}
