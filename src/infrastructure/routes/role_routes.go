package routes

import (
	"movie-reservation-system/src/application/usecases/role"
	"movie-reservation-system/src/infrastructure/middlewares"

	"github.com/gin-gonic/gin"
)

type RouterRole struct {
	getRoles   *role.GetRolesUseCase
	getRole    *role.GetRoleUseCase
	saveRole   *role.SaveRoleUseCase
	updateRole *role.UpdateRoleUseCase
	deleteRole *role.DeleteRoleUseCase

	middlewares *middlewares.Middlewares
}

func NewRouterRole(getRoles *role.GetRolesUseCase, getRole *role.GetRoleUseCase,
	saveRole *role.SaveRoleUseCase, updateRole *role.UpdateRoleUseCase,
	deleteRole *role.DeleteRoleUseCase, middlewares *middlewares.Middlewares) *RouterRole {
	return &RouterRole{getRoles, getRole, saveRole, updateRole, deleteRole, middlewares}
}

func (router *RouterRole) RegisterRoutesRole(engine *gin.Engine) {
	roles := engine.Group("/api/v1")
	roles.Use(router.middlewares.Authentication())
	{
		roles.GET("/roles", router.middlewares.Authorization("ADMIN"), router.getRoles.Execute)
		roles.GET("/roles/:id", router.middlewares.Authorization("ADMIN"), router.getRole.Execute)
		roles.POST("/roles", router.middlewares.Authorization("ADMIN"), router.saveRole.Execute)
		roles.PUT("/roles/:id", router.middlewares.Authorization("ADMIN"), router.updateRole.Execute)
		roles.DELETE("/roles/:id", router.middlewares.Authorization("ADMIN"), router.deleteRole.Execute)
	}
}
