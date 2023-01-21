package main

import (
	"fmt"
	"persona-backend-restfulapi/database"
)

func main() {
	db := database.Connect()
	//接続が終了したらクローズ
	defer db.Close()

	err := db.Ping()

	if err != nil {
		fmt.Println("データベース接続失敗")
		return
	} else {
		fmt.Println("データベース接続成功")
	}
}
