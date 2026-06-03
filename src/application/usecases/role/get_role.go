package role

import (
	"movie-reservation-system/src/infrastructure/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GetRoleUseCase struct {
	repository *repositories.RoleRepositoryGorm
}

func NewGetRoleUseCase(repository *repositories.RoleRepositoryGorm) *GetRoleUseCase {
	return &GetRoleUseCase{repository}
}

func (usecase *GetRoleUseCase) Execute(ctx *gin.Context) {
	rolID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "formato invalido en el param :id"})
		return
	}

	roleDB, err := usecase.repository.GetRoleByID(ctx.Request.Context(), uint(rolID))

	if err != gorm.ErrRecordNotFound && err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err == gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "role no existente"})
		return
	}

	ctx.JSON(http.StatusOK, roleDB)
}
