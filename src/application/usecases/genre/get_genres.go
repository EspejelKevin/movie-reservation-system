package genre

import (
	"movie-reservation-system/src/domain/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetGenresUseCase struct {
	repository interfaces.GenreRepository
}

func NewGetGenresUseCase(repository interfaces.GenreRepository) *GetGenresUseCase {
	return &GetGenresUseCase{repository}
}

func (usecase *GetGenresUseCase) Execute(ctx *gin.Context) {
	genres, err := usecase.repository.GetGenres(ctx.Request.Context())

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"genres": genres})
}
