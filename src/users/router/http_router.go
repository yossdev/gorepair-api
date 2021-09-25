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

	v1 := r.Web.Group("/api/v1/")
	v1.Get("/user/:username", middleware.JwtVerifyToken, userHandlers.GetUser)
	v1.Post("/user/register", userHandlers.Register)
	v1.Post("/user/login", userHandlers.Login)
	v1.Post("/user/refresh-token", middleware.JwtVerifyRefresh, userHandlers.Refresh)
}