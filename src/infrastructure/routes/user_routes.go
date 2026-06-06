package routes

import (
	"movie-reservation-system/src/application/usecases/auth"

	"github.com/gin-gonic/gin"
)

type RouterUser struct {
	registerUser *auth.RegisterUserUseCase
	loginUser    *auth.LoginUserUseCase
}

func NewRouterUser(registerUser *auth.RegisterUserUseCase, loginUser *auth.LoginUserUseCase) *RouterUser {
	return &RouterUser{registerUser, loginUser}
}

func (router *RouterUser) RegisterRoutesUser(engine *gin.Engine) {
	users := engine.Group("/api/v1")
	{
		users.POST("/login", router.loginUser.Execute)
		users.POST("/signup", router.registerUser.Execute)
	}
}
