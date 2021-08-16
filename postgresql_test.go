package main

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func postgresqlInit() {
	config := loadConfigFileName(".", ".env")
	fmt.Println(config)
	ConnectInstance = NewInstance("postgres", config)
	if ConnectInstance.GetConnectInstance().(*sql.DB) != nil {
		fmt.Println("Connected")
	} else {
		fmt.Println("Not Connect")
	}

}

func Test_postgresql(t *testing.T) {
	postgresqlInit()
	assert.True(t, ConnectInstance != nil)
	defer ConnectInstance.Disconnect()
}
