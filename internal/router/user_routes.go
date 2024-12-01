package router

import (
	"github.com/Kewere-CTRL/BuilderCompany/internal/handlers"
	"github.com/gin-gonic/gin"
)

// Регистрация
func RegisterUserRoutes(r *gin.Engine) {
	r.POST("/users", handlers.CreateUser)
}

// Авторизация
func LoginUserRoutes(r *gin.Engine) {
	r.POST("/auth/login", handlers.LoginUser)
}
