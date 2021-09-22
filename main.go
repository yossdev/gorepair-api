package main

import (
	"gorepair-rest-api/config"
	"gorepair-rest-api/routes"

	"github.com/spf13/viper"
)

func main() {
	config.LoadConfig()
	config.InitDB()
	config.InitMongo()

	e := routes.New()

	e.Logger.Fatal(e.Start(viper.GetString(`SERVER_ADDRESS`)))
}