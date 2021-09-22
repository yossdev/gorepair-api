package main

import (
	"gorepair-rest-api/config"
	"gorepair-rest-api/middlewares"
	"gorepair-rest-api/routes"

	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.InitDB()
	config.InitMongo()

	e := routes.New()

	e.Use(middleware.BodyDump(middlewares.BodyDumpLog))

	e.Logger.Fatal(e.Start(":8000"))
}