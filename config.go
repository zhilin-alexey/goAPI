package main

import "github.com/spf13/viper"

type Config struct {
	PostgresURL    string `mapstructure:"POSTGRES_URL"`
	GinLogLevel    string `mapstructure:"GIN_MODE"`
	TrustedProxies string `mapstructure:"TRUSTED_PROXIES"`
	Port           string `mapstructure:"PORT"`
	LogLevel       string `mapstructure:"LOG_LEVEL"`
}

func LoadConfig() (config Config, err error) {

	viper.SetConfigFile(".env")

	viper.SetDefault("PORT", "8080")
	viper.SetDefault("TRUSTED_PROXIES", "127.0.0.1 192.168.1.2 10.0.0.0/8")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
