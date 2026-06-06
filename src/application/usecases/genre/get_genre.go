package genre

import (
	"movie-reservation-system/src/infrastructure/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GetGenreUseCase struct {
	repository *repositories.GenreRepositoryGorm
}

func NewGetGenreUseCase(repository *repositories.GenreRepositoryGorm) *GetGenreUseCase {
	return &GetGenreUseCase{repository}
}

func (usecase *GetGenreUseCase) Execute(ctx *gin.Context) {
	genreID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "formato invalido en el param :id"})
		return
	}

	genreDB, err := usecase.repository.GetGenreByID(ctx.Request.Context(), uint(genreID))

	if err != gorm.ErrRecordNotFound && err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err == gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "genre no existente"})
		return
	}

	ctx.JSON(http.StatusOK, genreDB)
}
