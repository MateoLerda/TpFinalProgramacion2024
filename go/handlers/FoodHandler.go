package handlers

import (
	"Status418/go/dto"
	"Status418/go/services"
	"Status418/go/utils"
	"log"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type FoodHandler struct {
	foodService services.FoodServiceInterface
}

func NewFoodHandler(foodService services.FoodServiceInterface) *FoodHandler {
	return &FoodHandler{
		foodService: foodService,
	}
}

func (foodHandler *FoodHandler) GetAll(c *gin.Context) {
	user := utils.GetUserInfoFromContext(c)
	// if user.Code == "" {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "userInfo is required"})
	// 	return
	// }
	var filters dto.FiltersDto
	filters.Aproximation = c.Query("filter_aproximation")
	filters.Type = c.Query("filter_type")
	filters.All, _ = strconv.ParseBool(c.Query("filter_all"))
	log.Printf("[handler: FoodHandler][method: GetAll]")
	foods, err := foodHandler.foodService.GetAll(user.Code, filters)

	if err != nil && err.Error() == "nocontent" {
		c.JSON(http.StatusOK, gin.H{"message": "Not found any foods"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	log.Printf("[handler: FoodHandler][method: GetAll] minimumList: %v", filters.All)
	c.JSON(http.StatusOK, foods)
}

func (foodHandler *FoodHandler) GetByCode(c *gin.Context) {
	user := utils.GetUserInfoFromContext(c)
	// if user.Code == "" {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "userInfo is required"})
	// 	return
	// }
	foodCode := c.Param("foodcode")
	log.Printf("[handler: FoodHandler][method: GetByCode]")
	food, err := foodHandler.foodService.GetByCode(foodCode, user.Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("[handler: FoodHandler][method: GetByCode] food: %v", food)
	c.JSON(http.StatusOK, food)
}

func (foodHandler *FoodHandler) Create(c *gin.Context) {
	var newFood dto.FoodDto
	if err := c.ShouldBindJSON(&newFood); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data", "details": err.Error()})
		return
	}
	user := utils.GetUserInfoFromContext(c)
	log.Printf("[handler: FoodHandler][method: Create]")
	insertedId, err := foodHandler.foodService.Create(newFood, user.Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create food item", "details": err.Error()})
		return
	}
	log.Printf("[handler: FoodHandler][method: Create] insertedId: %v", insertedId)
	c.JSON(http.StatusCreated, gin.H{"message": "Food created successfully", "details": insertedId})
}

func (foodHandler *FoodHandler) Update(c *gin.Context) {
	var updateFood dto.FoodDto
	updateCode := c.Param("foodcode")
	if err := c.ShouldBindJSON(&updateFood); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request body"})
		return
	}
	updateFood.Code = updateCode
	log.Printf("[handler: FoodHandler][method: Update]")
	res, err := foodHandler.foodService.Update(updateFood)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update food item", "details": err.Error()})
		return
	}

	if res.ModifiedCount == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "not food was update"})
		return
	}
	log.Printf("[handler: FoodHandler][method: Create] res: %v", res)
	c.JSON(http.StatusOK, gin.H{"message": "Food updated successfully"})
}

func (foodHandler *FoodHandler) Delete(c *gin.Context) {
	foodCode := c.Param("foodcode")
	user := utils.GetUserInfoFromContext(c)
	// if user.Code == "" {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "userInfo is required"})
	// 	return
	// }
	log.Printf("[handler: FoodHandler][method: Delete]")
	_, err := foodHandler.foodService.Delete(user.Code, foodCode)
	if err != nil && err.Error() == "notfound" {
		c.JSON(http.StatusOK, gin.H{"message": "Not found the requested food to delete"})
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete food item", "details": err.Error()})
		return
	}
	log.Printf("[handler: FoodHandler][method: Delete] foodCode: %v", foodCode)
	c.JSON(http.StatusOK, gin.H{"message": "Food deleted successfully"})
}
