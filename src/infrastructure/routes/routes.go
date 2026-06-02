package routes

import (
	"movie-reservation-system/src/application/usecases/auth"

	"github.com/gin-gonic/gin"
)

type Router struct {
	registerUser *auth.RegisterUserUseCase
}

func NewRouter(registerUser *auth.RegisterUserUseCase) *Router {
	return &Router{registerUser}
}

func (router *Router) RegisterRoutes(engine *gin.Engine) {
	users := engine.Group("/api/v1")
	{
		users.POST("/login", func(ctx *gin.Context) {})
		users.POST("/signup", router.registerUser.Execute)
	}

	roles := engine.Group("/api/v1")
	{
		roles.GET("/roles/:id", func(ctx *gin.Context) {})
		roles.POST("/roles", func(ctx *gin.Context) {})
		roles.PUT("/roles/:id", func(ctx *gin.Context) {})
		roles.DELETE("/roles/:id", func(ctx *gin.Context) {})
	}
}
