package auth

import (
	"movie-reservation-system/src/domain/dto"
	"movie-reservation-system/src/domain/entities"
	"movie-reservation-system/src/domain/interfaces"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RegisterUserUseCase struct {
	repository interfaces.UserRepository
	validate   *validator.Validate
}

func NewRegisterUserUseCase(repository interfaces.UserRepository,
	validate *validator.Validate) *RegisterUserUseCase {

	return &RegisterUserUseCase{repository, validate}
}

func (usecase *RegisterUserUseCase) Execute(ctx *gin.Context) {
	var userDTO dto.UserDTO
	var roleID = 2

	if err := ctx.ShouldBindJSON(&userDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := usecase.validate.Struct(userDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userDB, err := usecase.repository.GetUserByEmail(ctx.Request.Context(), userDTO.Email)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !reflect.DeepEqual(userDB, entities.User{}) {
		ctx.JSON(http.StatusConflict, gin.H{"error": "usuario ya existente"})
		return
	}

	newUser := entities.User{
		Name:     userDTO.Name,
		UserName: userDTO.UserName,
		Email:    userDTO.Email,
		Password: userDTO.Password,
		RoleID:   uint(roleID),
	}

	if err := usecase.repository.Save(ctx.Request.Context(), newUser); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "registro exitoso"})
}
