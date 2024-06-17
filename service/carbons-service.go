package service

import (
	"sampah/dto"
	"sampah/entity"
	"sampah/repository"
)

type CarbonsService interface {
	GetDetailCarbons(user_id int) []entity.Carbons
	GetFootPrint(user_id int) float32
	InsertCarbons(carb dto.CarbonsCreateDTO) entity.Carbons
}

type carbonsService struct {
	carbonsRepository repository.CarbonsRepository
}

func NewCarbonsService(carbRepo repository.CarbonsRepository) CarbonsService {
	return &carbonsService{
		carbonsRepository: carbRepo,
	}
}

func (serv *carbonsService) GetDetailCarbons(user_id int) []entity.Carbons {
	return serv.carbonsRepository.GetDetailCarbons(user_id)
}
func (serv *carbonsService) GetFootPrint(user_id int) float32 {
	return serv.carbonsRepository.GetFootPrint(user_id)
}

func (serv *carbonsService) InsertCarbons(carb dto.CarbonsCreateDTO) entity.Carbons {
	carbons := entity.Carbons{
		Electriccity: carb.Electriccity,
		Gas: carb.Gas,
		Transportation: carb.Transportation,
		Food_type: carb.Food_type,       
		Food: carb.Food,
		Organic_waste: carb.Organic_waste,   
		Inorganic_waste: carb.Inorganic_waste,
		Carbon_footprint: carb.Carbon_footprint, 
		User_id: carb.User_id,
	}

	res := serv.carbonsRepository.InsertCarbons(carbons)
	return res
}



