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

type RecipeHandler struct {
	recipeService services.RecipeServiceInterface
}

func NewRecipeHandler(recipeService services.RecipeServiceInterface) *RecipeHandler {
	return &RecipeHandler{
		recipeService: recipeService,
	}
}

func (recipeHandler *RecipeHandler) GetAll(c *gin.Context) {
	user := utils.GetUserInfoFromContext(c)
	// if user.Code == "" {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "userInfo is required"})
	// 	return
	// }
	var filters dto.FiltersDto
	filters.Aproximation = c.Query("filter_aproximation")
	filters.Moment = c.Query("filter_moment")
	filters.Type = c.Query("filter_type")
	filters.All, _ = strconv.ParseBool(c.Query("filter_all"))
	log.Printf("[handler: RecipeHandler][method: GetAll]")
	recipes, err := recipeHandler.recipeService.GetAll(user.Code, filters)

	if err != nil && err.Error() == "internal" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get recipes from database",
		})
		return
	}

	if len(*recipes) == 0  {
		c.JSON(http.StatusOK, gin.H{
			"error": "Failed to get recipes from database",
			"result": recipes,
		})
		return
	}
	// log.Printf("[handler: RecipeHandler][method: GetAll] lenght of List: %v", len(*recipes))
	c.JSON(http.StatusOK,  gin.H{"result": recipes,})
}

func (recipeHandler *RecipeHandler) Create(c *gin.Context) {
	var recipe dto.RecipeDto
	err := c.ShouldBindJSON(&recipe)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := utils.GetUserInfoFromContext(c)
	recipe.UserCode = user.Code
	// if recipe.UserCode == "" {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "userInfo is required"})
	// 	return
	// }
	log.Printf("[handler: RecipeHandler][method: Create]")
	res, err := recipeHandler.recipeService.Create(recipe)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create recipe " + err.Error()})
		return
	}
	log.Printf("[handler: RecipeHandler][method: Create] recipe: %v", recipe)
	c.JSON(http.StatusOK, res)
}

func (recipeHandler *RecipeHandler) Delete(c *gin.Context) {
	id := c.Param("recipeid")
	log.Printf("[handler: RecipeHandler][method: Delete]")
	user := utils.GetUserInfoFromContext(c)
	_, err := recipeHandler.recipeService.Delete(id, user.Code)

	if err != nil && err.Error() == "internal" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete recipe with id: " + id})
		return
	}

	if err != nil && err.Error() == "notfound" {
		c.JSON(http.StatusOK, gin.H{"message": "Not found any recipe with id: " + id})
		return
	}
	log.Printf("[handler: RecipeHandler][method: Delete] recipeId: %v", id)
	c.JSON(http.StatusOK, gin.H{"message": "Successfully deleted recipe with id: " + id})
}

func (recipeHandler *RecipeHandler) Update(c *gin.Context) {
	id := c.Param("recipeid")
	var recipe dto.RecipeDto
	err := c.ShouldBindJSON(&recipe)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	recipe.Id = id
	log.Printf("[handler: RecipeHandler][method: Update]")
	res, err := recipeHandler.recipeService.Update(recipe)

	if err != nil && err.Error() == "internal" {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Failed to delete recipe"})
	}
	if err != nil && err.Error() == "notfound" {
		c.JSON(http.StatusOK, gin.H{"message": "Not found any recipe with id: " + id})
	}
	log.Printf("[handler: RecipeHandler][method: Update] recipeId: %v", id)
	c.JSON(http.StatusOK, res)
}

func (recipeHandler *RecipeHandler) Cook(c *gin.Context) {
	recipeId := c.Param("recipeid")
	recipeObjectId := utils.GetObjectIDFromStringID(recipeId)
	userInfo := utils.GetUserInfoFromContext(c)
    
	if userInfo.Code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userInfo is required"})
		return
	}
	cancel, _ := strconv.ParseBool(c.Query("cancel"))
	log.Printf("[handler: RecipeHandler][method: Cook]")
	res, err := recipeHandler.recipeService.Cook(userInfo.Code, recipeObjectId, cancel)

	if err != nil && err.Error() == "internal" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel recipe"})
		return
	}
    
	if err != nil && err.Error() == "notfound" {
		c.JSON(http.StatusOK, gin.H{"message": "Not found any recipe with id: " + recipeId})
		return
	}
	log.Printf("[handler: RecipeHandler][method: Cook] recipeId: %v", recipeId)
	var response dto.ResponseCook
	if(res){

		response.Succes= true;
	}
	c.JSON(http.StatusOK, response)
}
