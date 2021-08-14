package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	db *sql.DB
}

const URL = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"

var MysqlConnect *MySQL

func NewConnectionMySQL(config ...Config) *MySQL {
	c := &MySQL{}
	if len(config) > 0 {
		_config := config[0]
		c.Connect(_config)
	}

	return c
}

func (c *MySQL) GetConnectionMySQL() interface{} {
	return c.db
}

func (c *MySQL) Connect(config Config) {
	address := fmt.Sprintf(
		URL,
		config.DbUsername,
		config.DbPassword,
		config.DbHost,
		config.DbPort,
		config.DbName,
	)
	db, err := sql.Open(config.DbDriver, address)
	if err != nil {
		log.Fatal("Failed to connect to MySQL database", err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	c.db = db
}

func (c *MySQL) Disconnect() {
	err := c.db.Close()
	if err != nil {
		log.Fatal("Failed to connect to MySQL database", err)
	}
	log.Println("Disconnected from MySQL")
}
