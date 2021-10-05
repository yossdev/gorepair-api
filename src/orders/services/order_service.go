package services

import (
	"errors"
	"gorepair-rest-api/src/orders/entities"
	"strconv"
)

type orderService struct {
	orderMysqlRepository   entities.OrderMysqlRepositoryInterface
	orderScribleRepository entities.OrderScribleRepositoryInterface
}

func NewOrderService(
	orderMysqlRepository entities.OrderMysqlRepositoryInterface,
	orderScribleRepository entities.OrderScribleRepositoryInterface,
) entities.OrderService {
	return &orderService{
		orderMysqlRepository:   orderMysqlRepository,
		orderScribleRepository: orderScribleRepository,
	}
}

func (c *orderService) OrderNew(payload *entities.Orders, userId string) (*entities.Orders, error) {
	uID, err := strconv.ParseUint(userId, 10, 64)
	if err != nil {
		return nil, err
	}

	if err := c.orderScribleRepository.FindUserRefreshToken(userId); err != nil {
		return nil, err
	}

	res, err := c.orderMysqlRepository.OrderNew(payload, uID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *orderService) GetUserOrderDetails(orderId, userId string) (*entities.Orders, error) {
	orderID, err := strconv.ParseUint(orderId, 10, 64)
	if err != nil {
		return nil, err
	}

	userID, err := strconv.ParseUint(userId, 10, 64)
	if err != nil {
		return nil, err
	}

	if err := c.orderScribleRepository.FindUserRefreshToken(userId); err != nil {
		return nil, err
	}

	res, err := c.orderMysqlRepository.GetUserOrderDetails(orderID, userID)
	if err != nil {
		return nil, err
	}

	if res.UserID != userID {
		return nil, errors.New("")
	}

	return res, nil
}

func (c *orderService) GetWorkshopOrderDetails(orderId, workshopId string) (*entities.Orders, error) {
	orderID, err := strconv.ParseUint(orderId, 10, 64)
	if err != nil {
		return nil, err
	}

	workshopID, err := strconv.ParseUint(workshopId, 10, 64)
	if err != nil {
		return nil, err
	}

	if err := c.orderScribleRepository.FindWorkshopRefreshToken(workshopId); err != nil {
		return nil, err
	}

	res, err := c.orderMysqlRepository.GetWorkshopOrderDetails(orderID, workshopID)
	if err != nil {
		return nil, err
	}

	if res.WorkshopID != workshopID {
		return nil, errors.New("")
	}

	return res, nil
}

func (c *orderService) UserCancelOrder(orderId, userId, username string) error {
	orderID, err := strconv.ParseUint(orderId, 10, 64)
	if err != nil {
		return err
	}

	userID, err := strconv.ParseUint(userId, 10, 64)
	if err != nil {
		return err
	}

	if err := c.orderScribleRepository.FindUserRefreshToken(userId); err != nil {
		return err
	}

	if err := c.orderMysqlRepository.UserCancelOrder(orderID, userID, username); err != nil {
		return err
	}

	return nil
}