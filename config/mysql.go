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

	db := mysql.NewDB()

	rs := db.FetchOne("SELECT * FROM pio_meta_accesskeys")

	for _, r := range rs {
		//r
		for k, v := range r.(map[string]string) {
			fmt.Println("mysql", k, v)
		}
	}

	//var (
	//	db   *sql.DB
	//	conn *sql.Conn
	//	err  error
	//)
	//
	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//
	//if db, err = sql.Open("mysql", mysqlUrl); err != nil {
	//	fmt.Println(err)
	//}
	//
	//if err = db.Ping(); err != nil {
	//	fmt.Println(err)
	//}
	//
	//db.SetMaxOpenConns(10)
	//db.SetMaxIdleConns(10)
	//
	//if conn, err = db.Conn(ctx); err != nil {
	//	fmt.Println(err)
	//}
	//
	//var stmt *sql.Stmt
	//
	//if stmt, err = conn.PrepareContext(ctx, "S"); err != nil {
	//
	//}
	//
	//stmt.QueryRow("")
	//
	//fmt.Println(conn)
	//
	//fmt.Println(db)
	//
	//fmt.Println(ctx)

}
