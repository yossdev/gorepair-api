package main

import (
	"gorepair-rest-api/config"
	"gorepair-rest-api/infrastructures/db"
	"gorepair-rest-api/infrastructures/local_db"
	"gorepair-rest-api/internal/routes"
	"gorepair-rest-api/internal/web"
	"gorepair-rest-api/src/users/entities"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func DbMigrate(db *gorm.DB) {
	db.AutoMigrate(&entities.User{})
}

func main() {
	// you didn't define port in env file
	// the default port is random from fiber
	
	appPort := config.Get().AppPort
	log.Println("Server running on PORT", appPort)
	app := fiber.New()

	mysqlDB := db.NewMysqlClient()
	scribleDB := local_db.NewScribleClient()

	DbMigrate(db.NewMysqlClient().DB())

	routeStruct := routes.RouterStruct{
		RouterStruct: web.RouterStruct{
			Web:       app,
			MysqlDB:   mysqlDB,
			ScribleDB: scribleDB,
		},
	}
	router := routes.NewHttpRoute(routeStruct)
	router.GetRoutes()

	log.Fatal(app.Listen(":3000"))
	// app.Listen(fmt.Sprintf(":%s", appPort))
}