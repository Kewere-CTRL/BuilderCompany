package company

import (
	"github.com/Kewere-CTRL/BuilderCompany/internal/db"
	"github.com/Kewere-CTRL/BuilderCompany/internal/db/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCalculationServicesCost(r *gin.Context) {

	label := r.Param("label")
	if label == "" {
		r.JSON(http.StatusBadRequest, gin.H{"error": "label is required"})
	}
	var services models.Services

	if err := db.DB.Where("label = ?", label).First(&services).Error; err != nil {
		r.JSON(http.StatusUnauthorized, gin.H{"error": "Services not found"})
		return
	}
	var totalCost = float64(services.Price) * 1.25

	r.JSON(http.StatusOK, gin.H{
		"message":       "Calculation successful",
		"totalCost":     totalCost,
		"servicesLabel": label,
	})
}
