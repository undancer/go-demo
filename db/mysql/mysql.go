package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

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

	columns, err := rows.Columns()

	columnTypes, err := rows.ColumnTypes()

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))

	//cols, err := rows.Columns() // Remember to check err afterwards
	//vals := make([]interface{}, len(cols))
	//for i, _ := range cols {
	//	vals[i] = new(sql.RawBytes)
	//}

	for i := range scanArgs {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			fmt.Println("sss", err)
		}

		result := make(map[string]string, len(columnTypes))

		for i, col := range scanArgs {
			cp := col.(*sql.RawBytes)
			cv := *cp

			fmt.Println("scan", i, cv)
		}

		for i, col := range values {
			//c := &col

			fmt.Println("vals", i, col)
		}

		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}

			result[columns[i]] = value

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
