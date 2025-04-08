package dto

type PurchaseDto struct {
	TotalCost float64           `json:"total_cost,omitempty"`
	Foods     []FoodQuantityDTO `json:"foods,omitempty"`
}

