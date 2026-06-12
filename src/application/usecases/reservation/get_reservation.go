package reservation

import (
	"movie-reservation-system/src/domain/dto"
	"movie-reservation-system/src/domain/interfaces"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GetReservationUseCase struct {
	repository interfaces.ReservationRepository
}

func NewGetReservationUseCase(repository interfaces.ReservationRepository) *GetReservationUseCase {
	return &GetReservationUseCase{repository}
}

func (usecase *GetReservationUseCase) Execute(ctx *gin.Context) {
	reservationID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "formato invalido en el param :id"})
		return
	}

	reservationDB, err := usecase.repository.GetReservationByID(ctx.Request.Context(), uint(reservationID))

	if err != gorm.ErrRecordNotFound && err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err == gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "reservation no existente"})
		return
	}

	ctx.JSON(http.StatusOK, dto.ReservationResponse{
		ID:     reservationDB.ID,
		Folio:  reservationDB.Folio,
		Status: reservationDB.Status,
		ShowTime: dto.ShowResponse{
			ID:                     reservationDB.ShowTime.ID,
			Folio:                  reservationDB.ShowTime.Folio,
			StartDate:              reservationDB.ShowTime.StartDate.Format("2006-01-02 15:04:05"),
			EndDate:                reservationDB.ShowTime.EndDate.Format("2006-01-02 15:04:05"),
			AvailableQuantitySeats: reservationDB.ShowTime.QuantitySeats - *reservationDB.ShowTime.UnavailableSeats,
			UnavailableSeats:       reservationDB.ShowTime.UnavailableSeats,
			Movie: dto.MovieResponse{
				ID:          reservationDB.ShowTime.Movie.ID,
				Title:       reservationDB.ShowTime.Movie.Title,
				Description: reservationDB.ShowTime.Movie.Description,
				Image:       reservationDB.ShowTime.Movie.Image.String,
				Genre:       reservationDB.ShowTime.Movie.Genre.Name,
			},
		},
		User: dto.UserResponse{
			Name:     reservationDB.User.Name,
			UserName: reservationDB.User.Email,
			Email:    reservationDB.User.Email,
		},
	})
}
