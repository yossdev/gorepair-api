package router

import (
	"gorepair-rest-api/internal/middleware"
	"gorepair-rest-api/internal/utils/auth"
	"gorepair-rest-api/src/orders/handlers"
	"gorepair-rest-api/src/orders/repositories"
	"gorepair-rest-api/src/orders/services"
	"log"
)

func NewHttpRoute(structs RouterStruct) RouterStruct {
	log.Println("Setup HTTP Orders Route")

	structs.jwtAuth = auth.NewJwt(structs.ScribleDB)

	return structs
}

func (r *RouterStruct) GetRoute() {
	orderMysqlRepo := repositories.NewOrderMysqlRepository(r.MysqlDB)
	orderScribleRepo := repositories.NewOrderScribleRepositoryInterface(r.ScribleDB)
	orderService := services.NewOrderService(orderMysqlRepo, orderScribleRepo)
	orderHandlers := handlers.NewHttpHandler(orderService)

	v1 := r.Web.Group("/api/v1/orders")
	v1.Post("/", middleware.JwtVerifyRefresh, middleware.UserRestricted, orderHandlers.OrderNew)
	v1.Get("/user/:orderId", middleware.JwtVerifyRefresh, middleware.UserRestricted, orderHandlers.GetUserOrderDetails)
	v1.Get("/workshop/:orderId", middleware.JwtVerifyRefresh, middleware.WorkshopRestricted, orderHandlers.GetWorkshopOrderDetails)
	v1.Delete("/user/:username/:orderId/cancel", middleware.JwtVerifyRefresh, middleware.UserRestricted, orderHandlers.UserCancelOrder)
}