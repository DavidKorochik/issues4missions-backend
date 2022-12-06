package config

import "github.com/spf13/viper"

type Config struct {
	DBUrl string `mapstructure:"DB_URL"`
	Port  string `mapstructure:"PORT"`
}

func LoadEnvVariables(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("server")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
