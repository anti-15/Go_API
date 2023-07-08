package main

import (
	"database/sql"
	"go_api/controllers"
	"go_api/model"
	"log"
)

var Db *sql.DB

func main() {
	controllers.StartServer()

	// init関数で閉じれば接続エラーになるので、main関数でdbを閉じる。
	err := model.Db.Close()
	if err != nil {
		log.Fatal(err)
	}

}
