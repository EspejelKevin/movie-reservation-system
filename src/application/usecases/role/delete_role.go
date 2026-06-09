package role

import (
	"movie-reservation-system/src/domain/interfaces"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DeleteRoleUseCase struct {
	repository interfaces.RoleRepository
}

func NewDeleteRoleUseCase(repository interfaces.RoleRepository) *DeleteRoleUseCase {
	return &DeleteRoleUseCase{repository}
}

func (usecase *DeleteRoleUseCase) Execute(ctx *gin.Context) {
	rolID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "formato invalido en el param :id"})
		return
	}

	_, err = usecase.repository.GetRoleByID(ctx.Request.Context(), uint(rolID))

	if err != gorm.ErrRecordNotFound && err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err == gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "role no existente"})
		return
	}

	_, err = usecase.repository.Delete(ctx.Request.Context(), uint(rolID))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "eliminacion exitosa"})
}
