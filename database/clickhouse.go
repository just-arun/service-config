package database

import (
	"database/sql"
	"fmt"
)

// GetClickhouseDB connect to clickhouse db connection
func GetClickhouseDB(host string, port int, debug bool) *sql.DB {
	connectionString := fmt.Sprintf(
		"tcp://%v:%v?debug=%v",
		host,
		port,
		debug,
	)
	connect, err := sql.Open("clickhouse", connectionString)
	if err != nil {
		panic(err)
	}
	fmt.Println("Clickhouse db connected...")
	return connect
}
