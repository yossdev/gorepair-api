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

func (u *userMysqlRepository) Register(payload *entities.Users, street string) (*entities.Users, error) {
	user := fromDomain(*payload)
	user.Address = UserAddress{Street: street}
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

func (u *userMysqlRepository) UpdateAccount(payload *entities.Users, id uint64) (*entities.Users, error) {
	user := User{}

	u.DB.DB().First(&user, "id = ?", id)

	fromDomainAccount(payload, &user)
	
	res := u.DB.DB().Save(&user)
	if res.Error != nil {
		return nil, res.Error
	}

	return user.toDomain(), nil
}

func (u *userMysqlRepository) UpdateAddress(payload *entities.UserAddress, id uint64) (*entities.UserAddress, error) {
	address := UserAddress{}

	u.DB.DB().First(&address, "user_id = ?", id)

	fromDomainAddress(payload, &address)

	res := u.DB.DB().Save(&address)
	if res.Error != nil {
		return nil, res.Error
	}

	return address.toDomain(), nil
}

func (u *userMysqlRepository) GetAddress(id uint64) (*entities.UserAddress, error) {
	address := UserAddress{}
	if err := u.DB.DB().First(&address, "user_id = ?", id).Error; err != nil {
		return nil, err
	}

	return address.toDomain(), nil
}