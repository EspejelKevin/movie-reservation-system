package routes

import (
	"movie-reservation-system/src/application/usecases/admin"
	"movie-reservation-system/src/application/usecases/auth"

	"github.com/gin-gonic/gin"
)

type Router struct {
	registerUser *auth.RegisterUserUseCase
	loginUser    *auth.LoginUserUseCase
	loginAdmin   *admin.LoginAdminUseCase
}

func NewRouter(registerUser *auth.RegisterUserUseCase, loginUser *auth.LoginUserUseCase,
	loginAdmin *admin.LoginAdminUseCase) *Router {
	return &Router{registerUser, loginUser, loginAdmin}
}

func (router *Router) RegisterRoutes(engine *gin.Engine) {
	admin := engine.Group("/api/v1/admin")
	{
		admin.POST("/login", router.loginAdmin.Execute)
	}

	users := engine.Group("/api/v1")
	{
		users.POST("/login", router.loginUser.Execute)
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
