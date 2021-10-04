package repositories

import (
	"gorepair-rest-api/infrastructures/local_db"
	"gorepair-rest-api/src/orders/entities"
)

type orderScribleRepository struct {
	scribleDB local_db.ScribleDB
}

func NewOrderScribleRepositoryInterface(scribleDB local_db.ScribleDB) entities.OrderScribleRepositoryInterface {
	return &orderScribleRepository{
		scribleDB: scribleDB,
	}
}

func (c orderScribleRepository) FindWorkshopRefreshToken(id string) error {
	err := c.scribleDB.DB().Read("refresh_token_workshop", id, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c orderScribleRepository) FindUserRefreshToken(id string) error {
	err := c.scribleDB.DB().Read("refresh_token_user", id, nil)
	if err != nil {
		return err
	}
	return nil
}