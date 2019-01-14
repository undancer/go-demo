package mysql

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

const (
	url      = "tcp(localhost)"
	username = "pio"
	password = "pio"
)

func Main() {
	var (
		cxt  context.Context
		db   *sql.DB
		conn *sql.Conn
		err  error
	)
	cxt, _ = context.WithTimeout(context.Background(), 10*time.Second)
	if db, err = sql.Open("mysql", "pio:pio@tcp(localhost)/pio"); err != nil {
		log.Println(err.Error())
	}

	if conn, err = db.Conn(cxt); err != nil {
		log.Println(err.Error())
	}
	defer conn.Close()

	rows, err := conn.QueryContext(cxt, "SELECT * FROM pio_meta_accesskeys")

	columns, err := rows.Columns()

	if err != nil {

	}
	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {

		}

		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}

			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("-----------------------------------")

	}
	if err = rows.Err(); err != nil {
		log.Println(err.Error())
		//panic(err.Error()) // proper error handling instead of panic in your app
	}
}
