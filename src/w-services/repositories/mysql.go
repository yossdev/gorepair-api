package repositories

import (
	"gorepair-rest-api/infrastructures/db"
	"gorepair-rest-api/src/w-services/entities"
	_ws "gorepair-rest-api/src/workshops/entities"
	_w "gorepair-rest-api/src/workshops/repositories"
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

	u.DB.DB().Where("deleted_at = ?", nil).Find(&services)

	return toDomainSlice(services), nil
}

func (u *wservicesMysqlRepository) GetDetails(id uint64) (entities.WServices, error) {
	services := Service{}
	
	res := u.DB.DB().Where("id = ?", id).Find(&services)
	if res.Error != nil {
		return services.toDomain(), res.Error
	}

	return services.toDomain(), nil
}

func (u *wservicesMysqlRepository) GetAllWorkshop(city string) ([]_ws.WorkshopAddress, error) {
	address := []_w.WorkshopAddress{}
	res := u.DB.DB().Where("city = ?", city).Find(&address)
	if res.Error != nil {
		return nil, res.Error
	}

	return toDomainWS(address), nil
}