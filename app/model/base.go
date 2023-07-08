package model

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {

	//接続情報を格納
	connect := "host=db port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"

	var err error
	Db, err = sql.Open("postgres", connect)
	if err != nil {
		fmt.Println("failed!!!!!!!!!!")
	} else {
		fmt.Println("db connected!!")
	}

}
