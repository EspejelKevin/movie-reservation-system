package routes

import (
	"movie-reservation-system/src/application/usecases/admin"
	"movie-reservation-system/src/application/usecases/auth"
	"movie-reservation-system/src/application/usecases/role"
	"movie-reservation-system/src/infrastructure/middlewares"

	"github.com/gin-gonic/gin"
)

type Router struct {
	registerUser *auth.RegisterUserUseCase
	loginUser    *auth.LoginUserUseCase
	loginAdmin   *admin.LoginAdminUseCase
	getRoles     *role.GetRolesUseCase

	middlewares *middlewares.Middlewares
}

func NewRouter(registerUser *auth.RegisterUserUseCase, loginUser *auth.LoginUserUseCase,
	loginAdmin *admin.LoginAdminUseCase, getRoles *role.GetRolesUseCase,
	middlewares *middlewares.Middlewares) *Router {
	return &Router{registerUser, loginUser, loginAdmin, getRoles, middlewares}
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
	roles.Use(router.middlewares.Authentication())
	{
		roles.GET("/roles", router.middlewares.Authorization("ADMIN"), router.getRoles.Execute)
		roles.GET("/roles/:id", router.middlewares.Authorization("ADMIN"), func(ctx *gin.Context) {})
		roles.POST("/roles", router.middlewares.Authorization("ADMIN"), func(ctx *gin.Context) {})
		roles.PUT("/roles/:id", router.middlewares.Authorization("ADMIN"), func(ctx *gin.Context) {})
		roles.DELETE("/roles/:id", router.middlewares.Authorization("ADMIN"), func(ctx *gin.Context) {})
	}
}
