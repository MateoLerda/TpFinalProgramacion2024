package services

import (
	"Status418/go/dto"
	"Status418/go/models"
	"Status418/go/repositories"
	"Status418/go/utils"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type FoodServiceInterface interface {
	GetAll(userCode string, filter dto.FiltersDto) (*[]dto.FoodDto, error)
	GetByCode(foodCode string, userCode string) (*dto.FoodDto, error)
	Create(newFood dto.FoodDto, userCode string) (*mongo.InsertOneResult, error)
	Update(updateFood dto.FoodDto) (*mongo.UpdateResult, error)
	Delete(userCode string, foodcode string) (*mongo.DeleteResult, error)
}

type FoodService struct {
	foodRepository   repositories.FoodRepositoryInterface
	recipeRepository repositories.RecipeRepositoryInterface
}

func NewFoodService(foodRepository repositories.FoodRepositoryInterface, recipeRepository repositories.RecipeRepositoryInterface) *FoodService {
	return &FoodService{
		foodRepository:   foodRepository,
		recipeRepository: recipeRepository,
	}
}

func (foodService *FoodService) GetAll(userCode string, filter dto.FiltersDto) (*[]dto.FoodDto, error) {
	var foodsDTO []dto.FoodDto
	foods, err := foodService.foodRepository.GetAll(userCode, filter.GetModel())
	if err != nil {
		return nil, err
	}

	for _, food := range foods {
		foodDTO := dto.NewFoodDto(food)
		foodsDTO = append(foodsDTO, *foodDTO)
	}
	return &foodsDTO, nil
}

func (foodService *FoodService) GetByCode(foodCode string, userCode string) (*dto.FoodDto, error) {
	food, err := foodService.foodRepository.GetByCode(utils.GetObjectIDFromStringID(foodCode), userCode)
	if err != nil {
		return nil, err
	}
	foodDto := dto.NewFoodDto(food)
	return foodDto, nil
}

func validateDto(foodDto dto.FoodDto) error {

	if foodDto.Type == "" {
		return errors.New("Invalid food type")
	}

	if len(foodDto.Moments) == 0 {
		return errors.New("Invalid food moments")
	}

	if foodDto.Name == "" {
		return errors.New("Invalid food name")
	}

	if foodDto.UnitPrice <= 0 {
		return errors.New("Food unit price cannot be equals or lower than 0")
	}

	if foodDto.CurrentQuantity < 0 {
		return errors.New("Food current quantity cannot be lower than 0")
	}

	if foodDto.MinimumQuantity < 0 {
		return errors.New("Food minimum quantity cannot be lower than 0")
	}

	return nil
}

func (foodService *FoodService) Create(foodDto dto.FoodDto, userCode string) (*mongo.InsertOneResult, error) {
	err := validateDto(foodDto)
	if err != nil {
		return nil, err
	}
	food := foodDto.GetModel()
	food.CreationDate = time.Now().String()
	food.UserCode = userCode
	res, err := foodService.foodRepository.Create(food)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (foodService *FoodService) Update(foodDto dto.FoodDto) (*mongo.UpdateResult, error) {
	if foodDto.UnitPrice <= 0 {
		return nil, errors.New("Food unit price cannot be equals or lower than 0")
	}

	if foodDto.CurrentQuantity < 0 {
		return nil, errors.New("Food current quantity cannot be lower than 0")
	}

	if foodDto.MinimumQuantity < 0 {
		return nil, errors.New("Food minimum quantity cannot be lower than 0")
	}
	food := foodDto.GetModel()
	res, err := foodService.foodRepository.Update(food, false)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (foodService *FoodService) Delete(userCode string, foodCode string) (*mongo.DeleteResult, error) {
	filter := models.Filter{All: true}
	recipes, _ := foodService.recipeRepository.GetAll(userCode, filter)
	foodObjectId := utils.GetObjectIDFromStringID(foodCode)
	for _, recipe := range recipes {
		for _, food := range recipe.Ingredients {
			if food.FoodCode == foodObjectId {
				foodService.recipeRepository.Delete(recipe.Id)
				break
			}
		}
	}

	res, err := foodService.foodRepository.Delete(foodObjectId)
	if err != nil {
		return nil, err
	}
	return res, nil
}
