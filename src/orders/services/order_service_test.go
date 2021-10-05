package services

import (
	"errors"
	"gorepair-rest-api/src/orders/entities"
	"gorepair-rest-api/src/orders/entities/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	orderMysqlRepo mocks.OrderMysqlRepositoryInterface
	orderScribleRepo mocks.OrderScribleRepositoryInterface

	orderUsecase entities.OrderService
	orderDomain entities.Orders
)

func setup() {
	orderUsecase = NewOrderService(&orderMysqlRepo, &orderScribleRepo)

	orderDomain = entities.Orders{
		ID: 1,
		UserID: 1,
		Phone: "082212345678",
		Street: "Jl. Jember",
		WorkshopID: 1,
		ServiceID: 2,
		OnSite: false,
		PickUp: true,
		Note: "sesuai alamat",
		TotalPrice: 123456789,
		Placed: true,
	}
}


func TestUserCancelOrder(t *testing.T) {
	setup()

	t.Run("Test UserCancelOrder 1 | Valid", func(t *testing.T) {
		orderMysqlRepo.On("UserCancelOrder",
			mock.AnythingOfType("uint64"),
			mock.AnythingOfType("uint64"),
			mock.AnythingOfType("string")).Return(nil).Once()
		orderScribleRepo.On("FindUserRefreshToken",
			mock.AnythingOfType("string")).Return(nil).Once()

		err := orderUsecase.UserCancelOrder("1", "1", "jojo")

		assert.Nil(t, err)
	})

	t.Run("Test UserCancelOrder 2 | Error", func(t *testing.T) {
		orderMysqlRepo.On("UserCancelOrder",
			mock.AnythingOfType("uint64"),
			mock.AnythingOfType("uint64"),
			mock.AnythingOfType("string")).Return(errors.New("")).Once()
		orderScribleRepo.On("FindUserRefreshToken",
			mock.AnythingOfType("string")).Return(errors.New("")).Once()

		err := orderUsecase.UserCancelOrder("1", "1", "jojo")

		assert.NotNil(t, err)
	})

	t.Run("Test UserCancelOrder 3 | Invalid", func(t *testing.T) {
		orderMysqlRepo.On("UserCancelOrder",
			mock.AnythingOfType("uint64"),
			mock.AnythingOfType("uint64"),
			mock.AnythingOfType("string")).Return(nil).Once()
		orderScribleRepo.On("FindUserRefreshToken",
			mock.AnythingOfType("string")).Return(nil).Once()

		err := orderUsecase.UserCancelOrder("1", "1a", "jojo")

		assert.NotNil(t, err)
	})

	t.Run("Test UserCancelOrder 4 | Invalid", func(t *testing.T) {
		orderMysqlRepo.On("UserCancelOrder",
			mock.AnythingOfType("uint64"),
			mock.AnythingOfType("uint64"),
			mock.AnythingOfType("string")).Return(nil).Once()
		orderScribleRepo.On("FindUserRefreshToken",
			mock.AnythingOfType("string")).Return(nil).Once()

		err := orderUsecase.UserCancelOrder("1a", "1", "jojo")

		assert.NotNil(t, err)
	})
}

func TestOrderNew(t *testing.T) {
	setup()

	t.Run("Test OrderNew 1 | Valid", func(t *testing.T) {
		orderMysqlRepo.On("OrderNew",
			mock.Anything,
			mock.AnythingOfType("uint64")).Return(&orderDomain, nil).Once()
		orderScribleRepo.On("FindUserRefreshToken",
			mock.AnythingOfType("string")).Return(nil).Once()

		_, err := orderUsecase.OrderNew(&entities.Orders{
			WorkshopID: 1,
			ServiceID: 2,
			OnSite: false,
			PickUp: true,
			Note: "sesuai alamat",
		}, "1")

		assert.Nil(t, err)
	})

	t.Run("Test OrderNew 2 | Invalid", func(t *testing.T) {
		orderMysqlRepo.On("OrderNew",
			mock.Anything,
			mock.AnythingOfType("uint64")).Return(&orderDomain, nil).Once()
		orderScribleRepo.On("FindUserRefreshToken",
			mock.AnythingOfType("string")).Return(nil).Once()

		_, err := orderUsecase.OrderNew(&entities.Orders{
			WorkshopID: 1,
			ServiceID: 2,
			OnSite: false,
			PickUp: true,
			Note: "sesuai alamat",
		}, "1a")

		assert.NotNil(t, err)
	})
}

func TestGetUserOrderDetails(t *testing.T) {
	setup()

	t.Run("Test GetUserOrderDetails 1 | Valid", func(t *testing.T) {
		orderMysqlRepo.On("GetUserOrderDetails",
			mock.AnythingOfType("uint64"),
			mock.AnythingOfType("uint64")).Return(&orderDomain, nil).Once()
		orderScribleRepo.On("FindUserRefreshToken",
			mock.AnythingOfType("string")).Return(nil).Once()

		_, err := orderUsecase.GetUserOrderDetails("1", "1")

		assert.Nil(t, err)
	})

	t.Run("Test GetUserOrderDetails 2 | Error", func(t *testing.T) {
		orderMysqlRepo.On("GetUserOrderDetails",
			mock.AnythingOfType("uint64"),
			mock.AnythingOfType("uint64")).Return(&orderDomain, errors.New("")).Once()
		orderScribleRepo.On("FindUserRefreshToken",
			mock.AnythingOfType("string")).Return(errors.New("")).Once()

		_, err := orderUsecase.GetUserOrderDetails("1", "1")

		assert.NotNil(t, err)
	})

	t.Run("Test GetUserOrderDetails 3 | Invalid", func(t *testing.T) {
		orderMysqlRepo.On("GetUserOrderDetails",
			mock.AnythingOfType("uint64"),
			mock.AnythingOfType("uint64")).Return(&orderDomain, nil).Once()
		orderScribleRepo.On("FindUserRefreshToken",
			mock.AnythingOfType("string")).Return(nil).Once()

		_, err := orderUsecase.GetUserOrderDetails("1", "1s")

		assert.NotNil(t, err)
	})

	t.Run("Test GetUserOrderDetails 4 | Invalid", func(t *testing.T) {
		orderMysqlRepo.On("GetUserOrderDetails",
			mock.AnythingOfType("uint64"),
			mock.AnythingOfType("uint64")).Return(&orderDomain, nil).Once()
		orderScribleRepo.On("FindUserRefreshToken",
			mock.AnythingOfType("string")).Return(nil).Once()

		_, err := orderUsecase.GetUserOrderDetails("1a", "1")

		assert.NotNil(t, err)
	})


}

func TestGetWorkshopOrderDetails(t *testing.T) {
	setup()

	t.Run("Test GetWorkshopOrderDetails 1 | Valid", func(t *testing.T) {
		orderMysqlRepo.On("GetWorkshopOrderDetails",
			mock.AnythingOfType("uint64"),
			mock.AnythingOfType("uint64")).Return(&orderDomain, nil).Once()
		orderScribleRepo.On("FindWorkshopRefreshToken",
			mock.AnythingOfType("string")).Return(nil).Once()

		_, err := orderUsecase.GetWorkshopOrderDetails("1", "1")

		assert.Nil(t, err)
	})

	t.Run("Test GetWorkshopOrderDetails 2 | Error", func(t *testing.T) {
		orderMysqlRepo.On("GetWorkshopOrderDetails",
			mock.AnythingOfType("uint64"),
			mock.AnythingOfType("uint64")).Return(&orderDomain, errors.New("")).Once()
		orderScribleRepo.On("FindWorkshopRefreshToken",
			mock.AnythingOfType("string")).Return(errors.New("")).Once()

		_, err := orderUsecase.GetWorkshopOrderDetails("1", "1")

		assert.NotNil(t, err)
	})

	t.Run("Test GetWorkshopOrderDetails 3 | Invalid", func(t *testing.T) {
		orderMysqlRepo.On("GetWorkshopOrderDetails",
			mock.AnythingOfType("uint64"),
			mock.AnythingOfType("uint64")).Return(&orderDomain, nil).Once()
		orderScribleRepo.On("FindWorkshopRefreshToken",
			mock.AnythingOfType("string")).Return(nil).Once()

		_, err := orderUsecase.GetWorkshopOrderDetails("1", "1a")

		assert.NotNil(t, err)
	})

	t.Run("Test GetWorkshopOrderDetails 4 | Invalid", func(t *testing.T) {
		orderMysqlRepo.On("GetWorkshopOrderDetails",
			mock.AnythingOfType("uint64"),
			mock.AnythingOfType("uint64")).Return(&orderDomain, nil).Once()
		orderScribleRepo.On("FindWorkshopRefreshToken",
			mock.AnythingOfType("string")).Return(nil).Once()

		_, err := orderUsecase.GetWorkshopOrderDetails("1a", "1")

		assert.NotNil(t, err)
	})
}
