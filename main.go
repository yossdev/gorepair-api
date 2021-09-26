package main

import (
	"gorepair-rest-api/config"
	"gorepair-rest-api/infrastructures/db"
	"gorepair-rest-api/infrastructures/local_db"
	"gorepair-rest-api/internal/routes"
	s "gorepair-rest-api/internal/utils/start-server"
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
	// mongoDB := db.NewMongoClient()

	routeStruct := routes.RouterStruct{
		RouterStruct: web.RouterStruct{
			Web:       app,
			MysqlDB:   mysqlDB,
			// MongoDB:   mongoDB,
			ScribleDB: scribleDB,
		},
	}
	router := routes.NewHttpRoute(routeStruct)
	router.GetRoutes()

	s.StartServer(app)
	// s.StartServerWithGracefulShutdown(app)

}