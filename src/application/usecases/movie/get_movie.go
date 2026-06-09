package movie

import (
	"movie-reservation-system/src/domain/dto"
	"movie-reservation-system/src/domain/interfaces"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GetMovieUseCase struct {
	repository interfaces.MovieRepository
}

func NewGetMovieUseCase(repository interfaces.MovieRepository) *GetMovieUseCase {
	return &GetMovieUseCase{repository}
}

func (usecase *GetMovieUseCase) Execute(ctx *gin.Context) {
	movieID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "formato invalido en el param :id"})
		return
	}

	movieDB, err := usecase.repository.GetMovieByID(ctx.Request.Context(), uint(movieID))

	if err != gorm.ErrRecordNotFound && err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err == gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "movie no existente"})
		return
	}

	ctx.JSON(http.StatusOK, dto.MovieResponse{
		ID:          movieDB.ID,
		Title:       movieDB.Title,
		Description: movieDB.Description,
		Image:       movieDB.Image.String,
		Genre:       movieDB.Genre.Name,
	})
}
