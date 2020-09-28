package initialize

import (
	"github.com/hwb2017/CMDBDemo/global"
	"github.com/hwb2017/CMDBDemo/lib/config"
)

func InitConfiguration() {
	configuration, err := config.NewConfiguration()
	configuration.ReadSection("Server", &global.ServerConfiguration)
	if err != nil {
		panic(err)
	}
	configuration.ReadSection("Database", &global.DatabaseConfiguration)
	if err != nil {
		panic(err)
	}
	configuration.ReadSection("Log", &global.LogConfiguration)
	envConfig := config.ReadEnvVars()
	global.CloudApiConfiguration = envConfig.CloudApiConfig
}