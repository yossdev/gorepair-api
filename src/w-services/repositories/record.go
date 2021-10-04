package repositories

import (
	"gorepair-rest-api/src/w-services/entities"
	"time"
)

type Service struct {
	ID          uint64
	WorkshopID  uint64
	Vehicle     string
	VehicleType string
	Services    string
	Price       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (rec Service) toDomain() entities.WServices {
	return entities.WServices{
		ID:          rec.ID,
		WorkshopID:  rec.WorkshopID,
		Vehicle:     rec.Vehicle,
		VehicleType: rec.VehicleType,
		Services:    rec.Services,
		Price:       rec.Price,
	}
}

func toDomainSlice(rec []Service) []entities.WServices {
	res := []entities.WServices{}

	for _, val := range rec {
		res = append(res, val.toDomain())
	}
	return res
}