package services

import (
	"gorepair-rest-api/internal/utils/auth"
	"gorepair-rest-api/src/workshops/entities"
	"gorepair-rest-api/src/workshops/entities/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var workshopMysqlRepository mocks.WorkshopMysqlRepositoryInterface
var workshopScribleRepository mocks.WorkshopScribleRepositoryInterface
var jwtAuth mocks.JwtTokenInterface

var workshopUsecase entities.WorkshopService
var workshopDomain *entities.Workshops
var workshopJwt auth.TokenStruct

func setup() {
	workshopUsecase = NewWorkshopService(&workshopMysqlRepository, &workshopScribleRepository, &jwtAuth)
	workshopDomain = &entities.Workshops{
		ID: 				1,
		Username: 			"zc",
		Email: 				"zc@gmail.com",
		Password: 			"$2a$04$pyOz6LbPAV.DaTqWDWYMAuLhJoUVjCp3J6KCHn5J58Ff/qxLDuBK6",
		Name: 				"zc",
		Phone: 				"082212341234",
		OperationalStart: 	"Monday",
		OperationalEnd: 	"Friday",
		Description: entities.Descriptions{
			ID: 1,
			WorkshopID: 1,
			Description: "ZC is the best workshop in the world!",
		},
		WorkshopAddress: entities.WorkshopAddress{
			ID:             1,
			WorkshopID:     1,
			BuildingNumber: "12A",
			Street:         "Jl. Jember",
			City:           "Jember",
			Country:        "Indonesia",
			PostalCode:     "1111",
			Province:       "Jatim",
		},
		Services: entities.Services{
			ID: 1,
			WorkshopID: 1,
			Vehicle: "BMW 770",
			VehicleType: "Sport Car",
			Services: "Body repair, Engine repair, and Performance booster",
			Price: 100000000,
		},
	}

	workshopJwt = auth.TokenStruct{
		Type:         "Bearer",
		Token:        "asdf",
		RefreshToken: "refreshasdf",
	}
}

func TestLogin(t *testing.T) {
	setup()

	workshopMysqlRepository.On("FindByEmail",
		mock.AnythingOfType("string")).Return(workshopDomain).Twice()

	jwtAuth.On("Sign", mock.AnythingOfType("MapClaims")).Return(workshopJwt).Twice()

	t.Run("Test Case 1 | Valid Login", func(t *testing.T) {
		token, err := workshopUsecase.Login(&entities.Workshops{
			Email:    "zc@gmail.com",
			Password: "asdf123",
		})

		assert.Nil(t, err)
		assert.Equal(t, workshopJwt, token)
	})

	t.Run("Test Case 2 | Invalid Password", func(t *testing.T) {
		_, err := workshopUsecase.Login(&entities.Workshops{
			Email:    "zc@gmail.com",
			Password: "jojo",
		})

		assert.NotNil(t, err)
	})
}

func TestGetWorkshop(t *testing.T) {
	setup()

	workshopMysqlRepository.On("GetWorkshop",
		mock.AnythingOfType("string")).Return(workshopDomain, nil).Once()

	jwtAuth.On("Sign", mock.AnythingOfType("MapClaims")).Return(workshopJwt).Once()

	t.Run("Test Get Workshop", func(t *testing.T) {
		workshop, err := workshopUsecase.GetWorkshop("zc")

		assert.Nil(t, err)
		assert.Equal(t, "zc", workshop.Username)
	})
}

func TestRegister(t *testing.T) {
	setup()

	workshopMysqlRepository.On("Register",
		mock.Anything,
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string")).Return(workshopDomain, nil).Once()

	jwtAuth.On("Sign", mock.AnythingOfType("MapClaims")).Return(workshopJwt).Once()

	t.Run("Test Register", func(t *testing.T) {
		workshop, err := workshopUsecase.Register(&entities.Workshops{
			Username: "zc",
			Email:    "zc@gmail.com",
			Password: "asdf123",
			Name:     "zc",
			Phone:    "0822",
			OperationalStart: "Monday",
			OperationalEnd: "Friday",
		}, "Jl. Jember", "ZC is the best workshop in the world!")

		assert.Nil(t, err)
		assert.Equal(t, "Jl. Jember", workshop.WorkshopAddress.Street)
	})
}

func TestFindByID(t *testing.T) {
	setup()

	workshopScribleRepository.On("FindWorkshopRefreshToken",
		mock.AnythingOfType("string")).Return(nil).Twice()

	jwtAuth.On("Sign", mock.AnythingOfType("MapClaims")).Return(workshopJwt).Twice()

	t.Run("Test FindByID", func(t *testing.T) {
		err := workshopUsecase.FindByID("1")

		assert.Nil(t, err)
	})
}

func TestLogOut(t *testing.T) {
	setup()

	workshopScribleRepository.On("DeleteWorkshopRefreshToken",
		mock.AnythingOfType("string")).Return(nil).Once()

	jwtAuth.On("Sign", mock.AnythingOfType("MapClaims")).Return(workshopJwt).Once()

	t.Run("Test LogOut", func(t *testing.T) {
		err := workshopUsecase.Logout("1")

		assert.Nil(t, err)
	})
}

func TestUpdateAccount(t *testing.T) {
	setup()

	workshopMysqlRepository.On("UpdateAccount",
		mock.Anything,
		mock.AnythingOfType("uint64")).Return(workshopDomain, nil).Once()

	jwtAuth.On("Sign", mock.AnythingOfType("MapClaims")).Return(workshopJwt).Once()

	t.Run("Test UpdateAccount", func(t *testing.T) {
		workshop, err := workshopUsecase.UpdateAccount(&entities.Workshops{
			Username: "zc",
			Email:    "zc@gmail.com",
			Password: "asdf123",
			Name:     "zc",
			Phone:    "0822",
			OperationalStart: "Monday",
			OperationalEnd: "Friday",
		}, 1)

		assert.Nil(t, err)
		assert.Equal(t, uint64(1), workshop.ID)
	})
}

func TestUpdateAddress(t *testing.T) {
	setup()

	workshopMysqlRepository.On("UpdateAddress",
		mock.Anything,
		mock.AnythingOfType("uint64")).Return(&workshopDomain.WorkshopAddress, nil).Once()

	jwtAuth.On("Sign", mock.AnythingOfType("MapClaims")).Return(workshopJwt).Once()

	t.Run("Test UpdateAddress", func(t *testing.T) {
		address, err := workshopUsecase.UpdateAddress(&entities.WorkshopAddress{
			BuildingNumber: "12A",
			Street:         "Jl. Jember",
			City:           "Jember",
			Country:        "Indonesia",
			PostalCode:     "1111",
			Province:       "Jatim",
		}, 1)

		assert.Nil(t, err)
		assert.Equal(t, "Jember", address.City)
		assert.NotEqual(t, "Jakarta", address.City)
	})
}

func TestGetAddress(t *testing.T) {
	setup()

	workshopMysqlRepository.On("GetAddress",
		mock.AnythingOfType("uint64")).Return(&workshopDomain.WorkshopAddress, nil).Once()

	jwtAuth.On("Sign", mock.AnythingOfType("MapClaims")).Return(workshopJwt).Once()

	t.Run("Test GetAddress", func(t *testing.T) {
		_, err := workshopUsecase.GetAddress(1)

		assert.Nil(t, err)
	})
}

func TestUpdateDescription(t *testing.T) {
	setup()

	workshopMysqlRepository.On("UpdateDescription",
	mock.Anything,
	mock.AnythingOfType("uint64")).Return(&workshopDomain.Description, nil).Once()

	jwtAuth.On("Sign", mock.AnythingOfType("MapClaims")).Return(workshopJwt).Once()

	t.Run("Test UpdateDescription", func(t *testing.T) {
		_, err := workshopUsecase.UpdateDescription(&entities.Descriptions{
			ID: 1,
			WorkshopID: 1,
			Description: "ZC is the best workshop in the world!",
		}, 1)

		assert.Nil(t, err)
	})
}

func TestServicesNew(t *testing.T) {
	setup()

	workshopMysqlRepository.On("ServicesNew",
	mock.Anything,
	mock.AnythingOfType("uint64")).Return(&workshopDomain.Services, nil).Once()

	jwtAuth.On("Sign", mock.AnythingOfType("MapClaims")).Return(workshopJwt).Once()

	t.Run("Test ServicesNew", func(t *testing.T) {
		_, err := workshopUsecase.ServicesNew(&entities.Services{
			ID: 1,
			WorkshopID: 1,
			Vehicle: "BMW 770",
			VehicleType: "Sport Car",
			Services: "Body repair, Engine repair, and Performance booster",
			Price: 100000000,
		}, 1)

		assert.Nil(t, err)
	})
}

func TestUpdateServices(t *testing.T) {
	setup()

	workshopMysqlRepository.On("UpdateServices",
	mock.Anything,
	mock.AnythingOfType("uint64"),
	mock.AnythingOfType("uint64")).Return(&workshopDomain.Services, nil).Once()

	jwtAuth.On("Sign", mock.AnythingOfType("MapClaims")).Return(workshopJwt).Once()

	t.Run("Test UpdateServices 1", func(t *testing.T) {
		_, err := workshopUsecase.UpdateServices(&entities.Services{
			ID: 1,
			WorkshopID: 1,
			Vehicle: "BMW 770",
			VehicleType: "Sport Car",
			Services: "Body repair, Engine repair, and Performance booster",
			Price: 100000000,
		}, 1, "1")

		assert.Nil(t, err)
	})

	t.Run("Test UpdateServices 2", func(t *testing.T) {
		_, err := workshopUsecase.UpdateServices(&entities.Services{
			ID: 1,
			WorkshopID: 1,
			Vehicle: "BMW 770",
			VehicleType: "Sport Car",
			Services: "Body repair, Engine repair, and Performance booster",
			Price: 100000000,
		}, 1, "1a")

		assert.NotNil(t, err)
	})
}

func TestDeleteServices(t *testing.T) {
	setup()

	workshopMysqlRepository.On("DeleteServices",
	mock.AnythingOfType("uint64"),
	mock.AnythingOfType("uint64")).Return(nil).Once()

	jwtAuth.On("Sign", mock.AnythingOfType("MapClaims")).Return(workshopJwt).Once()

	t.Run("Test DeleteServices 1", func(t *testing.T) {
		err := workshopUsecase.DeleteServices(1, "1")

		assert.Nil(t, err)
	})

	t.Run("Test DeleteServices 2", func(t *testing.T) {
		err := workshopUsecase.DeleteServices(1, "a")

		assert.NotNil(t, err)
	})
}