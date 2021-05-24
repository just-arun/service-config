package database

import (
	"fmt"

	aero "github.com/aerospike/aerospike-client-go"
)

func getAerospikeInstance(host string, port int) *aero.Client {
	client, err := aero.NewClient(host, port)
	if err != nil {
		fmt.Println("[DB ERR]: aerospike db connection error")
		panic(err)
	}
	fmt.Println("Aerospike connection established...")
	return client
}
