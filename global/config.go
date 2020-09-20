package global

import (
	"github.com/hwb2017/CMDBDemo/lib/config"
)

var (
	DatabaseConfiguration config.DatabaseConfig
	ServerConfiguration config.ServerConfig
	LogConfiguration config.LogConfig
	CloudApiConfiguration config.CloudApiConfig
)

func InitConfiguration() {
	configuration, err := config.NewConfiguration()
	configuration.ReadSection("Server", &ServerConfiguration)
	if err != nil {
		panic(err)
	}
	configuration.ReadSection("Database", &DatabaseConfiguration)
	if err != nil {
		panic(err)
	}
	configuration.ReadSection("Log", &LogConfiguration)
	envConfig := config.ReadEnvVars()
	CloudApiConfiguration = envConfig.CloudApiConfig
}