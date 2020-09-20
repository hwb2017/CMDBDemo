package global

import (
	"github.com/hwb2017/CMDBDemo/lib/config"
)

var (
	DatabaseConfiguration config.DatabaseConfig
	ServerConfiguration config.ServerConfig
)

func InitConfiguration() {
	config, err := config.NewConfiguration()
	config.ReadSection("Server", &ServerConfiguration)
	if err != nil {
		panic(err)
	}
	config.ReadSection("Database", &DatabaseConfiguration)
	if err != nil {
		panic(err)
	}
}