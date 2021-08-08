package main

import (
	"database/sql"
	"fmt"
)

func initMysql() {
	config := loadConfigFileName(".", ".env")
	fmt.Println(config)
	Connection = New(config)
	if Connection.GetConnection().(*sql.DB) != nil {
		fmt.Println("Connected")
	} else {
		fmt.Println("Not Connect")
	}
	defer Connection.Disconnect()
}

func main() {
	initMysql()
}
