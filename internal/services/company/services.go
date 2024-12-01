package company

import (
	"github.com/Kewere-CTRL/BuilderCompany/internal/db"
	"github.com/Kewere-CTRL/BuilderCompany/internal/db/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetServices(r *gin.Context) {
	var services []models.Services

	if err := db.DB.Find(&services).Error; err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить данные"})
		return
	}
	if len(services) == 0 {
		r.JSON(http.StatusOK, gin.H{"message": "Services not found"})
	}

	r.JSON(http.StatusOK, gin.H{"data": services})

}
