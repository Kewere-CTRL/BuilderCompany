package company

import (
	"github.com/Kewere-CTRL/BuilderCompany/internal/db"
	"github.com/Kewere-CTRL/BuilderCompany/internal/db/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TelephoneInput struct {
	TelephoneNumber string `json:"telephoneNumber" binding:"required"`
}

func PostCallbackForm(r *gin.Context) {

	var telephoneNumber TelephoneInput
	userId := r.Param("userId")
	userIDInt, err := strconv.Atoi(userId)

	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := r.ShouldBindJSON(&telephoneNumber); err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	callbackForm := &models.CallbackForm{
		UserID:          uint(userIDInt),
		TelephoneNumber: telephoneNumber.TelephoneNumber,
	}

	if err = db.DB.Create(&callbackForm).Error; err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{"error": "Create error"})
		return
	}
	r.JSON(http.StatusOK, gin.H{"message": "Callback form created successfully"})
}
