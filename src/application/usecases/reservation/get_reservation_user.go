package reservation

import (
	"movie-reservation-system/src/domain/dto"
	"movie-reservation-system/src/domain/entities"
	"movie-reservation-system/src/domain/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetReservationUserUseCase struct {
	repository interfaces.ReservationRepository
}

func NewGetReservationUserUseCase(repository interfaces.ReservationRepository) *GetReservationUserUseCase {
	return &GetReservationUserUseCase{repository}
}

func (usecase *GetReservationUserUseCase) Execute(ctx *gin.Context) {
	id, exists := ctx.Get("userID")

	if !exists {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "no se puede determinar el userID"})
		return
	}

	userID, ok := id.(uint64)

	if !ok {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "valor incorrecto para el campo userID"})
		return
	}

	reservations, err := usecase.repository.GetReservationsByUserID(ctx.Request.Context(), uint(userID))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"reservations": usecase.mapReservations(reservations)})
}

func (usecase *GetReservationUserUseCase) mapReservations(reservations []entities.Reservation) []dto.ReservationResponse {
	reservationResponse := make([]dto.ReservationResponse, 0, len(reservations))

	for _, reservation := range reservations {
		reservationResponse = append(reservationResponse, dto.ReservationResponse{
			ID:     reservation.ID,
			Folio:  reservation.Folio,
			Status: reservation.Status,
			ShowTime: dto.ShowResponse{
				ID:                     reservation.ShowTime.ID,
				Folio:                  reservation.ShowTime.Folio,
				StartDate:              reservation.ShowTime.StartDate.Format("2006-01-02 15:04:05"),
				EndDate:                reservation.ShowTime.EndDate.Format("2006-01-02 15:04:05"),
				AvailableQuantitySeats: reservation.ShowTime.QuantitySeats - *reservation.ShowTime.UnavailableSeats,
				UnavailableSeats:       reservation.ShowTime.UnavailableSeats,
				Movie: dto.MovieResponse{
					ID:          reservation.ShowTime.Movie.ID,
					Title:       reservation.ShowTime.Movie.Title,
					Description: reservation.ShowTime.Movie.Description,
					Image:       reservation.ShowTime.Movie.Image.String,
					Genre:       reservation.ShowTime.Movie.Genre.Name,
				},
			},
			User: dto.UserResponse{
				Name:     reservation.User.Name,
				UserName: reservation.User.Email,
				Email:    reservation.User.Email,
			},
		})
	}

	return reservationResponse
}
