package main

import (
	"gorepair-rest-api/config"
	"gorepair-rest-api/infrastructures/db"
	"gorepair-rest-api/infrastructures/local_db"
	"gorepair-rest-api/internal/routes"
	"gorepair-rest-api/internal/web"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	
	app := fiber.New()

	appPort := config.Get().AppPort
	log.Println("Server running on PORT", appPort)

	mysqlDB := db.NewMysqlClient()
	scribleDB := local_db.NewScribleClient()
	db.NewMysqlClient()

	routeStruct := routes.RouterStruct{
		RouterStruct: web.RouterStruct{
			Web:       app,
			MysqlDB:   mysqlDB,
			ScribleDB: scribleDB,
		},
	}
	router := routes.NewHttpRoute(routeStruct)
	router.GetRoutes()

	log.Fatal(app.Listen(":"+appPort))

}