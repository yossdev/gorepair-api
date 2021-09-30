package main

import (
	"gorepair-rest-api/config"
	"gorepair-rest-api/infrastructures/db"
	"gorepair-rest-api/infrastructures/local_db"
	"gorepair-rest-api/internal/routes"
	_s "gorepair-rest-api/internal/utils/start-server"
	"gorepair-rest-api/internal/web"
	"gorepair-rest-api/src/users/repositories"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&repositories.User{},
		&repositories.UserAddress{},
	)
}

// @title API
// @version 1.0
// @BasePath /api
func main() {

	app := fiber.New()

	appPort := config.Get().AppPort
	log.Println("Server running on PORT", appPort)

	mysqlDB := db.NewMysqlClient()
	dbMigrate(mysqlDB.DB())

	scribleDB := local_db.NewScribleClient()
	mongoDB := db.NewMongoClient()

	routeStruct := routes.RouterStruct{
		RouterStruct: web.RouterStruct{
			Web:       app,
			MysqlDB:   mysqlDB,
			MongoDB:   mongoDB,
			ScribleDB: scribleDB,
		},
	}
	router := routes.NewHttpRoute(routeStruct)
	router.GetRoutes()

	_s.StartServer(app)
	// _s.StartServerWithGracefulShutdown(app)

}