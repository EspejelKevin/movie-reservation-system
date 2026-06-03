package auth

import (
	"movie-reservation-system/src/application/usecases/auth/token"
	"movie-reservation-system/src/domain/dto"
	"movie-reservation-system/src/domain/interfaces"
	"movie-reservation-system/src/domain/settings"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type LoginUserUseCase struct {
	repository interfaces.UserRepository
	settings   *settings.Settings
	validate   *validator.Validate
}

func NewLoginUserUseCase(repository interfaces.UserRepository,
	settings *settings.Settings,
	validate *validator.Validate) *LoginUserUseCase {

	return &LoginUserUseCase{repository, settings, validate}
}

func (usecase *LoginUserUseCase) Execute(ctx *gin.Context) {
	var userLogin dto.LoginRequestDTO
	var role = "REGULAR_USER"

	if err := ctx.ShouldBindJSON(&userLogin); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := usecase.validate.Struct(userLogin); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userDB, err := usecase.repository.GetUserByEmailOrUserName(ctx.Request.Context(), userLogin.Email, "")

	if err != gorm.ErrRecordNotFound && err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err == gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "usuario no existente"})
		return
	}

	if !userLogin.VerifyPassword(userDB.Password) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "contraseña incorrecta"})
		return
	}

	jwt, err := token.CreateToken(usecase.settings.Key, userDB.ID, role)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "login exitoso", "token": jwt})
}
