package repositories

import (
	"gorepair-rest-api/infrastructures/db"
	"gorepair-rest-api/src/w-services/entities"
)

type wservicesMysqlRepository struct {
	DB db.MysqlDB
}

func NewWServicesMysqlRepository(DB db.MysqlDB) entities.WServicesMysqlRepositoryInterface {
	return &wservicesMysqlRepository{
		DB: DB,
	}
}

func (u *wservicesMysqlRepository) GetAll() ([]entities.WServices, error) {
	services := []Service{}

	u.DB.DB().Find(&services)

	return toDomainSlice(services), nil
}

func (u *wservicesMysqlRepository) GetDetails(id uint64) (entities.WServices, error) {
	services := Service{}
	
	res := u.DB.DB().First(&services, "id = ?", id)
	if res.Error != nil {
		return services.toDomain(), res.Error
	}

	return services.toDomain(), nil
}