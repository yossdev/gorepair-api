package services

import (
	"database/sql"
	"gorepair-rest-api/internal/utils/auth"
	"gorepair-rest-api/internal/utils/helper"
	"gorepair-rest-api/src/workshops/entities"
	"strconv"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type workshopService struct {
	workshopMysqlRepository   entities.WorkshopMysqlRepositoryInterface
	workshopScribleRepository entities.WorkshopScribleRepositoryInterface
	jwtAuth                   auth.JwtTokenInterface
}

func NewWorkshopService(
	workshopMysqlRepository 	entities.WorkshopMysqlRepositoryInterface,
	workshopScribleRepository 	entities.WorkshopScribleRepositoryInterface,
	jwtAuth 					auth.JwtTokenInterface,
) entities.WorkshopService {
	return &workshopService{
		workshopMysqlRepository:   workshopMysqlRepository,
		workshopScribleRepository: workshopScribleRepository,
		jwtAuth:                   jwtAuth,
	}
}

func (c *workshopService) FindByID(id string) error {
	err := c.workshopScribleRepository.FindWorkshopRefreshToken(id)
	if err != nil {
		return err
	}
	return nil
}

func (c *workshopService) GetWorkshop(username string) (*entities.Workshops, error) {
	workshop, err := c.workshopMysqlRepository.GetWorkshop(username)
	return workshop, err
}

func (c *workshopService) Register(payload *entities.Workshops , street, description string) (*entities.Workshops, error) {
	payload.Password, _ = helper.Hash(payload.Password)
	workshop, err := c.workshopMysqlRepository.Register(payload, street, description)
	return workshop, err
}

func (c *workshopService) Login(payload *entities.Workshops) (interface{}, error) {
	workshop := c.workshopMysqlRepository.FindByEmail(payload.Email)
	if workshop.ID == 0 {
		return nil, sql.ErrNoRows
	}

	err := bcrypt.CompareHashAndPassword([]byte(workshop.Password), []byte(payload.Password))
	if err != nil {
		return nil, err
	}

	token := c.jwtAuth.Sign(jwt.MapClaims{
		"id": workshop.ID,
		"role": "workshop",
	})

	return token, nil
}

func (c *workshopService) Logout(id string) error {
	err := c.workshopScribleRepository.DeleteWorkshopRefreshToken(id)
	if err != nil {
		return err
	}
	return nil
}

func (c *workshopService) UpdateAccount(payload *entities.Workshops, id uint64) (*entities.Workshops, error) {
	payload.Password, _ = helper.Hash(payload.Password)
	workshop, err := c.workshopMysqlRepository.UpdateAccount(payload, id)
	if err != nil {
		return workshop, err
	}
	return workshop, nil
}

func (c *workshopService) UpdateAddress(payload *entities.WorkshopAddress, id uint64) (*entities.WorkshopAddress, error) {
	res, err := c.workshopMysqlRepository.UpdateAddress(payload, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *workshopService) GetAddress(id uint64) (*entities.WorkshopAddress, error)  {
	address, err := c.workshopMysqlRepository.GetAddress(id)
	return address, err
}

func (c *workshopService) UpdateDescription(payload *entities.Descriptions, id uint64) (*entities.Descriptions, error) {
	res, err := c.workshopMysqlRepository.UpdateDescription(payload, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *workshopService) ServicesNew(payload *entities.Services, id uint64) (*entities.Services, error) {
	res, err := c.workshopMysqlRepository.ServicesNew(payload, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *workshopService) UpdateServices(payload *entities.Services, id uint64, servicesId string) (*entities.Services, error) {
	servID, e := strconv.ParseUint(servicesId, 10, 64)
	if e != nil {
		return nil, e
	}

	res, err := c.workshopMysqlRepository.UpdateServices(payload, id, servID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *workshopService) DeleteServices(id uint64, servicesId string) error {
	servID, e := strconv.ParseUint(servicesId, 10, 64)
	if e != nil {
		return e
	}

	err := c.workshopMysqlRepository.DeleteServices(id, servID)
	if err != nil {
		return err
	}
	return nil
}