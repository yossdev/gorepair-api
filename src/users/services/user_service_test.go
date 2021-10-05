package services

import (
	"errors"
	"gorepair-rest-api/internal/utils/auth"
	"gorepair-rest-api/src/users/entities"
	"gorepair-rest-api/src/users/entities/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	userMysqlRepository mocks.UserMysqlRepositoryInterface
	userScribleRepository mocks.UserScribleRepositoryInterface
	jwtAuth mocks.JwtTokenInterface

	userUsecase entities.UserService
	userDomain *entities.Users
	userJwt auth.TokenStruct
)

func setup() {
	userUsecase = NewUserService(&userMysqlRepository, &userScribleRepository, &jwtAuth)
	userDomain = &entities.Users{
		ID:       1,
		Username: "jojo123",
		Email:    "jojo@gmail.com",
		Password: "$2a$04$pyOz6LbPAV.DaTqWDWYMAuLhJoUVjCp3J6KCHn5J58Ff/qxLDuBK6",
		Name:     "jojo",
		Gender:   "M",
		DOB:      time.Now(),
		Phone:    "0822",
		UserAddress: entities.UserAddress{
			ID:             1,
			UserID:         1,
			BuildingNumber: "12A",
			Street:         "Jl. Jember",
			City:           "Jember",
			Country:        "Indonesia",
			PostalCode:     "1111",
			Province:       "Jatim",
		},
	}

	userJwt = auth.TokenStruct{
		Type: "Bearer",
		Token: "asdf",
		RefreshToken: "refreshasdf",
	}
}

func TestLogin(t *testing.T) {
	setup()

	userMysqlRepository.On("FindByEmail",
		mock.AnythingOfType("string")).Return(userDomain).Twice()

	jwtAuth.On("Sign", mock.AnythingOfType("MapClaims")).Return(userJwt).Twice()

	t.Run("Test Case 1 | Valid Login", func(t *testing.T) {
		token, err := userUsecase.Login(&entities.Users{
			Email:    "jojo@gmail.com",
			Password: "asdf123",
		})

		assert.Nil(t, err)
		assert.Equal(t, userJwt, token)
	})

	t.Run("Test Case 2 | Invalid Password", func(t *testing.T) {
		_, err := userUsecase.Login(&entities.Users{
			Email:    "jojo@gmail.com",
			Password: "jojo",
		})

		assert.NotNil(t, err)
	})
}

func TestGetUser(t *testing.T) {
	setup()

	t.Run("Test Get User 1 Valid", func(t *testing.T) {
		userScribleRepository.On("FindUserRefreshToken",
			mock.AnythingOfType("string")).Return(nil).Once()
		userMysqlRepository.On("GetUser",
			mock.AnythingOfType("string")).Return(userDomain, nil).Once()

		jwtAuth.On("Sign", mock.AnythingOfType("MapClaims")).Return(userJwt).Once()

		user, err := userUsecase.GetUser("jojo123")

		assert.Nil(t, err)
		assert.Equal(t, "jojo123", user.Username)
	})

	t.Run("Test Get User 2 Error", func(t *testing.T) {
		userScribleRepository.On("FindUserRefreshToken",
			mock.AnythingOfType("string")).Return(errors.New("")).Once()
		userMysqlRepository.On("GetUser",
			mock.AnythingOfType("string")).Return(userDomain, nil).Once()

		jwtAuth.On("Sign", mock.AnythingOfType("MapClaims")).Return(userJwt).Once()
		
		_, err := userUsecase.GetUser("jojo123")

		assert.NotNil(t, err)
	})
}

func TestRegister(t *testing.T) {
	setup()

	t.Run("Test Register 1 Valid", func(t *testing.T) {
		userMysqlRepository.On("Register",
			mock.Anything,
			mock.AnythingOfType("string")).Return(userDomain, nil).Once()

		jwtAuth.On("Sign", mock.AnythingOfType("MapClaims")).Return(userJwt).Once()

		user, err := userUsecase.Register(&entities.Users{
			Username: "jojo123",
			Email:    "jojo@gmail.com",
			Password: "asdf123",
			Name:     "jojo",
			Gender:   "M",
			DOB:      time.Now(),
			Phone:    "0822",
		}, "Jl. Jember")

		assert.Nil(t, err)
		assert.Equal(t, "Jl. Jember", user.UserAddress.Street)
	})

	t.Run("Test Register 2 Error", func(t *testing.T) {
		userMysqlRepository.On("Register",
			mock.Anything,
			mock.AnythingOfType("string")).Return(userDomain, errors.New("")).Once()

		jwtAuth.On("Sign", mock.AnythingOfType("MapClaims")).Return(userJwt).Once()

		_, err := userUsecase.Register(&entities.Users{
			Username: "jojo123",
			Email:    "jojo@gmail.com",
			Password: "asdf123",
			Name:     "jojo",
			Gender:   "M",
			DOB:      time.Now(),
			Phone:    "0822",
		}, "Jl. Jember")

		assert.NotNil(t, err)
	})
}

func TestLogOut(t *testing.T) {
	setup()

	t.Run("Test LogOut 1", func(t *testing.T) {
		userScribleRepository.On("FindUserRefreshToken",
			mock.AnythingOfType("string")).Return(nil).Once()
		userScribleRepository.On("DeleteUserRefreshToken",
			mock.AnythingOfType("string")).Return(nil).Once()
		userMysqlRepository.On("GetUser",
			mock.AnythingOfType("string")).Return(userDomain, nil).Once()

		jwtAuth.On("Sign", mock.AnythingOfType("MapClaims")).Return(userJwt).Once()
	
		err := userUsecase.Logout("1", "jojo123")

		assert.Nil(t, err)
	})

	t.Run("Test LogOut 2", func(t *testing.T) {
		userScribleRepository.On("FindUserRefreshToken",
			mock.AnythingOfType("string")).Return(errors.New("")).Once()
		userScribleRepository.On("DeleteUserRefreshToken",
			mock.AnythingOfType("string")).Return(nil).Once()
		userMysqlRepository.On("GetUser",
			mock.AnythingOfType("string")).Return(userDomain, nil).Once()

		jwtAuth.On("Sign", mock.AnythingOfType("MapClaims")).Return(userJwt).Once()
	
		err := userUsecase.Logout("1", "jojo12")

		assert.NotNil(t, err)
	})
}

func TestUpdateAccount(t *testing.T) {
	setup()

	t.Run("Test UpdateAccount 1", func(t *testing.T) {
		userScribleRepository.On("FindUserRefreshToken",
			mock.AnythingOfType("string")).Return(nil).Once()
		userMysqlRepository.On("GetUser",
			mock.AnythingOfType("string")).Return(userDomain, nil).Once()
		userMysqlRepository.On("UpdateAccount",
			mock.Anything,
			mock.AnythingOfType("uint64")).Return(userDomain, nil).Once()

		jwtAuth.On("Sign", mock.AnythingOfType("MapClaims")).Return(userJwt).Once()

		user, err := userUsecase.UpdateAccount(&entities.Users{
			Username: "jojo123",
			Email:    "jojo@gmail.com",
			Password: "asdf123",
			Name:     "jojo",
			Gender:   "M",
			DOB:      time.Now(),
			Phone:    "0822",
		}, "jojo123")

		assert.Nil(t, err)
		assert.Equal(t, "jojo123", user.Username)
	})

	t.Run("Test UpdateAccount 2", func(t *testing.T) {
		userScribleRepository.On("FindUserRefreshToken",
			mock.AnythingOfType("string")).Return(errors.New("")).Once()
		userMysqlRepository.On("GetUser",
			mock.AnythingOfType("string")).Return(userDomain, nil).Once()
		userMysqlRepository.On("UpdateAccount",
			mock.Anything,
			mock.AnythingOfType("uint64")).Return(userDomain, nil).Once()

		jwtAuth.On("Sign", mock.AnythingOfType("MapClaims")).Return(userJwt).Once()

		_, err := userUsecase.UpdateAccount(&entities.Users{
			Username: "jojo123",
			Email:    "jojo@gmail.com",
			Password: "asdf123",
			Name:     "jojo",
			Gender:   "M",
			DOB:      time.Now(),
			Phone:    "0822",
		}, "jojo123")

		assert.NotNil(t, err)
	})
}

func TestUpdateAddress(t *testing.T) {
	setup()

	t.Run("Test UpdateAddress 1 Valid", func(t *testing.T) {
		userScribleRepository.On("FindUserRefreshToken",
			mock.AnythingOfType("string")).Return(nil).Once()
		userMysqlRepository.On("GetUser",
			mock.AnythingOfType("string")).Return(userDomain, nil).Once()
		userMysqlRepository.On("UpdateAddress",
			mock.Anything,
			mock.AnythingOfType("uint64")).Return(&userDomain.UserAddress, nil).Once()

		jwtAuth.On("Sign", mock.AnythingOfType("MapClaims")).Return(userJwt).Once()

		address, err := userUsecase.UpdateAddress(&entities.UserAddress{
			BuildingNumber: "12A",
			Street:         "Jl. Jember",
			City:           "Jember",
			Country:        "Indonesia",
			PostalCode:     "1111",
			Province:       "Jatim",
		}, "jojo123")

		assert.Nil(t, err)
		assert.Equal(t, "Jember", address.City)
		assert.NotEqual(t, "Jakarta", address.City)
	})

	t.Run("Test UpdateAddress 2 Error", func(t *testing.T) {
		userScribleRepository.On("FindUserRefreshToken",
			mock.AnythingOfType("string")).Return(errors.New("")).Once()
		userMysqlRepository.On("GetUser",
			mock.AnythingOfType("string")).Return(userDomain, nil).Once()
		userMysqlRepository.On("UpdateAddress",
			mock.Anything,
			mock.AnythingOfType("uint64")).Return(&userDomain.UserAddress, nil).Once()

		jwtAuth.On("Sign", mock.AnythingOfType("MapClaims")).Return(userJwt).Once()

		_, err := userUsecase.UpdateAddress(&entities.UserAddress{
			BuildingNumber: "12A",
			Street:         "Jl. Jember",
			City:           "Jember",
			Country:        "Indonesia",
			PostalCode:     "1111",
			Province:       "Jatim",
		}, "jojo123")

		assert.NotNil(t, err)
	})
}

func TestGetAddress(t *testing.T) {
	setup()

	t.Run("Test GetAddress 1", func(t *testing.T) {
		userScribleRepository.On("FindUserRefreshToken",
			mock.AnythingOfType("string")).Return(nil).Once()
		userMysqlRepository.On("GetUser",
			mock.AnythingOfType("string")).Return(userDomain, nil).Once()
		userMysqlRepository.On("GetAddress",
			mock.AnythingOfType("uint64")).Return(&userDomain.UserAddress, nil).Once()

		jwtAuth.On("Sign", mock.AnythingOfType("MapClaims")).Return(userJwt).Once()

		_, err := userUsecase.GetAddress("jojo123")

		assert.Nil(t, err)
	})

	t.Run("Test GetAddress 2", func(t *testing.T) {
		userScribleRepository.On("FindUserRefreshToken",
			mock.AnythingOfType("string")).Return(errors.New("")).Once()
		userMysqlRepository.On("GetUser",
			mock.AnythingOfType("string")).Return(userDomain, nil).Once()
		userMysqlRepository.On("GetAddress",
			mock.AnythingOfType("uint64")).Return(&userDomain.UserAddress, nil).Once()

		jwtAuth.On("Sign", mock.AnythingOfType("MapClaims")).Return(userJwt).Once()

		_, err := userUsecase.GetAddress("jojo123")

		assert.NotNil(t, err)
	})

	t.Run("Test GetAddress 3", func(t *testing.T) {
		userScribleRepository.On("FindUserRefreshToken",
			mock.AnythingOfType("string")).Return(errors.New("")).Once()
		userMysqlRepository.On("GetUser",
			mock.AnythingOfType("string")).Return(nil, errors.New("")).Once()
		userMysqlRepository.On("GetAddress",
			mock.AnythingOfType("uint64")).Return(&userDomain.UserAddress, nil).Once()

		jwtAuth.On("Sign", mock.AnythingOfType("MapClaims")).Return(userJwt).Once()

		_, err := userUsecase.GetAddress("jojo123")

		assert.NotNil(t, err)
	})
}