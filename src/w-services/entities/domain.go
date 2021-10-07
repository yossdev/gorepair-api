package entities

import _ws "gorepair-rest-api/src/workshops/entities"

type WServices struct {
	ID          uint64
	WorkshopID  uint64
	Vehicle     string
	VehicleType string
	Services    string
	Price       int
}

type WServicesService interface {
	GetAll() ([]WServices, error)
	GetDetails(id string) (WServices, error)
	GetAllWorkshop(ip string) ([]_ws.WorkshopAddress, error)
}

type WServicesMysqlRepositoryInterface interface {
	GetAll() ([]WServices, error)
	GetDetails(id uint64) (WServices, error)
	GetAllWorkshop(city string) ([]_ws.WorkshopAddress, error)
}