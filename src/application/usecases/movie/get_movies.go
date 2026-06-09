package movie

import (
	"movie-reservation-system/src/domain/dto"
	"movie-reservation-system/src/domain/entities"
	"movie-reservation-system/src/domain/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetMoviesUseCase struct {
	repository interfaces.MovieRepository
}

func NewGetMoviesUseCase(repository interfaces.MovieRepository) *GetMoviesUseCase {
	return &GetMoviesUseCase{repository}
}

func (usecase *GetMoviesUseCase) Execute(ctx *gin.Context) {
	movies, err := usecase.repository.GetMovies(ctx.Request.Context())

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"movies": usecase.mapMovies(movies)})
}

func (usecase *GetMoviesUseCase) mapMovies(movies []entities.Movie) []dto.MovieResponse {
	var movieResponse []dto.MovieResponse

	for _, movie := range movies {
		movieResponse = append(movieResponse, dto.MovieResponse{
			ID:          movie.ID,
			Title:       movie.Title,
			Description: movie.Description,
			Image:       movie.Image.String,
			Genre:       movie.Genre.Name,
		})
	}

	return movieResponse
}
