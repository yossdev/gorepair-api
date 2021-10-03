package repositories

import (
	"gorepair-rest-api/infrastructures/db"
	"gorepair-rest-api/src/workshops/entities"
)

type workshopMysqlRepository struct {
	DB db.MysqlDB
}

func NewWorkshopMysqlRepository(DB db.MysqlDB) entities.WorkshopMysqlRepositoryInterface {
	return &workshopMysqlRepository{
		DB: DB,
	}
}

func (u *workshopMysqlRepository) GetWorkshop(param string) (*entities.Workshops, error) {
	workshop := Workshop{}
	if err := u.DB.DB().Preload("Description").First(&workshop, "username = ?", param).Error; err != nil {
		return nil, err
	}

	return workshop.toDomain(), nil
}

func (u *workshopMysqlRepository) Register(payload *entities.Workshops, street, description string) (*entities.Workshops, error) {
	workshop := fromDomain(*payload)
	workshop.Address = WorkshopAddress{Street: street}
	workshop.Description = Description{Description: description}
	e := u.DB.DB().Create(&workshop)
	if e.Error != nil {
		return nil, e.Error
	}

	return workshop.toDomain(), nil
}

func (u *workshopMysqlRepository) FindByEmail(email string) *entities.Workshops {
	workshop := Workshop{}
	u.DB.DB().Where("email = ?", email).First(&workshop)

	return workshop.toDomain()
}

func (u *workshopMysqlRepository) UpdateAccount(payload *entities.Workshops, id uint64) (*entities.Workshops, error) {
	workshop := Workshop{}

	u.DB.DB().First(&workshop, "id = ?", id)

	fromDomainAccount(payload, &workshop)

	res := u.DB.DB().Save(&workshop)
	if res.Error != nil {
		return nil, res.Error
	}

	return workshop.toDomain(), nil
}

func (u *workshopMysqlRepository) UpdateAddress(payload *entities.WorkshopAddress, id uint64) (*entities.WorkshopAddress, error) {
	address := WorkshopAddress{}

	u.DB.DB().First(&address, "workshop_id = ?", id)

	fromDomainAddress(payload, &address)

	res := u.DB.DB().Save(&address)
	if res.Error != nil {
		return nil, res.Error
	}

	return address.toDomain(), nil
}

func (u *workshopMysqlRepository) GetAddress(id uint64) (*entities.WorkshopAddress, error) {
	address := WorkshopAddress{}
	if err := u.DB.DB().First(&address, "workshop_id = ?", id).Error; err != nil {
		return nil, err
	}

	return address.toDomain(), nil
}