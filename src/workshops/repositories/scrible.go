package repositories

import (
	"gorepair-rest-api/infrastructures/local_db"
	"gorepair-rest-api/src/workshops/entities"
)

type workshopScribleRepository struct {
	scribleDB local_db.ScribleDB
}

func NewWorkshopScribleRepositoryInterface(scribleDB local_db.ScribleDB) entities.WorkshopScribleRepositoryInterface {
	return &workshopScribleRepository{
		scribleDB: scribleDB,
	}
}

func (c workshopScribleRepository) FindWorkshopRefreshToken(workshopID string) error {
	err := c.scribleDB.DB().Read("refresh_token_workshop", workshopID, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c workshopScribleRepository) DeleteWorkshopRefreshToken(workshopID string) error {
	err := c.scribleDB.DB().Delete("refresh_token_workshop", workshopID)
	if err != nil {
		return err
	}
	return nil
}