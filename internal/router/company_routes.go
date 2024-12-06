package router

import (
	"github.com/Kewere-CTRL/BuilderCompany/internal/services/company"
	"github.com/gin-gonic/gin"
)

// Проекты компании
func GetCompanyRoutes(r *gin.Engine) {
	r.GET("/company/project", company.GetCompany)
}

// Услуги компании
func GetServicesRoutes(r *gin.Engine) {
	r.GET("/company/services", company.GetServices)
}

// Расчет стоимости услуг
func GetCalculationServicesCostRoutes(r *gin.Engine) {
	r.GET("/company/calculation/:label", company.GetCalculationServicesCost)
}

// Получение всех заявок
func GetCallbackCallbackFormRoutes(r *gin.Engine) {
	r.GET("/company/callback/all", company.GetCallback)
}

// Создание заявки
func PostCallbackFormRoutes(r *gin.Engine) {
	r.POST("/company/callback/:userId", company.PostCallbackForm)
}
