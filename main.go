package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func initPostgresql() {
	config := loadConfigFileName(".", ".env")
	fmt.Println(config)
	PostgreConnect = NewConnectPostgresSQL(config)
	if PostgreConnect.GetConnectPostgresSQL().(*sql.DB) != nil {
		fmt.Println("Connected")
	} else {
		fmt.Println("Not Connect")
	}
}

type User struct {
	Id     int8
	Number uint64
}

func main() {
	initPostgresql()
	sqlString := "select id, number from test_table"
	sqlQuery := PostgreConnect.GetConnectPostgresSQL().(*sql.DB)
	row := sqlQuery.QueryRow(sqlString)
	users := make([]User, 2)
	if row != nil {
		var user User
		if err := row.Scan(&user.Id, &user.Number); err != nil {
			fmt.Println(err.Error())
		} else {
			users = append(users, user)
		}

	}
	fmt.Println(users)

	//defer sqlQuery.Close()
}
