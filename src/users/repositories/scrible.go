package repositories

import (
	"gorepair-rest-api/infrastructures/local_db"
	"gorepair-rest-api/src/users/entities"
)

type userScribleRepository struct {
	scribleDB local_db.ScribleDB
}

func NewUserScribleRepositoryInterface(scribleDB local_db.ScribleDB) entities.UserScribleRepositoryInterface {
	return &userScribleRepository{
		scribleDB: scribleDB,
	}
}

func (c userScribleRepository) FindUserRefreshToken(userID string) error {
	err := c.scribleDB.DB().Read("refresh_token_user", userID, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c userScribleRepository) DeleteUserRefreshToken(userID string) error {
	err := c.scribleDB.DB().Delete("refresh_token_user", userID)
	if err != nil {
		return err
	}
	return nil
}