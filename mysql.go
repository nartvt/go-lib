package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	conn *sql.DB
}

const URL = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"

var Connection *MySQL

func New(config ...Config) *MySQL {
	c := &MySQL{}
	if len(config) > 0 {
		_config := config[0]
		c.Connect(_config)
	}

	return c
}

func (c *MySQL) GetConnection() interface{} {
	return c.conn
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
	conn, err := sql.Open(config.DbDriver, address)
	if err != nil {
		log.Fatal("Failed to connect to MySQL database", err)
	}
	c.conn = conn
}

func (c *MySQL) Disconnect() {
	err := c.conn.Close()
	if err != nil {
		log.Fatal("Failed to connect to MySQL database", err)
	}
	log.Println("Disconnected from MariaDB")
}
