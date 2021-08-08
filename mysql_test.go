package golib

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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

}

func Test_main(t *testing.T) {
	initMysql()
	assert.True(t, Connection != nil)
	defer Connection.Disconnect()
}
