package main

import (
	"encoding/json"
	"fmt"
	"go_api/model"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// APIサーバーからデータを取得
	resp, err := http.Get("http://localhost:8080/users")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	var user []model.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(user)

}
