package painterconfig

type server struct {
	Port string `toml:"port"`
	Host string `toml:"host"`
}

type mongodb struct {
	MongoURI string `toml:"mongoUri"`
}

type aeroSpike struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

type database struct {
	Mongo     mongodb   `toml:"mongo"`
	Aerospike aeroSpike `toml:"aeroSpike"`
}

type Config struct {
	Server   server   `toml:"server"`
	Database database `toml:"database"`
}
