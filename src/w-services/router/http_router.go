package router

import (
	"gorepair-rest-api/infrastructures/third-party/freegeoapi"
	"gorepair-rest-api/src/w-services/handlers"
	"gorepair-rest-api/src/w-services/repositories"
	"gorepair-rest-api/src/w-services/services"
	"log"
)

func NewHttpRoute(structs RouterStruct) RouterStruct {
	log.Println("Setup HTTP Services Route")

	return structs
}

func (r *RouterStruct) GetRoute() {
	wservicesMysqlRepo := repositories.NewWServicesMysqlRepository(r.MysqlDB)
	ipgeo := freegeoapi.NewIpAPI()
	wservicesService := services.NewWServicesService(wservicesMysqlRepo, ipgeo)
	wservicesHandlers := handlers.NewHttpHandler(wservicesService)

	v1 := r.Web.Group("/api/v1/services")
	v1.Get("/", wservicesHandlers.GetAll)
	v1.Get("/workshops", wservicesHandlers.GetAllWorkshop)
	v1.Get("/:serviceId", wservicesHandlers.GetDetails)
}