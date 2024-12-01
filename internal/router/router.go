package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5137",
			"http://localhost:5138",
		},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:5137"
		},
		MaxAge: 12 * time.Hour,
	}))

	RegisterUserRoutes(r)
	LoginUserRoutes(r)
	GetCompanyRoutes(r)
	GetServicesRoutes(r)
	return r
}
