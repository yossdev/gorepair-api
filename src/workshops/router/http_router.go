package router

import (
	"gorepair-rest-api/internal/middleware"
	"gorepair-rest-api/internal/utils/auth"
	"gorepair-rest-api/src/workshops/handlers"
	"gorepair-rest-api/src/workshops/repositories"
	"gorepair-rest-api/src/workshops/services"
	"log"
)

func NewHttpRoute(structs RouterStruct) RouterStruct {
	log.Println("Setup HTTP Workshops Route")

	structs.jwtAuth = auth.NewJwt(structs.ScribleDB)

	return structs
}

func (r *RouterStruct) GetRoute() {
	workshopMysqlRepo := repositories.NewWorkshopMysqlRepository(r.MysqlDB)
	workshopScribleRepo := repositories.NewWorkshopScribleRepositoryInterface(r.ScribleDB)
	workshopService := services.NewWorkshopService(workshopMysqlRepo, workshopScribleRepo, r.jwtAuth)
	workshopHandlers := handlers.NewHttpHandler(workshopService)

	v1 := r.Web.Group("/api/v1/workshops")
	v1.Post("/", workshopHandlers.Login)
	v1.Get("/:username/logout", middleware.JwtVerifyRefresh, middleware.WorkshopRestricted, workshopHandlers.Logout)
	v1.Post("/register", workshopHandlers.Register)
	v1.Get("/:username", middleware.JwtVerifyRefresh, workshopHandlers.GetWorkshop)
	v1.Put("/:username/account", middleware.JwtVerifyRefresh, middleware.WorkshopRestricted, workshopHandlers.UpdateAccount)
	v1.Put("/:username/address/update", middleware.JwtVerifyRefresh, middleware.WorkshopRestricted, workshopHandlers.UpdateAddress)
	v1.Get("/:username/address", middleware.JwtVerifyRefresh, workshopHandlers.GetAddress)
	v1.Put("/:username/description", middleware.JwtVerifyRefresh, middleware.WorkshopRestricted, workshopHandlers.UpdateDescription)

	v1.Post("/:username/services", middleware.JwtVerifyRefresh, middleware.WorkshopRestricted, workshopHandlers.ServicesNew)
	v1.Put("/:username/services/:serviceId", middleware.JwtVerifyRefresh, middleware.WorkshopRestricted, workshopHandlers.UpdateServices)
	v1.Delete("/:username/services/:serviceId", middleware.JwtVerifyRefresh, middleware.WorkshopRestricted, workshopHandlers.DeleteServices)
}