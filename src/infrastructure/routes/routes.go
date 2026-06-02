package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(engine *gin.Engine) {
	users := engine.Group("/api/v1")
	{
		users.POST("/login", func(ctx *gin.Context) {})
		users.POST("/signup", func(ctx *gin.Context) {})
	}

	roles := engine.Group("/api/v1")
	{
		roles.GET("/roles/:id", func(ctx *gin.Context) {})
		roles.POST("/roles", func(ctx *gin.Context) {})
		roles.PUT("/roles/:id", func(ctx *gin.Context) {})
		roles.DELETE("/roles/:id", func(ctx *gin.Context) {})
	}
}
