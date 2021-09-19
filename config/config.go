package config

import "github.com/spf13/viper"

type Config struct {
	HERE_API_KEY string `mapstructure: "HERE_API_KEY"`
	DBUsername   string `mapstructure: "DB_USER"`
	DBPassword   string `mapstructure: "DB_PASSWORD"`
	DBHost       string `mapstructure: "DB_HOST"`
	DBPort       string `mapstructure: "DB_PORT"`
	DBName       string `mapstructure: "DB_NAME"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}