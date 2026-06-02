package routes

import (
	"movie-reservation-system/src/application/usecases/auth"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	registerUser *auth.RegisterUserUseCase
}

func (routes *Routes) RegisterRoutes(engine *gin.Engine) {
	users := engine.Group("/api/v1")
	{
		users.POST("/login", func(ctx *gin.Context) {})
		users.POST("/signup", routes.registerUser.Execute)
	}

	roles := engine.Group("/api/v1")
	{
		roles.GET("/roles/:id", func(ctx *gin.Context) {})
		roles.POST("/roles", func(ctx *gin.Context) {})
		roles.PUT("/roles/:id", func(ctx *gin.Context) {})
		roles.DELETE("/roles/:id", func(ctx *gin.Context) {})
	}
}
