package config

import (
	"sync"
	"time"

	"github.com/spf13/viper"
)

var config appConfigStruct
var doOnce sync.Once

type appConfigStruct struct {
	AppPort 			string
	AppKey  			string // all off local encryption will use this key
	LogPath 			string
	// MySql database config
	DbHost     			string
	DbPort     			string
	DbName     			string
	DbUsername 			string
	DbPassword 			string
	// MongoDB
	MongoDb_Address	   	string
	MongoDb_Name 	   	string
	MongoDb_Collection 	string
	MongoDb_Username 	string
	MongoDb_Password 	string
	// key
	PrivateKey 			string
	PublicKey  			string
	// jwt
	JwtTokenType      	string
	JwtTokenExpired   	time.Duration // in second
	JwtRefreshExpired 	time.Duration // in second
}

func init() {
	doOnce.Do(func() {
		viper.SetConfigFile(`.env`)
		viper.AutomaticEnv()
		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		}

		config = load()
	})
}

func load() appConfigStruct {

	jwtTokenExp := viper.GetString("JWT_TOKEN_EXPIRED")
	jwtRefreshExp := viper.GetString("JWT_REFRESH_EXPIRED")

	jwtTokenDuration, _ := time.ParseDuration(jwtTokenExp)
	jwtRefreshDuration, _ := time.ParseDuration(jwtRefreshExp)

	return appConfigStruct{
		AppPort: 			viper.GetString("APP_PORT"),
		AppKey:  			viper.GetString("APP_KEY"),
		LogPath: 			viper.GetString("LOG_PATH"),
		// db configure
		DbHost:     		viper.GetString("DB_HOST"),
		DbPort:     		viper.GetString("DB_PORT"),
		DbName:     		viper.GetString("DB_NAME"),
		DbUsername: 		viper.GetString("DB_USERNAME"),
		DbPassword: 		viper.GetString("DB_PASSWORD"),
		// MongoDB
		MongoDb_Address:	viper.GetString("MongoDb_Address"),
		MongoDb_Name: 		viper.GetString("MongoDb_Name"),
		MongoDb_Collection: viper.GetString("MongoDb_Collection"),
		MongoDb_Username: 	viper.GetString("MongoDb_Username"),
		MongoDb_Password: 	viper.GetString("MongoDb_Password"),
		// key
		PrivateKey: 		viper.GetString("PRIVATE_KEY"),
		PublicKey:  		viper.GetString("PUBLIC_KEY"),
		// Jwt Configuration
		JwtTokenType:      	"Bearer",
		JwtTokenExpired:   	jwtTokenDuration,   // in second
		JwtRefreshExpired: 	jwtRefreshDuration, // in second
	}
}

func Get() appConfigStruct {
	return config
}