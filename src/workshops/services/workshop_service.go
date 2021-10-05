package services

import (
	"database/sql"
	"errors"
	"fmt"
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

func (c *workshopService) GetWorkshop(username string) (*entities.Workshops, error) {
	workshop, err := c.workshopMysqlRepository.GetWorkshop(username)
	if err != nil {
		return nil, err
	}

	return workshop, nil
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

func (c *workshopService) Logout(ctxId, username string) error {
	if err := c.workshopScribleRepository.FindWorkshopRefreshToken(ctxId); err != nil {
		return err
	}

	workshop, err := c.workshopMysqlRepository.GetWorkshop(username)
	if err != nil {
		return err
	}

	id := fmt.Sprintf("%d", workshop.ID)
	
	if id != ctxId {
		return errors.New("")
	}

	if err := c.workshopScribleRepository.DeleteWorkshopRefreshToken(id); err != nil {
		return err
	}
	return nil
}

func (c *workshopService) UpdateAccount(payload *entities.Workshops, username string) (*entities.Workshops, error) {
	w, e := c.workshopMysqlRepository.GetWorkshop(username)
	if e != nil {
		return nil, e
	}

	id := fmt.Sprintf("%d", w.ID)

	if err := c.workshopScribleRepository.FindWorkshopRefreshToken(id); err != nil {
		return nil, err
	}

	payload.Password, _ = helper.Hash(payload.Password)
	workshop, err := c.workshopMysqlRepository.UpdateAccount(payload, w.ID)
	if err != nil {
		return workshop, err
	}
	return workshop, nil
}

func (c *workshopService) UpdateAddress(payload *entities.WorkshopAddress, username string) (*entities.WorkshopAddress, error) {
	w, e := c.workshopMysqlRepository.GetWorkshop(username)
	if e != nil {
		return nil, e
	}

	id := fmt.Sprintf("%d", w.ID)
	
	if err := c.workshopScribleRepository.FindWorkshopRefreshToken(id); err != nil {
		return nil, err
	}

	res, err := c.workshopMysqlRepository.UpdateAddress(payload, w.ID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *workshopService) GetAddress(username string) (*entities.WorkshopAddress, error)  {
	w, e := c.workshopMysqlRepository.GetWorkshop(username)
	if e != nil {
		return nil, e
	}

	id := fmt.Sprintf("%d", w.ID)
	
	if err := c.workshopScribleRepository.FindWorkshopRefreshToken(id); err != nil {
		return nil, err
	}

	address, err := c.workshopMysqlRepository.GetAddress(w.ID)
	return address, err
}

func (c *workshopService) UpdateDescription(payload *entities.Descriptions, username string) (*entities.Descriptions, error) {
	w, e := c.workshopMysqlRepository.GetWorkshop(username)
	if e != nil {
		return nil, e
	}

	id := fmt.Sprintf("%d", w.ID)
	
	if err := c.workshopScribleRepository.FindWorkshopRefreshToken(id); err != nil {
		return nil, err
	}

	res, err := c.workshopMysqlRepository.UpdateDescription(payload, w.ID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *workshopService) ServicesNew(payload *entities.Services, username string) (*entities.Services, error) {
	w, e := c.workshopMysqlRepository.GetWorkshop(username)
	if e != nil {
		return nil, e
	}

	id := fmt.Sprintf("%d", w.ID)
	
	if err := c.workshopScribleRepository.FindWorkshopRefreshToken(id); err != nil {
		return nil, err
	}

	res, err := c.workshopMysqlRepository.ServicesNew(payload, w.ID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *workshopService) UpdateServices(payload *entities.Services, username, servicesId string) (*entities.Services, error) {
	servID, e := strconv.ParseUint(servicesId, 10, 64)
	if e != nil {
		return nil, e
	}

	w, e := c.workshopMysqlRepository.GetWorkshop(username)
	if e != nil {
		return nil, e
	}

	id := fmt.Sprintf("%d", w.ID)
	
	if err := c.workshopScribleRepository.FindWorkshopRefreshToken(id); err != nil {
		return nil, err
	}

	res, err := c.workshopMysqlRepository.UpdateServices(payload, w.ID, servID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *workshopService) DeleteServices(username, servicesId string) error {
	servID, e := strconv.ParseUint(servicesId, 10, 64)
	if e != nil {
		return e
	}

	w, e := c.workshopMysqlRepository.GetWorkshop(username)
	if e != nil {
		return e
	}

	id := fmt.Sprintf("%d", w.ID)
	
	if err := c.workshopScribleRepository.FindWorkshopRefreshToken(id); err != nil {
		return err
	}

	err := c.workshopMysqlRepository.DeleteServices(w.ID, servID)
	if err != nil {
		return err
	}
	return nil
}