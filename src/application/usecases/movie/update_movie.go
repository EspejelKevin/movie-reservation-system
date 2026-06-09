package movie

import (
	"movie-reservation-system/src/domain/dto"
	"movie-reservation-system/src/domain/interfaces"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UpdateMovieUseCase struct {
	repository interfaces.MovieRepository
	validate   *validator.Validate
}

func NewUpdateMovieUseCase(repository interfaces.MovieRepository,
	validate *validator.Validate) *UpdateMovieUseCase {
	return &UpdateMovieUseCase{repository, validate}
}

func (usecase *UpdateMovieUseCase) Execute(ctx *gin.Context) {
	movieID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "formato invalido en el param :id"})
		return
	}

	var movieDTO dto.MovieDTO

	if err := ctx.ShouldBindJSON(&movieDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := usecase.validate.Struct(movieDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = usecase.repository.GetMovieByID(ctx.Request.Context(), uint(movieID))

	if err != gorm.ErrRecordNotFound && err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err == gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "movie no existente"})
		return
	}

	_, err = usecase.repository.Update(ctx.Request.Context(), uint(movieID), movieDTO)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "actualizacion exitosa"})
}
