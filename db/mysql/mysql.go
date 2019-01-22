package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

//参考：http://go-database-sql.org/varcols.html

type DB struct {
	db *sql.DB
}

func NewDB() (mdb *DB) {
	username := "pio"
	password := "pio"
	host := "localhost"
	port := "3306"
	database := "pio"
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=skip-verify&autocommit=true", username, password, host, port, database)

	var db *sql.DB
	var err error

	if db, err = sql.Open("mysql", url); err != nil {
		fmt.Println(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(10 * time.Second)

	mdb = &DB{db}
	return

}

func (d *DB) Close() {

}

func Close() {

}

func fetch(db *sql.DB, query string) (results []interface{}) {
	rows, err := db.Query(query)

	columnTypes, err := rows.ColumnTypes()

	values := make([]interface{}, len(columnTypes))

	for i := range values {
		values[i] = new(sql.RawBytes)
	}

	for rows.Next() {
		err = rows.Scan(values...)
		if err != nil {
			fmt.Println(err)
		}

		result := make(map[string]string, len(columnTypes))

		var value string
		for i, col := range values {
			key := columnTypes[i].Name()

			cp := col.(*sql.RawBytes)
			cv := *cp
			col := cv

			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}

			result[key] = value
		}

		results = append(results, result)

	}

	if err = rows.Err(); err != nil {
		log.Println(err.Error())
		//panic(err.Error()) // proper error handling instead of panic in your app
	}
	return
}

func (m *DB) FetchOne(sql string) (result []interface{}) {
	return fetch(m.db, sql)
}
