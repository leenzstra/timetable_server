package config

import "github.com/spf13/viper"

type Config struct {
	PostgresLogin  string `mapstructure:"POSTGRES_LOGIN"`
	PostgresPass string `mapstructure:"POSTGRES_PASS"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./common/config")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	return
}