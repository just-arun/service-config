package database

import (
	"database/sql"

	aero "github.com/aerospike/aerospike-client-go"
	"github.com/dgraph-io/dgo"
	"github.com/jackc/pgx/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type CancelFunc func()

type DB struct {
	AerospikeClient *aero.Client
	CockroachConnectConfig *pgx.ConnConfig
	ClickHouseClient *sql.DB
	DGraphClient *dgo.Dgraph
	Mongo *mongo.Client
}


func (s *DB) InitClickHouse(host string, port int, debug bool) *DB {
	s.ClickHouseClient = getClickhouseDB(host, port, debug)
	return s
}
func (s *DB) InitCockroach(userName, password, hostName string, port int) *DB {
	s.CockroachConnectConfig = getCockroach(userName, password, hostName, port)
	return s
}
func (s *DB) InitAeroSpike(host string, port int) *DB {
	s.AerospikeClient = getAerospikeInstance(host, port)
	return s
}
func (s *DB) InitDGraph(param GetDGraphDBParam) *DB {
	client, _ := getDGraphDB(param)
	s.DGraphClient = client
	return s
}
func (s *DB) InitMongo(mongoUri string) *DB {
	s.Mongo = getMongoInstance(mongoUri)
	return s
}
