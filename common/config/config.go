package config

import "github.com/spf13/viper"

type Config struct {
	PostgresLogin  string `mapstructure:"POSTGRES_LOGIN"`
	PostgresPass string `mapstructure:"POSTGRES_PASS"`
	PostgresHost string `mapstructure:"POSTGRES_HOST"`
	PostgresPort int `mapstructure:"POSTGRES_PORT"`
}

func LoadConfig(prod bool) (c Config, err error) {
	viper.AddConfigPath("./common/config")
	if prod {
		viper.SetConfigName("prod")
	} else {
		viper.SetConfigName("dev")
	}
	
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	return
}