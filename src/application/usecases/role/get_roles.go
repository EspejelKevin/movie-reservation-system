package role

import (
	"movie-reservation-system/src/domain/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetRolesUseCase struct {
	repository interfaces.RoleRepository
}

func NewGetRolesUseCase(repository interfaces.RoleRepository) *GetRolesUseCase {
	return &GetRolesUseCase{repository}
}

func (usecase *GetRolesUseCase) Execute(ctx *gin.Context) {
	roles, err := usecase.repository.GetRoles(ctx.Request.Context())

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"roles": roles})
}
