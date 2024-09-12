package config

import (
	"github.com/spf13/viper"
)

func Init(env string) {

	if env != "production" && env != "staging" && env != "local" {
		env = "local"
	}

	envConfigName := env + ".config.yml"

	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/")

	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	viper.SetConfigName(envConfigName)

	err = viper.MergeInConfig()
	if err != nil {
		panic(err)
	}
}
