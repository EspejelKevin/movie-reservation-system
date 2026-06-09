package movie

import (
	"movie-reservation-system/src/domain/dto"
	"movie-reservation-system/src/domain/entities"
	"movie-reservation-system/src/domain/interfaces"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type SaveMovieUseCase struct {
	repository interfaces.MovieRepository
	validate   *validator.Validate
}

func NewSaveMovieUseCase(repository interfaces.MovieRepository, validate *validator.Validate) *SaveMovieUseCase {
	return &SaveMovieUseCase{repository, validate}
}

func (usecase *SaveMovieUseCase) Execute(ctx *gin.Context) {
	var movieDTO dto.MovieDTO

	if err := ctx.ShouldBindJSON(&movieDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := usecase.validate.Struct(movieDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	movieDB, err := usecase.repository.GetMovieByTitle(ctx.Request.Context(), movieDTO.Title)

	if err != gorm.ErrRecordNotFound && err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !reflect.DeepEqual(movieDB, entities.Movie{}) {
		ctx.JSON(http.StatusConflict, gin.H{"error": "movie ya existente"})
		return
	}

	newMovie := entities.Movie{
		Title:       movieDTO.Title,
		Description: movieDTO.Description,
		GenreID:     movieDTO.GenreID,
	}

	if err := usecase.repository.Save(ctx.Request.Context(), newMovie); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "registro exitoso"})
}
