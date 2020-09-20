package config

import (
	"github.com/spf13/viper"
)

type Configuration struct {
	vp *viper.Viper
}

func NewConfiguration() (*Configuration, error) {
	vp := viper.New()
    vp.SetConfigName("config")
	vp.AddConfigPath("config/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Configuration{vp}, nil
}