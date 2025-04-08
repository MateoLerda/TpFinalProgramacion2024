package dto

import (
	"Status418/go/enums"
	"Status418/go/models"
	"Status418/go/utils"
)

type FoodDto struct {
	Code            string   `json:"_id"`
	Type            string   `json:"type" validate:"required" required:"food type cannot be empty"`
	Moments         []string `json:"moments" validate:"required" required:"food moments cannot be empty"`
	Name            string   `json:"name" validate:"required,min=3,max=50" required:"food name cannot be empty"`
	UnitPrice       float64  `json:"unit_price" validate:"required,gte=0.0" required:"food unitprice cannot be empty"`
	CurrentQuantity int      `json:"current_quantity" validate:"required,gte=0,numeric" required:"food currentquantity cannot be empty"`
	MinimumQuantity int      `json:"minimum_quantity" validate:"required,gte=0,numeric" required:"food minimumquantity cannot be empty"`
}

func NewFoodDto(model models.Food) *FoodDto {

	return &FoodDto{
		Code:            utils.GetStringIDFromObjectID(model.Code),
		Type:            model.Type.String(),
		Moments:         enums.ArrayString(model.Moments),
		Name:            model.Name,
		UnitPrice:       model.UnitPrice,
		CurrentQuantity: model.CurrentQuantity,
		MinimumQuantity: model.MinimumQuantity,
	}
}

func (dto FoodDto) GetModel() models.Food {
	return models.Food{
		Code:            utils.GetObjectIDFromStringID(dto.Code),
		Type:            enums.GetTypeEnum(dto.Type),
		Moments:         enums.GetArrayMoments(dto.Moments),
		Name:            dto.Name,
		UnitPrice:       dto.UnitPrice,
		CurrentQuantity: dto.CurrentQuantity,
		MinimumQuantity: dto.MinimumQuantity,
	}
}
