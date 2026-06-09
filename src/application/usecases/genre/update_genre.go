package genre

import (
	"movie-reservation-system/src/domain/dto"
	"movie-reservation-system/src/domain/interfaces"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UpdateGenreUseCase struct {
	repository interfaces.GenreRepository
	validate   *validator.Validate
}

func NewUpdateGenreUseCase(repository interfaces.GenreRepository,
	validate *validator.Validate) *UpdateGenreUseCase {
	return &UpdateGenreUseCase{repository, validate}
}

func (usecase *UpdateGenreUseCase) Execute(ctx *gin.Context) {
	genreID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "formato invalido en el param :id"})
		return
	}

	var genreDTO dto.GenreDTO

	if err := ctx.ShouldBindJSON(&genreDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := usecase.validate.Struct(genreDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	_, err = usecase.repository.Update(ctx.Request.Context(), uint(genreID), genreDTO.Name)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "actualizacion exitosa"})
}
