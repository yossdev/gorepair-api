package services

import (
	"gorepair-rest-api/src/w-services/entities"
	"strconv"
)

type wservicesService struct {
	wservicesMysqlRepository entities.WServicesMysqlRepositoryInterface
}

func NewWServicesService(
	wservicesMysqlRepository entities.WServicesMysqlRepositoryInterface,
) entities.WServicesService {
	return &wservicesService{
		wservicesMysqlRepository: wservicesMysqlRepository,
	}
}

func (c *wservicesService) GetAll() ([]entities.WServices, error) {
	res, _ := c.wservicesMysqlRepository.GetAll()
	return res, nil
}

func (c *wservicesService) GetDetails(servicesId string) (entities.WServices, error) {
	servID, e := strconv.ParseUint(servicesId, 10, 64)
	if e != nil {
		return entities.WServices{}, e
	}

	res, err := c.wservicesMysqlRepository.GetDetails(servID)
	if err != nil {
		return res, err
	}
	return res, nil
}