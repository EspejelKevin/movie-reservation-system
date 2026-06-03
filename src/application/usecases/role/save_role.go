package role

import (
	"movie-reservation-system/src/domain/dto"
	"movie-reservation-system/src/domain/entities"
	"movie-reservation-system/src/infrastructure/repositories"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type SaveRole struct {
	repository *repositories.RoleRepositoryGorm
	validate   *validator.Validate
}

func NewSaveRole(repository *repositories.RoleRepositoryGorm, validate *validator.Validate) *SaveRole {
	return &SaveRole{repository, validate}
}

func (usecase *SaveRole) Execute(ctx *gin.Context) {
	var roleDTO dto.RoleDTO

	if err := ctx.ShouldBindJSON(&roleDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := usecase.validate.Struct(roleDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	roleDB, err := usecase.repository.GetRoleByName(ctx.Request.Context(), roleDTO.Name)

	if err != gorm.ErrRecordNotFound && err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !reflect.DeepEqual(roleDB, entities.Role{}) {
		ctx.JSON(http.StatusConflict, gin.H{"error": "role ya existente"})
		return
	}

	newRole := entities.Role{
		Name: roleDTO.Name,
	}

	if err := usecase.repository.Save(ctx.Request.Context(), newRole); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "registro exitoso"})
}
