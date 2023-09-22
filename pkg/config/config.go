package config

import (
	"github.com/spf13/viper"
)

var Instance = &config{}

type config struct {
	Environment string
}

func init() {
	viper.AutomaticEnv()

	Instance.Environment = viper.GetString("ENVIRONMENT")
}
