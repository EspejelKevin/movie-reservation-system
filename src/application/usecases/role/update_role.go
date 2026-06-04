package role

import (
	"movie-reservation-system/src/domain/dto"
	"movie-reservation-system/src/infrastructure/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UpdateRoleUseCase struct {
	repository *repositories.RoleRepositoryGorm
	validate   *validator.Validate
}

func NewUpdateRoleUseCase(repository *repositories.RoleRepositoryGorm,
	validate *validator.Validate) *UpdateRoleUseCase {
	return &UpdateRoleUseCase{repository, validate}
}

func (usecase *UpdateRoleUseCase) Execute(ctx *gin.Context) {
	rolID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "formato invalido en el param :id"})
		return
	}

	var roleDTO dto.RoleDTO

	if err := ctx.ShouldBindJSON(&roleDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := usecase.validate.Struct(roleDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	_, err = usecase.repository.Update(ctx.Request.Context(), uint(rolID), roleDTO.Name)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "actualizacion exitosa"})
}
