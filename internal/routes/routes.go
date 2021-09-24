package routes

import (
	"gorepair-rest-api/internal/web"
	userRoute "gorepair-rest-api/src/users/router"
	"log"

	"github.com/gofiber/fiber/v2"
)

type RouterStruct struct {
	web.RouterStruct
}

func NewHttpRoute(r RouterStruct) RouterStruct {
	log.Println("Loading the HTTP Router")

	return r
}

func (c *RouterStruct) GetRoutes() {
	c.Web.Get("/", func(c *fiber.Ctx) error {
		return c.Send([]byte("Hello this is my first route in go fiber"))
	})

	webRouterConfig := web.RouterStruct{
		Web:       c.Web,
		MysqlDB:   c.MysqlDB,
		ScribleDB: c.ScribleDB,
	}
	// registering route from another modules
	userRouterStruct := userRoute.RouterStruct {
		RouterStruct: webRouterConfig,
	}
	userRouter := userRoute.NewHttpRoute(userRouterStruct)
	userRouter.GetRoute()

	// handling 404 error
	c.Web.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Sorry can't find that!")
	})
}