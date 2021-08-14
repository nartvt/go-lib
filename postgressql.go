package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type PostgreSQL struct {
	db *sql.DB
}

const PostgresqlUrl = "postgres://%s:%s@%s:%s/%s?sslmode=disable"

var PostgreConnect *PostgreSQL

func NewConnectPostgresSQL(config ...Config) *PostgreSQL {
	c := &PostgreSQL{}
	if len(config) > 0 {
		_config := config[0]
		c.ConnectPostgres(_config)
	}

	return c
}

func (c *PostgreSQL) GetConnectPostgresSQL() interface{} {
	return c.db
}

func (c *PostgreSQL) ConnectPostgres(config Config) {
	psqlInfo := fmt.Sprintf(
		PostgresqlUrl,
		config.DbUsername,
		config.DbPassword,
		config.DbHost,
		config.DbPort,

		config.DbName,
	)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	c.db = db
}

func (c *PostgreSQL) Close() {
	err := c.db.Close()
	if err != nil {
		log.Fatal("Failed to connect to Postgresql database", err)
	}
	log.Println("Disconnected from Postgresql")
}
