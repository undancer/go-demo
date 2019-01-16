package main

import (
	"fmt"
	"github.com/caarlos0/env"
	"log"
	"time"
)

type config struct {
	Home         string        `env:"HOME"`
	Port         int           `env:"PORT" envDefault:"3000"`
	IsProduction bool          `env:"PRODUCTION"`
	Hosts        []string      `env:"HOSTS" envSeparator:":"`
	Duration     time.Duration `env:"DURATION"`
	TempFolder   string        `env:"TEMP_FOLDER" envDefault:"${HOME}/tmp" envExpand:"true"`
}

func main() {

	var config = config{}

	if err := env.Parse(&config); err != nil {
		log.Println(err.Error())
	}

	fmt.Println(config)
	fmt.Printf("%+v\n", config)

}
