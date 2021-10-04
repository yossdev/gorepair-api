package repositories

import (
	"errors"
	"gorepair-rest-api/infrastructures/db"
	"gorepair-rest-api/src/orders/entities"
	_user "gorepair-rest-api/src/users/repositories"
	_service "gorepair-rest-api/src/workshops/repositories"
)

type orderMysqlRepository struct {
	DB db.MysqlDB
}

func NewOrderMysqlRepository(DB db.MysqlDB) entities.OrderMysqlRepositoryInterface {
	return &orderMysqlRepository{
		DB: DB,
	}
}

func (u *orderMysqlRepository) OrderNew(payload *entities.Orders, userId uint64) (*entities.Orders, error) {
	order := Order{}
	service := _service.Service{}

	price := u.DB.DB().First(&service, "id = ?", payload.ServiceID)
	if price.Error != nil {
		return nil, price.Error
	}

	fromDomainOrder(payload, &order)
	order.UserID = userId
	order.TotalPrice = service.Price

	res := u.DB.DB().Save(&order)
	if res.Error != nil {
		return nil, res.Error
	}

	return order.toDomain(), nil
}

func (u *orderMysqlRepository) GetUserOrderDetails(orderId, userId uint64) (*entities.Orders, error) {
	order := Order{}

	res := u.DB.DB().Where("id = ? AND user_id = ?", orderId, userId).Find(&order)
	if res.Error != nil {
		return nil, res.Error
	}

	user := _user.User{}
	if res := u.DB.DB().Preload("Address").First(&user, "id = ?", order.UserID); res.Error != nil {
		return nil, res.Error
	}
	
	return order.toDomainGet(user.Address.Street, user.Phone), nil
}

func (u *orderMysqlRepository) GetWorkshopOrderDetails(orderId, workshopId uint64) (*entities.Orders, error) {
	order := Order{}
	res := u.DB.DB().Where("id = ? AND workshop_id = ?", orderId, workshopId).Find(&order)
	if res.Error != nil {
		return nil, res.Error
	}

	user := _user.User{}
	if res := u.DB.DB().Preload("Address").First(&user, "id = ?", order.UserID); res.Error != nil {
		return nil, res.Error
	}

	return order.toDomainGet(user.Address.Street, user.Phone), nil
}

func (u *orderMysqlRepository) UserCancelOrder(orderId, userId uint64, username string) error {
	user := _user.User{}
	if res := u.DB.DB().First(&user, "username = ?", username); res.Error != nil {
		return res.Error
	}

	if userId != user.ID {
		return errors.New("")
	}

	order := Order{}
	res := u.DB.DB().First(&order, "id = ? AND user_id = ?", orderId, user.ID)
	if res.Error != nil {
		return res.Error
	}

	del := u.DB.DB().Delete(&order)
	if del.Error != nil {
		return del.Error
	}

	return nil
}