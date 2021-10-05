package entities

type Workshops struct {
	ID               uint64
	Username         string
	Email            string
	Password         string
	Name             string
	Phone            string
	OperationalStart string
	OperationalEnd   string
	Description      Descriptions
	WorkshopAddress
	Services
}

type WorkshopAddress struct {
	ID             uint64
	WorkshopID     uint64
	BuildingNumber string
	Street         string
	City           string
	Country        string
	PostalCode     string
	Province       string
}

type Descriptions struct {
	ID          uint64
	WorkshopID  uint64
	Description string
}

type Services struct {
	ID          uint64
	WorkshopID  uint64
	Vehicle     string
	VehicleType string
	Services    string
	Price       int
}

type WorkshopService interface {
	GetWorkshop(username string) (*Workshops, error)
	Register(payload *Workshops, street, description string) (*Workshops, error)
	Login(payload *Workshops) (interface{}, error)
	Logout(ctxId, username string) error
	UpdateAccount(payload *Workshops, username string) (*Workshops, error)
	UpdateAddress(payload *WorkshopAddress, username string) (*WorkshopAddress, error)
	GetAddress(username string) (*WorkshopAddress, error)
	UpdateDescription(payload *Descriptions, username string) (*Descriptions, error)
	ServicesNew(payload *Services, username string) (*Services, error)
	UpdateServices(payload *Services, username, servicesId string) (*Services, error)
	DeleteServices(username, servicesId string) error
}

type WorkshopMysqlRepositoryInterface interface {
	GetWorkshop(username string) (*Workshops, error)
	Register(payload *Workshops, street, description string) (*Workshops, error)
	FindByEmail(email string) *Workshops
	UpdateAccount(payload *Workshops, id uint64) (*Workshops, error)
	UpdateAddress(payload *WorkshopAddress, id uint64) (*WorkshopAddress, error)
	GetAddress(id uint64) (*WorkshopAddress, error)
	UpdateDescription(payload *Descriptions, id uint64) (*Descriptions, error)
	ServicesNew(payload *Services, id uint64) (*Services, error)
	UpdateServices(payload *Services, id uint64, servicesId uint64) (*Services, error)
	DeleteServices(id, servicesId uint64) error
}

type WorkshopScribleRepositoryInterface interface {
	FindWorkshopRefreshToken(workshopID string) error
	DeleteWorkshopRefreshToken(workshopID string) error
}