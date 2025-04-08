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

type PurchaseServiceInterface interface {
	Create(userCode string, newPurchase dto.PurchaseDto) (*mongo.InsertOneResult, error)
}

type PurchaseService struct {
	purchaseRepository repositories.PurchaseRepositoryInterface
	foodRepository repositories.FoodRepositoryInterface
}

func NewPurchaseService(purchaseRepository repositories.PurchaseRepositoryInterface, foodRepository repositories.FoodRepositoryInterface) *PurchaseService {
	return &PurchaseService{
		purchaseRepository: purchaseRepository,
		foodRepository: foodRepository,
	}
}

func calculatePurchaseAllFoods(foods []models.Food) models.Purchase {
	var purchase models.Purchase
	for _, food := range foods {
		purchase.TotalCost += food.UnitPrice * (float64)(food.MinimumQuantity-food.CurrentQuantity)
		purchase.Foods = append(purchase.Foods, models.FoodQuantity{
			FoodCode: food.Code,
			Name: food.Name,
			Quantity: food.MinimumQuantity - food.CurrentQuantity,
		})
	}
	return purchase
}


func (purchaseService *PurchaseService) Create(userCode string, newPurchase dto.PurchaseDto) (*mongo.InsertOneResult, error) {
	var foods []models.Food
	var err error
	var purchase models.Purchase
	
	if len(newPurchase.Foods) != 0 {
		var food models.Food
		for _, foodQuantity := range newPurchase.Foods {
			if foodQuantity.Quantity < 0 {
				return nil, errors.New("quantity cannot be negative")
			}
			foodObjectId := utils.GetObjectIDFromStringID(foodQuantity.FoodCode)
			food, err = purchaseService.foodRepository.GetByCode(foodObjectId, userCode)
			if err != nil {
				return nil, err
			}
			purchase.TotalCost += food.UnitPrice * float64(foodQuantity.Quantity)
			purchase.Foods = append(purchase.Foods, models.FoodQuantity{FoodCode: foodObjectId, Name: food.Name, Quantity: foodQuantity.Quantity})
		}
	} else {
		var filter models.Filter
		filter.All= false 
		foods, err = purchaseService.foodRepository.GetAll(userCode, filter)
		if err != nil {
			return nil, err
		}
		purchase = calculatePurchaseAllFoods(foods)
	}

	purchase.PurchaseDate = time.Now().String()
	purchase.UserCode = userCode
	create, err := purchaseService.purchaseRepository.Create(purchase)
	if err != nil {
		return nil, err
	}

	for _, food := range purchase.Foods {
		var updatedFood models.Food
		updatedFood.Code = food.FoodCode
		updatedFood.CurrentQuantity = food.Quantity
		_,err := purchaseService.foodRepository.Update(updatedFood, true)
		if err != nil {
			return nil, err
		}
	}
	return create, nil
}
