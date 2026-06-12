package showtime

import (
	"movie-reservation-system/src/domain/dto"
	"movie-reservation-system/src/domain/entities"
	"movie-reservation-system/src/domain/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetShowTimesUseCase struct {
	repository interfaces.ShowTimeRepository
}

func NewGetShowTimesUseCase(repository interfaces.ShowTimeRepository) *GetShowTimesUseCase {
	return &GetShowTimesUseCase{repository}
}

func (usecase *GetShowTimesUseCase) Execute(ctx *gin.Context) {
	shows, err := usecase.repository.GetShows(ctx.Request.Context())

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"shows": usecase.mapShows(shows)})
}

func (usecase *GetShowTimesUseCase) mapShows(shows []entities.ShowTime) []dto.ShowResponse {
	showResponse := make([]dto.ShowResponse, 0, len(shows))

	for _, show := range shows {
		showResponse = append(showResponse, dto.ShowResponse{
			ID:                     show.ID,
			Folio:                  show.Folio,
			StartDate:              show.StartDate.Format("2006-01-02 15:04:05"),
			EndDate:                show.EndDate.Format("2006-01-02 15:04:05"),
			AvailableQuantitySeats: show.QuantitySeats - *show.UnavailableSeats,
			UnavailableSeats:       show.UnavailableSeats,
			Movie: dto.MovieResponse{
				ID:          show.Movie.ID,
				Title:       show.Movie.Title,
				Description: show.Movie.Description,
				Image:       show.Movie.Image.String,
				Genre:       show.Movie.Genre.Name,
			},
		})
	}

	return showResponse
}
