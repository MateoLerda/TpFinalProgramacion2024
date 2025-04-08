package handlers

import (
	"Status418/go/services"
	"Status418/go/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReportHandler struct {
	reportService services.ReportServiceInterface
}

func NewReportHandler(reportService services.ReportServiceInterface) *ReportHandler {
	return &ReportHandler{
		reportService: reportService,
	}
}

func (reportHandler *ReportHandler) GetRecipeMomentReport(c *gin.Context) {
	user := utils.GetUserInfoFromContext(c)
	report, err := reportHandler.reportService.GetRecipesReport(user.Code, true)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to create report"})
		return
	}
	c.JSON(http.StatusOK, report)
}

func (reportHandler *ReportHandler) GetRecipeFoodTypeReport(c *gin.Context) {
	user := utils.GetUserInfoFromContext(c)
	report, err :=reportHandler.reportService.GetRecipesReport(user.Code, false)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to create report"})
		return
	}
	c.JSON(http.StatusOK, report)
}

func (reportHandler *ReportHandler) GetPurchaseReport(c *gin.Context) {	user := utils.GetUserInfoFromContext(c)
	report, err := reportHandler.reportService.GetCostReport(user.Code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to create report"})
		return
	}
	c.JSON(http.StatusOK, report)
}