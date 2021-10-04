package dto

import "gorepair-rest-api/src/orders/entities"

type OrderRequestBody struct {
	WorkshopID uint64 `json:"workshop_id" validate:"required"`
	ServiceID  uint64 `json:"service_id" validate:"required"`
	OnSite     bool   `json:"on_site"`
	PickUp     bool   `json:"pick_up"`
	Note       string `json:"note" validate:"required"`
}

type OrderUpdateRequestBody struct {
	UserID      uint64 `json:"user_id" validate:"required"`
	ServiceID  	uint64 `json:"service_id" validate:"required"`
	Placed      bool   `json:"placed" validate:"required"`
}

func (req *OrderRequestBody) ToDomain() *entities.Orders {
	return &entities.Orders{
		WorkshopID: req.WorkshopID,
		ServiceID: 	req.ServiceID,
		OnSite: 	req.OnSite,
		PickUp: 	req.PickUp,
		Note: 		req.Note,
	}
}