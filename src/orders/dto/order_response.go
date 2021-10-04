package dto

import "gorepair-rest-api/src/orders/entities"

type OrderResponseBody struct {
	ID          uint64 `json:"id"`
	UserID      uint64 `json:"user_id"`
	Phone 		string `json:"phone"`
	Street 		string `json:"street"`
	WorkshopID  uint64 `json:"workshop_id"`
	ServiceID  	uint64 `json:"service_id"`
	OnSite      bool   `json:"on_site"`
	PickUp      bool   `json:"pick_up"`
	Note        string `json:"note"`
	TotalPrice  int    `json:"total_price"`
	Placed      bool   `json:"placed"`
}

func FromDomainOrder(domain *entities.Orders) OrderResponseBody {
	return OrderResponseBody{
		ID:             domain.ID,
		UserID: 		domain.UserID,
		WorkshopID: 	domain.WorkshopID,
		ServiceID: 		domain.ServiceID,
		OnSite: 		domain.OnSite,
		PickUp: 		domain.PickUp,
		Note: 			domain.Note,
		TotalPrice: 	domain.TotalPrice,
		Placed: 		domain.Placed,
	}
}

func FromDomainOrderGet(domain *entities.Orders) OrderResponseBody {
	return OrderResponseBody{
		ID:             domain.ID,
		UserID: 		domain.UserID,
		Phone: 			domain.Phone,
		Street: 		domain.Street,
		WorkshopID: 	domain.WorkshopID,
		ServiceID: 		domain.ServiceID,
		OnSite: 		domain.OnSite,
		PickUp: 		domain.PickUp,
		Note: 			domain.Note,
		TotalPrice: 	domain.TotalPrice,
		Placed: 		domain.Placed,
	}
}