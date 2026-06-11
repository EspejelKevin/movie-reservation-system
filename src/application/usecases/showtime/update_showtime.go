package showtime

import (
	"movie-reservation-system/src/domain/dto"
	"movie-reservation-system/src/domain/interfaces"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UpdateShowTimeUseCase struct {
	repository interfaces.ShowTimeRepository
	validate   *validator.Validate
}

func NewUpdateShowTimeUseCase(repository interfaces.ShowTimeRepository,
	validate *validator.Validate) *UpdateShowTimeUseCase {
	return &UpdateShowTimeUseCase{repository, validate}
}

func (usecase *UpdateShowTimeUseCase) Execute(ctx *gin.Context) {
	showID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "formato invalido en el param :id"})
		return
	}

	var showDTO dto.ShowDTO

	if err := ctx.ShouldBindJSON(&showDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := usecase.validate.Struct(showDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	_, err = usecase.repository.Update(ctx.Request.Context(), uint(showID), showDTO)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "actualizacion exitosa"})
}
