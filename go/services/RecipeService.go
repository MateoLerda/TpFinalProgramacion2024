package services

import (
	"Status418/go/dto"
	"Status418/go/enums"
	"Status418/go/models"
	"Status418/go/repositories"
	"Status418/go/utils"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RecipeServiceInterface interface {
	GetAll(userCode string, filters dto.FiltersDto) (*[]dto.RecipeDto, error)
	Create(newRecipe dto.RecipeDto) (*mongo.InsertOneResult, error)
	Delete(recipeId string, userCode string) (*mongo.DeleteResult, error)
	Update(updateRecipe dto.RecipeDto) (*mongo.UpdateResult, error)
	Cook(userCode string, recipeId primitive.ObjectID, cancel bool) (bool, error)
}

type RecipeService struct {
	recipeRepository repositories.RecipeRepositoryInterface
	foodRepository   repositories.FoodRepositoryInterface
}

func NewRecipeService(recipeRepository repositories.RecipeRepositoryInterface, foodRepository repositories.FoodRepositoryInterface) *RecipeService {
	return &RecipeService{
		recipeRepository: recipeRepository,
		foodRepository:   foodRepository,
	}
}

func (recipeService *RecipeService) filterByType(recipes []models.Recipe, userCode string, fType enums.FoodType) ([]models.Recipe, error) {
	var filteredRecipes []models.Recipe
	for _, recipe := range recipes {
		for _, foodQuantity := range recipe.Ingredients {
			food, err := recipeService.foodRepository.GetByCode(foodQuantity.FoodCode, userCode)
			if err != nil {
				return nil, err
			}
			if food.Type == fType {
				filteredRecipes = append(filteredRecipes, recipe)
				break
			}
		}
	}
	return filteredRecipes, nil
}

func (recipeService *RecipeService) validateQuantity(recipe models.Recipe) bool {
	for i, foodQuantity := range recipe.Ingredients {
		food, err := recipeService.foodRepository.GetByCode(foodQuantity.FoodCode, recipe.UserCode)
		if err != nil {
			return false
		}
		if food.CurrentQuantity < foodQuantity.Quantity {
			break
		}
		if i+1 == len(recipe.Ingredients) {
			return true
		}
	}
	return false
}

func (recipeService *RecipeService) filterByQuantity(recipes []models.Recipe) ([]models.Recipe, error) {
	var filteredRecipes []models.Recipe
	for _, recipe := range recipes {
		if recipeService.validateQuantity(recipe) {
			filteredRecipes = append(filteredRecipes, recipe)
		}
	}
	return filteredRecipes, nil
}

func (recipeService *RecipeService) GetAll(userCode string, filters dto.FiltersDto) (*[]dto.RecipeDto, error) {
	var recipesDto []dto.RecipeDto
	recipes, err := recipeService.recipeRepository.GetAll(userCode, filters.GetModel())
	if err != nil {
		return nil, err
	}
	
	if !filters.All {
		recipes, err = recipeService.filterByQuantity(recipes)
		if err != nil {
			return nil, errors.New("nocontent")
		}
	}
	if filters.Type != "" {
		recipes, err = recipeService.filterByType(recipes, userCode, (filters.GetModel()).Type)
		if err != nil {
			return nil, errors.New("internal")
		}
	}
	for _, recipe := range recipes {
		recipeDto := dto.NewRecipeDto(recipe)
		recipesDto = append(recipesDto, *recipeDto)
	}

	return &recipesDto, nil
}

func (recipeService *RecipeService) Create(newRecipe dto.RecipeDto) (*mongo.InsertOneResult, error) {
	recipe := newRecipe.GetModel()
	recipe.CreationDate = time.Now().String()
	validation, err := recipeService.validateMoment(recipe.Ingredients, recipe.Moment, recipe.UserCode)
	if err != nil {
		return nil, err
	}
	if !validation {
		return nil, errors.New("the food moment doesn´t match with the recipe moment")
	}
	if !recipeService.validateQuantity(recipe) {
		return nil, errors.New("the foods are not enough for the recipe")
	}
	res, err := recipeService.recipeRepository.Create(recipe)
	if err != nil {
		return nil, err
	}
	objectId := res.InsertedID.(primitive.ObjectID)
	resultado, err := recipeService.Cook(newRecipe.UserCode, objectId, false)
	if !resultado {
		return nil, err
	}
	return res, nil
}

func (recipeService RecipeService) validateMoment(ingredients []models.FoodQuantity, recipeMoment enums.Moment, userCode string) (bool, error) {
	for _, ingredient := range ingredients {
		momentValidation := false
		food, err := recipeService.foodRepository.GetByCode(ingredient.FoodCode, userCode)
		if err != nil {
			return false, err
		}
		for _, moment := range food.Moments {
			if moment == recipeMoment {
				momentValidation = true
				break
			}
		}
		if !momentValidation {
			return false, nil
		}
	}
	return true, nil
}

func (recipeService *RecipeService) Delete(recipeId string, userCode string) (*mongo.DeleteResult, error) {
	recipeObjectId := utils.GetObjectIDFromStringID(recipeId)
	resultado, err := recipeService.Cook(userCode, recipeObjectId, true)
	if !resultado {
		return nil, err
	}
	res, err := recipeService.recipeRepository.Delete(recipeObjectId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (recipeService *RecipeService) Update(updateRecipe dto.RecipeDto) (*mongo.UpdateResult, error) {
	recipe := updateRecipe.GetModel()
	recipe.UpdateDate = time.Now().String()
	res, err := recipeService.recipeRepository.Update(recipe)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (recipeService *RecipeService) Cook(userCode string, recipeId primitive.ObjectID, cancel bool) (bool, error) {
	recipe, err := recipeService.recipeRepository.GetByCode(userCode, recipeId)
	if err != nil {
		return false, err
	}
	for _, foodQuantity := range recipe.Ingredients {
		var food models.Food
		food.Code = foodQuantity.FoodCode
		if !cancel {
			food.CurrentQuantity = -foodQuantity.Quantity
			//SI LA RECETA SE HACE, LA CANTIDAD SE RESTA DE LA ACTUAL
		} else {
			food.CurrentQuantity = foodQuantity.Quantity
			//SI LA RECETA SE HIZO PERO SE CANCELA LA CANTIDAD SE SUMA A LA ACTUAL
		}
		_, err = recipeService.foodRepository.Update(food, true)
		//SEGUN LO QUE HAYA SUCEDIDO, UPDATEAMOS EN LA BD EL ALIMENTO
		if err != nil {
			return false, err
		}
	}
	return true, nil
}
