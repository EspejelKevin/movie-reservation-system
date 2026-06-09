package movie

import (
	"fmt"
	"movie-reservation-system/src/domain/interfaces"
	"movie-reservation-system/src/domain/settings"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UpdateImageMovieUseCase struct {
	repository interfaces.MovieRepository
	settings   *settings.Settings
}

func NewUpdateImageMovieUseCase(repository interfaces.MovieRepository,
	settings *settings.Settings) *UpdateImageMovieUseCase {
	return &UpdateImageMovieUseCase{repository, settings}
}

func (usecase *UpdateImageMovieUseCase) Execute(ctx *gin.Context) {
	movieID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "formato invalido en el param :id"})
		return
	}

	file, err := ctx.FormFile("file")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	imageName := fmt.Sprintf("%s.png", uuid.NewString())
	dst := filepath.Join("./images", imageName)
	err = ctx.SaveUploadedFile(file, dst)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	imagePath := filepath.Join("/images", imageName)

	_, err = usecase.repository.UpdateImage(ctx.Request.Context(), uint(movieID), imagePath)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	os.Remove(filepath.Join(usecase.settings.ImagePath, filepath.Base(movieDB.Image.String)))

	ctx.JSON(http.StatusOK, gin.H{"message": "actualizacion exitosa"})
}
