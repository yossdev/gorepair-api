package entities

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
}

type WServicesMysqlRepositoryInterface interface {
	GetAll() ([]WServices, error)
	GetDetails(id uint64) (WServices, error)
}