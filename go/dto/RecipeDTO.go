package dto

import (
	"Status418/go/enums"
	"Status418/go/models"
	"Status418/go/utils"
)

type RecipeDto struct {
	Id          string            `json:"_id"`
	Name        string            `json:"recipe_name" validate:"required,min=3,max=100" required:"recipe name cannot be empty"`
	Ingredients []FoodQuantityDTO `json:"recipe_ingredients" validate:"required" required:"recipe ingredients cannot be empty"`
	Moment      string     		  `json:"recipe_moment" validate:"required" required:"recipe moment cannot be empty"`
	Description string            `json:"recipe_description" validate:"required,max=180" required:"recipe description cannot be empty"`
	UserCode    string            `json:"recipe_usercode" validate:"required" required:"recipe user code cannot be empty"`
}

func NewRecipeDto(model models.Recipe) *RecipeDto {
	var dtoIngredients []FoodQuantityDTO

	for _, food := range model.Ingredients {
		dtoIngredients = append(dtoIngredients, FoodQuantityDTO{FoodCode: utils.GetStringIDFromObjectID(food.FoodCode),Name: food.Name, Quantity: food.Quantity})
	}

	return &RecipeDto{
		Id:          utils.GetStringIDFromObjectID(model.Id),
		Name:        model.Name,
		Ingredients: dtoIngredients,
		Moment:      model.Moment.String(),
		Description: model.Description,
		UserCode:    model.UserCode,
	}
}

func (dto RecipeDto) GetModel() models.Recipe {
	var ingredients []models.FoodQuantity

	for _, food := range dto.Ingredients {
		ingredients = append(ingredients, models.FoodQuantity{
			FoodCode: utils.GetObjectIDFromStringID(food.FoodCode),
			Name: food.Name,
			Quantity: food.Quantity,
		})
	}

	return models.Recipe{
		Id:          utils.GetObjectIDFromStringID(dto.Id),
		Name:        dto.Name,
		Ingredients: ingredients,
		Moment:      enums.GetMomentEnum(dto.Moment),
		Description: dto.Description,
		UserCode:    dto.UserCode,
	}
}
