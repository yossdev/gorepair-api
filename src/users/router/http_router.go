package router

import (
	"gorepair-rest-api/internal/middleware"
	"gorepair-rest-api/internal/utils/auth"
	"gorepair-rest-api/src/users/handlers"
	"gorepair-rest-api/src/users/repositories"
	"gorepair-rest-api/src/users/services"
	"log"
)

func NewHttpRoute(
	structs RouterStruct,
) RouterStruct {
	log.Println("Setup HTTP Users Route")

	structs.jwtAuth = auth.NewJwt(structs.ScribleDB)

	return structs
}

func (r *RouterStruct) GetRoute() {
	userMysqlRepo := repositories.NewUserMysqlRepository(r.MysqlDB)
	userScribleRepo := repositories.NewUserScribleRepositoryInterface(r.ScribleDB)
	userService := services.NewUserService(userMysqlRepo, r.jwtAuth, userScribleRepo)
	userHandlers := handlers.NewHttpHandler(userService)

	// r.Web.Get("/user/:username", userHandlers.GetUser)
	r.Web.Post("/user/register", userHandlers.Register)
	r.Web.Post("/user/login", userHandlers.Login)
	r.Web.Post("/user/refresh-token", middleware.JwtVerifyRefresh, userHandlers.Refresh)
}