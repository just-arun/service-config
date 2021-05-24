package database

import (
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
)

func getCockroach(userName, password, hostName string, port int) *pgx.ConnConfig {
	connectionString := fmt.Sprintf(
		"postgres://%v:%v@%v:%v/bank?sslmode=require",
		userName,
		password,
		hostName,
		port,
	)
	config, err := pgx.ParseConfig(
		connectionString,
	)
	if err != nil {
		log.Fatal("error configuring the database: ", err)
	}
	return config
}
