package movie

import (
	"movie-reservation-system/src/domain/interfaces"
	"movie-reservation-system/src/domain/settings"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DeleteMovieUseCase struct {
	repository interfaces.MovieRepository
	settings   *settings.Settings
}

func NewDeleteMovieUseCase(repository interfaces.MovieRepository,
	settings *settings.Settings) *DeleteMovieUseCase {
	return &DeleteMovieUseCase{repository, settings}
}

func (usecase *DeleteMovieUseCase) Execute(ctx *gin.Context) {
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

	_, err = usecase.repository.Delete(ctx.Request.Context(), uint(movieID))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	os.Remove(filepath.Join(usecase.settings.ImagePath, filepath.Base(movieDB.Image.String)))

	ctx.JSON(http.StatusOK, gin.H{"message": "eliminacion exitosa"})

}
