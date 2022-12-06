package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBUrl string `mapstructure:"DB_URL"`
	Port  string `mapstructure:"PORT"`
}

func LoadEnvVariables(path string) (config Config) {
	viper.AddConfigPath(path)
	viper.SetConfigName("server")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal("Cannot read env file ", err)
		return
	}

	err = viper.Unmarshal(&config)

	return
}
