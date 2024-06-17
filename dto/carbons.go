package dto

type CarbonsCreateDTO struct {
	Electriccity     float32 `json:"electriccity" form:"electriccity"`
	Gas              float32 `json:"gas" form:"gas"`
	Transportation   float32 `json:"transportation" form:"transportation"`
	Food_type        string  `json:"food_type" form:"food_type"`
	Food             float32 `json:"food" form:"food"`
	Organic_waste    float32 `json:"organic_waste" form:"organic_waste"`
	Inorganic_waste  float32 `json:"inorganic_waste" form:"inorganic_waste"`
	Carbon_footprint float32 `json:"carbon_footprint" form:"carbon_footprint"`
	User_id          uint64  `json:"user_id" form:"user_id" binding:"required"`
}