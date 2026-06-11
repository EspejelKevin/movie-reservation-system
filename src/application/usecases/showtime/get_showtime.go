package showtime

import (
	"movie-reservation-system/src/domain/dto"
	"movie-reservation-system/src/domain/interfaces"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GetShowTimeUseCase struct {
	repository interfaces.ShowTimeRepository
}

func NewGetShowTimeUseCase(repository interfaces.ShowTimeRepository) *GetShowTimeUseCase {
	return &GetShowTimeUseCase{repository}
}

func (usecase *GetShowTimeUseCase) Execute(ctx *gin.Context) {
	showID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "formato invalido en el param :id"})
		return
	}

	showDB, err := usecase.repository.GetShowByID(ctx.Request.Context(), uint(showID))

	if err != gorm.ErrRecordNotFound && err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err == gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "showtime no existente"})
		return
	}

	ctx.JSON(http.StatusOK, dto.ShowResponse{
		ID:                     showDB.ID,
		Folio:                  showDB.Folio,
		StartDate:              showDB.StartDate.Format("2006-01-02 15:04:05"),
		EndDate:                showDB.EndDate.Format("2006-01-02 15:04:05"),
		AvailableQuantitySeats: showDB.AvailableQuantitySeats,
		UnavailableSeats:       showDB.UnavailableSeats,
		Movie: dto.MovieResponse{
			ID:          showDB.Movie.ID,
			Title:       showDB.Movie.Title,
			Description: showDB.Movie.Description,
			Image:       showDB.Movie.Image.String,
			Genre:       showDB.Movie.Genre.Name,
		}})
}
