package repositories

import (
	"gorepair-rest-api/infrastructures/db"
	"gorepair-rest-api/src/users/entities"
)

type userMysqlRepository struct {
	DB db.MysqlDB
}

func NewUserMysqlRepository(DB db.MysqlDB) entities.UserRepository {
	return &userMysqlRepository{
		DB: DB,
	}
}

func (u *userMysqlRepository) GetUser(param string) (*entities.Users, error) {
	user := entities.Users{}
	if err := u.DB.DB().First(&user, "username = ?", param).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userMysqlRepository) Register(payload *entities.Users) (*entities.Users, error) {
	user := fromDomain(*payload)

	tx := u.DB.DB().Begin()
	defer func() {
		if r := recover(); r != nil {
		tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return user.toDomain(), tx.Commit().Error
}

func (u *userMysqlRepository) FindByEmail(email string) *entities.Users {
	user := entities.Users{}
	u.DB.DB().Where("email = ?", email).First(&user)

	return &user
}

// func (u *userMysqlRepository) Account(payload *entities.Users) (*entities.Users, error) {
	
// }

// func (u *userMysqlRepository) Address(payload *entities.Users) (*entities.Users, error) {
	
// }