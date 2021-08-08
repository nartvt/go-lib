package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MariaDB struct {
	conn *sql.DB
}

const URL = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"

var Connection *MariaDB

func New(config ...Config) *MariaDB {
	c := &MariaDB{}
	if len(config) > 0 {
		_config := config[0]
		c.Connect(_config)
	}

	return c
}

func (c *MariaDB) GetConnection() interface{} {
	return c.conn
}

func (c *MariaDB) Connect(config Config) {
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

func (c *MariaDB) Disconnect() {
	err := c.conn.Close()
	if err != nil {
		log.Fatal("Failed to connect to MySQL database", err)
	}
	log.Println("Disconnected from MariaDB")
}
