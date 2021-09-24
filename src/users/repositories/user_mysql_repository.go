package repositories

import (
	"gorepair-rest-api/infrastructures/db"
	"gorepair-rest-api/src/users/entities"
)

type UserMysqlRepositoryInterface interface {
	FindAll() []entities.User
	FindByID(id uint) entities.User
	FindByEmail(email string) entities.User
	Register(uname, name, email, password, phone string) entities.User
}

type userMysqlRepository struct {
	DB db.MysqlDB
}

func NewUserMysqlRepository(DB db.MysqlDB) UserMysqlRepositoryInterface {
	return &userMysqlRepository{
		DB: DB,
	}
}

func (u *userMysqlRepository) Register(uname, name, email, password, phone string) entities.User {
	var user entities.User
	
	user.Username = uname
	user.Name = name
	user.Email = email
	user.Password = password
	user.Phone = phone

	e := u.DB.DB().Create(&user)
	if e.Error != nil {
		return user
	}
	return user
}

func (u *userMysqlRepository) FindAll() []entities.User {
	var users []entities.User
	u.DB.DB().Find(&users)

	return users
}

func (u *userMysqlRepository) FindByID(id uint) entities.User {
	var user entities.User
	u.DB.DB().First(&user, id)

	return user
}

func (u *userMysqlRepository) FindByEmail(email string) entities.User {
	var user entities.User
	u.DB.DB().Where("email = ?", email).First(&user)

	return user
}
