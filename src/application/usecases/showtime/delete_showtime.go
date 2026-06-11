package showtime

import (
	"movie-reservation-system/src/domain/interfaces"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DeleteShowTimeUseCase struct {
	repository interfaces.ShowTimeRepository
}

func NewDeleteShowTimeUseCase(repository interfaces.ShowTimeRepository) *DeleteShowTimeUseCase {
	return &DeleteShowTimeUseCase{repository}
}

func (usecase *DeleteShowTimeUseCase) Execute(ctx *gin.Context) {
	showID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "formato invalido en el param :id"})
		return
	}

	_, err = usecase.repository.GetShowByID(ctx.Request.Context(), uint(showID))

	if err != gorm.ErrRecordNotFound && err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err == gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "showtime no existente"})
		return
	}

	_, err = usecase.repository.Delete(ctx.Request.Context(), uint(showID))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "eliminacion exitosa"})
}
