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
	user := User{}
	if err := u.DB.DB().First(&user, "username = ?", param).Error; err != nil {
		return nil, err
	}

	return user.toDomain(), nil
}

func (u *userMysqlRepository) Register(payload *entities.Users) (*entities.Users, error) {
	user := fromDomain(*payload)
	e := u.DB.DB().Create(&user)
	if e.Error != nil {
		return nil, e.Error
	}

	return user.toDomain(), nil
}

func (u *userMysqlRepository) FindByEmail(email string) *entities.Users {
	user := User{}
	u.DB.DB().Where("email = ?", email).First(&user)

	return user.toDomain()
}

func (u *userMysqlRepository) UpdateAccount(payload *entities.Users, id string) (*entities.Users, error) {
	res := u.DB.DB().Save(*payload)
	if res.Error != nil {
		return nil, res.Error
	}

	return payload, nil
}