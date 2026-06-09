package genre

import (
	"movie-reservation-system/src/domain/interfaces"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DeleteGenreUseCase struct {
	repository interfaces.GenreRepository
}

func NewDeleteGenreUseCase(repository interfaces.GenreRepository) *DeleteGenreUseCase {
	return &DeleteGenreUseCase{repository}
}

func (usecase *DeleteGenreUseCase) Execute(ctx *gin.Context) {
	genreID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "formato invalido en el param :id"})
		return
	}

	_, err = usecase.repository.GetGenreByID(ctx.Request.Context(), uint(genreID))

	if err != gorm.ErrRecordNotFound && err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err == gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "genre no existente"})
		return
	}

	_, err = usecase.repository.Delete(ctx.Request.Context(), uint(genreID))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "eliminacion exitosa"})
}
