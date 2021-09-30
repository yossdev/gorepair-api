package routes

import (
	"gorepair-rest-api/internal/middleware"
	"gorepair-rest-api/internal/web"
	_userRoute "gorepair-rest-api/src/users/router"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type RouterStruct struct {
	web.RouterStruct
}

func NewHttpRoute(r RouterStruct) RouterStruct {
	log.Println("Loading the HTTP Router")

	return r
}

func (c *RouterStruct) GetRoutes() {

	c.Web.Use(logger.New(), cors.New())
	c.Web.Use(middleware.NewLogMongo(c.MongoDB).LogReqRes)

	c.Web.Get("/api", func(c *fiber.Ctx) error {
		return web.JsonResponse(c, http.StatusOK, "HOMEPAGE", nil)
	})

	webRouterConfig := web.RouterStruct{
		Web:       c.Web,
		MysqlDB:   c.MysqlDB,
		MongoDB:   c.MongoDB,
		ScribleDB: c.ScribleDB,
	}

	// registering route from another modules
	// User Route
	userRouterStruct := _userRoute.RouterStruct {
		RouterStruct: webRouterConfig,
	}
	userRouter := _userRoute.NewHttpRoute(userRouterStruct)
	userRouter.GetRoute()

	// handling 404 error
	c.Web.Use(func(c *fiber.Ctx) error {
		return web.JsonResponse(c, http.StatusNotFound, "Sorry can't find that!", nil)
	})
}