package reservation

import (
	"movie-reservation-system/src/domain/interfaces"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DeleteReservationUseCase struct {
	reservationRepo interfaces.ReservationRepository
	showtimeRepo    interfaces.ShowTimeRepository
}

func NewDeleteReservationUseCase(reservationRepo interfaces.ReservationRepository,
	showtimeRepo interfaces.ShowTimeRepository) *DeleteReservationUseCase {
	return &DeleteReservationUseCase{reservationRepo, showtimeRepo}
}

func (usecase *DeleteReservationUseCase) Execute(ctx *gin.Context) {
	reservationID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "formato invalido en el param :id"})
		return
	}

	reservationDB, err := usecase.reservationRepo.GetReservationByID(ctx.Request.Context(), uint(reservationID))

	if err != gorm.ErrRecordNotFound && err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err == gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "reservation no existente"})
		return
	}

	_, err = usecase.reservationRepo.Delete(ctx.Request.Context(), uint(reservationID))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newSeats := reservationDB.ReservedSeats - *reservationDB.ShowTime.UnavailableSeats
	usecase.showtimeRepo.UpdateUnavailableSeats(ctx.Request.Context(), reservationDB.ShowTimeID, newSeats)

	ctx.JSON(http.StatusOK, gin.H{"message": "eliminacion exitosa"})
}
