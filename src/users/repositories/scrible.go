package repositories

import (
	"gorepair-rest-api/infrastructures/local_db"
	"gorepair-rest-api/src/users/entities"
)

type UserScribleRepositoryInterface interface {
	FindUserRefreshToken(id string) (entities.UserRefreshToken, error)
}

type userScribleRepository struct {
	scribleDB local_db.ScribleDB
}

func NewUserScribleRepositoryInterface(scribleDB local_db.ScribleDB) UserScribleRepositoryInterface {
	return &userScribleRepository{
		scribleDB: scribleDB,
	}
}

func (c userScribleRepository) FindUserRefreshToken(id string) (entities.UserRefreshToken, error) {
	var userRefreshToken entities.UserRefreshToken
	err := c.scribleDB.DB().Read("refresh_token", id, &userRefreshToken)
	if err != nil {
		return entities.UserRefreshToken{}, err
	}
	return userRefreshToken, nil
}
