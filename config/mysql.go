package config

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/undancer/go-demo/db/mysql"
)

const mysqlUrl = "pio:pio@tcp(localhost)/pio"

func init() {
	fmt.Println("初始化mysql")
	configMysql()
}

func configMysql() {

	config := make(map[string]string)
	config["host"] = "localhost"
	config["port"] = "3306"
	config["username"] = "pio"
	config["password"] = "pio"
	config["database"] = "pio"

	db := mysql.NewDB(config)

	rs := db.FetchAll("SELECT * FROM pio.pio_event_1;")

	for i, r := range rs {
		//r
		fmt.Println(i, ")-----------------------------")
		for k, v := range r.(map[string]string) {
			fmt.Println("mysql", k, v)
		}
		fmt.Println("-----------------------------")
	}
	
}
