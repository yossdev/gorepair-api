package repositories

import (
	"gorepair-rest-api/infrastructures/db"
	"gorepair-rest-api/src/users/entities"
)

type userMysqlRepository struct {
	DB db.MysqlDB
}

func NewUserMysqlRepository(DB db.MysqlDB) entities.Repository {
	return &userMysqlRepository{
		DB: DB,
	}
}

func (u *userMysqlRepository) Register(data *entities.Users) (*entities.Users, error) {
	u.DB.DB().Create(&data)
	// if e.Error != nil {
	// 	return nil, e.Error
	// }

	return data, nil
}

func (u *userMysqlRepository) GetUser(param string) (*entities.Users, error) {
	user := entities.Users{}

	if e := u.DB.DB().First(&user, "username = ?", param).Error; e != nil {
		return nil, e
	}
	return &user, nil
}

// func (u *userMysqlRepository) FindAll() []entities.Users {
// 	var users []entities.Users
// 	u.DB.DB().Find(&users)

// 	return users
// }

func (u *userMysqlRepository) FindByID(id uint64) (*entities.Users, error) {
	var user entities.Users
	err := u.DB.DB().First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userMysqlRepository) FindByEmail(email string) *entities.Users {
	var user entities.Users
	u.DB.DB().Where("email = ?", email).First(&user)

	return &user
}
