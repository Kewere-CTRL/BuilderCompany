package router

import (
	"github.com/Kewere-CTRL/BuilderCompany/internal/services/company"
	"github.com/gin-gonic/gin"
)

// Проекты компании
func GetCompany(r *gin.Engine) {
	r.GET("/company/project", company.GetCompany)
}

// Услуги компании
func GetServices(r *gin.Engine) {
	r.GET("/company/services", company.GetServices)
}
