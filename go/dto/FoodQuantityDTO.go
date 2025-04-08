package dto

type FoodQuantityDTO struct {
	FoodCode string `json:"_id"`
	Name string		`bson:"name"`
	Quantity int    `json:"quantity" validate:"numeric"`
}
