package entity

import "time"

type Carbons struct {
	ID        uint64    `json:"id"`
	Electriccity     float32    `json:"electriccity"`
	Gas      float32    `json:"gas"`
	Transportation  float32    `json:"transportation"`
	Food_type     string    `json:"food_type"`
	Food   float32    `json:"food"`
	Organic_waste     float32    `json:"organic_waste"`
	Inorganic_waste     float32    `json:"inorganic_waste"`
	User_id     uint64    `json:"user_id"`
	Carbon_footprint     float32    `json:"carbon_footprint"`
	CreatedAt time.Time `json:"created_At"`
	UpdatedAt time.Time `json:"updated_At"`
}
