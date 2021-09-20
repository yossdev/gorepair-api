package main

import (
	"gorepair-rest-api/config"
	"gorepair-rest-api/routes"
)

func main() {
	config.InitDB() 
	
	e := routes.New()

	e.Logger.Fatal(e.Start(":8000"))
}