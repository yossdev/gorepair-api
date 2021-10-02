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
	Price       uint64
}

type WorkshopService interface {
	FindByID(id string) error
	GetWorkshop(username string) (*Workshops, error)
	Register(payload *Workshops, street string) (*Workshops, error)
	Login(payload *Workshops) (interface{}, error)
	Logout(id string) error
	UpdateAccount(payload *Workshops, id uint64) (*Workshops, error)
	UpdateAddress(payload *WorkshopAddress, id uint64) (*WorkshopAddress, error)
	GetAddress(id uint64) (*WorkshopAddress, error)
}

type WorkshopMysqlRepositoryInterface interface {
	GetWorkshop(username string) (*Workshops, error)
	Register(payload *Workshops, street string) (*Workshops, error)
	FindByEmail(email string) *Workshops
	UpdateAccount(payload *Workshops, id uint64) (*Workshops, error)
	UpdateAddress(payload *WorkshopAddress, id uint64) (*WorkshopAddress, error)
	GetAddress(id uint64) (*WorkshopAddress, error)
}

type WorkshopScribleRepositoryInterface interface {
	FindWorkshopRefreshToken(workshopID string) error
	DeleteWorkshopRefreshToken(workshopID string) error
}