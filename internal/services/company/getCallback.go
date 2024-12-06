package company

import (
	"github.com/Kewere-CTRL/BuilderCompany/internal/db"
	"github.com/Kewere-CTRL/BuilderCompany/internal/db/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCallback(r *gin.Context) {
	var callbackForm []models.CallbackForm

	if err := db.DB.Preload("User").Find(&callbackForm).Error; err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить данные"})
		return
	}
	if len(callbackForm) == 0 {
		r.JSON(http.StatusOK, gin.H{"message": "CallbackForm not found"})
	}

	r.JSON(http.StatusOK, gin.H{"data": callbackForm})

}
