package handlers

import (
	"Status418/go/dto"
	"Status418/go/services"
	"Status418/go/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PurchaseHandler struct {
	purchaseService services.PurchaseServiceInterface
}

func NewPurchaseHandler(purchaseService services.PurchaseServiceInterface) *PurchaseHandler {
	return &PurchaseHandler{
		purchaseService: purchaseService,
	}
}

func (purchaseHandler *PurchaseHandler) Create(c *gin.Context) {
	user := (utils.GetUserInfoFromContext(c))
	var newPurchase dto.PurchaseDto
	if c.Request.Body != http.NoBody {
		if err := c.ShouldBindJSON(&newPurchase); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data", "details": err.Error()})
			return
		}
	}
	log.Printf("[handler: PurchaseHandler][method: Create]")
	purchase, err := purchaseHandler.purchaseService.Create(user.Code, newPurchase)

	if err != nil && err.Error() == "nocontent" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create purchase", "details": "You must select a food to buy"})
		return
	}
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create purchase", "details": err.Error()})
		return
	}
	log.Printf("[handler: PurchaseHandler][method: Create] purchase: %v", purchase)
	c.JSON(http.StatusAccepted, purchase)

}
