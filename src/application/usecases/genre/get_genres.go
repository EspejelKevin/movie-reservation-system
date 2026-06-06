package genre

import (
	"movie-reservation-system/src/infrastructure/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetGenresUseCase struct {
	repository *repositories.GenreRepositoryGorm
}

func NewGetGenresUseCase(repository *repositories.GenreRepositoryGorm) *GetGenresUseCase {
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
