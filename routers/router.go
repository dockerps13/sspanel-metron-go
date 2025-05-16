package routers

import (
	"sspanel-metron-go/controllers"
	"sspanel-metron-go/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 公共路由
	r.POST("/api/auth/login", controllers.Login)
	r.POST("/api/auth/register", controllers.Register)

	// 认证后路由
	auth := r.Group("/api")
	auth.Use(middleware.JWTMiddleware())
	{
		auth.GET("/nodes", controllers.ListNodes)
		auth.GET("/subscription", controllers.GetSubscription)
		auth.POST("/payment/callback", controllers.PaymentCallback)
	}

	return r
}
