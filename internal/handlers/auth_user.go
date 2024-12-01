package handlers

import (
	"github.com/Kewere-CTRL/BuilderCompany/internal/db"
	"github.com/Kewere-CTRL/BuilderCompany/internal/db/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func LoginUser(r *gin.Context) {
	var LoginData struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	if err := r.ShouldBindJSON(&LoginData); err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Json"})
		return
	}

	var user models.User

	if err := db.DB.Where("name = ?", LoginData.Name).First(&user).Error; err != nil {
		r.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(LoginData.Password))
	if err != nil {
		r.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Password"})
		return
	}
	r.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user})
}
