package admin

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

type LoginAdminUseCase struct {
	repository interfaces.UserRepository
	settings   *settings.Settings
	validate   *validator.Validate
}

func NewLoginAdminUseCase(repository interfaces.UserRepository,
	settings *settings.Settings,
	validate *validator.Validate) *LoginAdminUseCase {

	return &LoginAdminUseCase{repository, settings, validate}
}

func (usecase *LoginAdminUseCase) Execute(ctx *gin.Context) {
	var adminLogin dto.LoginRequestDTO
	var role = "ADMIN"

	if err := ctx.ShouldBindJSON(&adminLogin); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := usecase.validate.Struct(adminLogin); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	adminDB, err := usecase.repository.GetUserByEmailOrUserName(ctx.Request.Context(), adminLogin.Email, "")

	if err != gorm.ErrRecordNotFound && err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err == gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "usuario admin no existente"})
		return
	}

	if !adminLogin.VerifyPassword(adminDB.Password) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "contraseña incorrecta"})
		return
	}

	jwt, err := token.CreateToken(usecase.settings.Key, adminDB.ID, role)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "login exitoso", "token": jwt})
}
