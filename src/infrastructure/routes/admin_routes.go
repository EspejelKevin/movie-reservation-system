package routes

import (
	"movie-reservation-system/src/application/usecases/admin"

	"github.com/gin-gonic/gin"
)

type RouterAdmin struct {
	loginAdmin *admin.LoginAdminUseCase
}

func NewRouterAdmin(loginAdmin *admin.LoginAdminUseCase) *RouterAdmin {
	return &RouterAdmin{loginAdmin}
}

func (router *RouterAdmin) RegisterRoutesAdmin(engine *gin.Engine) {
	admin := engine.Group("/api/v1/admin")
	{
		admin.POST("/login", router.loginAdmin.Execute)
	}
}
