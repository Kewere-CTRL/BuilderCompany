package router

import (
	"github.com/Kewere-CTRL/BuilderCompany/internal/services/user"
	"github.com/gin-gonic/gin"
)

// Регистрация
func RegisterUserRoutes(r *gin.Engine) {
	r.POST("/users", user.CreateUser)
}

// Авторизация
func LoginUserRoutes(r *gin.Engine) {
	r.GET("/auth/login", user.LoginUser)
}
