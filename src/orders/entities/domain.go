package entities

type Orders struct {
	ID         uint64
	UserID     uint64
	Phone      string
	Street     string
	WorkshopID uint64
	ServiceID  uint64
	OnSite     bool
	PickUp     bool
	Note       string
	TotalPrice int
	Placed     bool
}

type OrderService interface {
	FindUserByID(id string) error
	FindWorkshopByID(id string) error
	OrderNew(payload *Orders, userId string) (*Orders, error)
	GetUserOrderDetails(orderId, userId string) (*Orders, error)
	GetWorkshopOrderDetails(orderId, workshopId string) (*Orders, error)
	UserCancelOrder(orderId, userId, username string) error
}

type OrderMysqlRepositoryInterface interface {
	OrderNew(payload *Orders, userId uint64) (*Orders, error)
	GetUserOrderDetails(orderId, userId uint64) (*Orders, error)
	GetWorkshopOrderDetails(orderId, workshopId uint64) (*Orders, error)
	UserCancelOrder(orderId, userId uint64, username string) error
}

type OrderScribleRepositoryInterface interface {
	FindWorkshopRefreshToken(id string) error
	FindUserRefreshToken(id string) error
}