package company

import (
	"github.com/Kewere-CTRL/BuilderCompany/internal/db"
	"github.com/Kewere-CTRL/BuilderCompany/internal/db/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCompany(r *gin.Context) {
	var project []models.Project

	if err := db.DB.Find(&project).Error; err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения данных"})
		return
	}

	if len(project) == 0 {
		r.JSON(http.StatusOK, gin.H{"message": "Project not found"})
	}

	r.IndentedJSON(http.StatusOK, gin.H{"data": project})
}
