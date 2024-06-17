package repository

import (
	"sampah/entity"

	"gorm.io/gorm"
)

type CarbonsRepository interface {
	GetDetailCarbons(user_id int) []entity.Carbons
	GetFootPrint(user_id int) float32
	InsertCarbons(carbon entity.Carbons) entity.Carbons
}

type carbonsConnection struct {
	connection *gorm.DB
}

func NewCarbonsRepository(db *gorm.DB) CarbonsRepository {
	return &carbonsConnection{
		connection: db,
	}
}

func (db *carbonsConnection) GetDetailCarbons(user_id int) []entity.Carbons {
	var carbons []entity.Carbons
	db.connection.Where("user_id = ?", user_id).Last(&carbons)
	return carbons
}

func (db *carbonsConnection) GetFootPrint(user_id int) float32 {
	var carbons entity.Carbons
	db.connection.Where("user_id = ?", user_id).Select("carbon_footprint").Last(&carbons)
	return float32(carbons.Carbon_footprint)
}


func (db *carbonsConnection) InsertCarbons(carbon entity.Carbons) entity.Carbons {
	db.connection.Save(&carbon)
	return carbon
}