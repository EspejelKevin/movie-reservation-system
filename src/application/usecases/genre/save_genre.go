package genre

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

type SaveGenreUseCase struct {
	repository *repositories.GenreRepositoryGorm
	validate   *validator.Validate
}

func NewSaveGenreUseCase(repository *repositories.GenreRepositoryGorm, validate *validator.Validate) *SaveGenreUseCase {
	return &SaveGenreUseCase{repository, validate}
}

func (usecase *SaveGenreUseCase) Execute(ctx *gin.Context) {
	var genreDTO dto.GenreDTO

	if err := ctx.ShouldBindJSON(&genreDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := usecase.validate.Struct(genreDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	genreDB, err := usecase.repository.GetGenreByName(ctx.Request.Context(), genreDTO.Name)

	if err != gorm.ErrRecordNotFound && err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !reflect.DeepEqual(genreDB, entities.Genre{}) {
		ctx.JSON(http.StatusConflict, gin.H{"error": "genre ya existente"})
		return
	}

	newGenre := entities.Genre{
		Name: genreDTO.Name,
	}

	if err := usecase.repository.Save(ctx.Request.Context(), newGenre); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "registro exitoso"})
}
