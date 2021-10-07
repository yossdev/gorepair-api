package services

import (
	ipgeo "gorepair-rest-api/src/ip-geo"
	"gorepair-rest-api/src/w-services/entities"
	_ws "gorepair-rest-api/src/workshops/entities"
	"strconv"
)

type wservicesService struct {
	wservicesMysqlRepository entities.WServicesMysqlRepositoryInterface
	ipgeo ipgeo.Repository
}

func NewWServicesService(
	wservicesMysqlRepository entities.WServicesMysqlRepositoryInterface,
	ipgeo ipgeo.Repository,
) entities.WServicesService {
	return &wservicesService{
		wservicesMysqlRepository: wservicesMysqlRepository,
		ipgeo: ipgeo,
	}
}

func (c *wservicesService) GetAll() ([]entities.WServices, error) {
	res, err := c.wservicesMysqlRepository.GetAll()
	if err != nil {
		return []entities.WServices{}, err
	}
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

func (c *wservicesService) GetAllWorkshop(ip string) ([]_ws.WorkshopAddress, error) {
	ipgeo, err := c.ipgeo.GetLocationByIP(ip)
	if err != nil {
		return nil, err
	}

	res, err := c.wservicesMysqlRepository.GetAllWorkshop(ipgeo.City)
	if err != nil {
		return nil, err
	}

	return res, nil
}