package reservation

import (
	"movie-reservation-system/src/domain/entities"
	"movie-reservation-system/src/domain/interfaces"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CancelReservationUseCase struct {
	repository interfaces.ReservationRepository
}

func NewCancelReservationUseCase(repository interfaces.ReservationRepository) *CancelReservationUseCase {
	return &CancelReservationUseCase{repository}
}

func (usecase *CancelReservationUseCase) Execute(ctx *gin.Context) {
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

	reservationID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "formato invalido en el param :id"})
		return
	}

	reservationDB, err := usecase.repository.GetReservationByIDAndUserID(ctx.Request.Context(),
		uint(reservationID), uint(userID))

	if err != gorm.ErrRecordNotFound && err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err == gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "reservation no existente"})
		return
	}

	if reservationDB.Status == "CANCELED" {
		ctx.JSON(http.StatusConflict, gin.H{"error": "reservation ya cancelada previamente"})
		return
	}

	if !usecase.canCancel(reservationDB) {
		ctx.JSON(http.StatusConflict, gin.H{
			"error": "las reservaciones solo pueden cancelarse con al menos 24 horas de anticipacion a la funcion"})
		return
	}

	_, err = usecase.repository.UpdateReservationStatus(ctx.Request.Context(), uint(reservationID), uint(userID), "CANCELED")

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "reservation cancelada"})
}

func (usecase *CancelReservationUseCase) canCancel(reservation entities.Reservation) bool {
	loc, _ := time.LoadLocation("America/Mexico_City")
	eventStart := reservation.ShowTime.StartDate.In(loc)
	now := time.Now().In(loc)
	deadline := eventStart.Add(-24 * time.Hour)
	return now.Before(deadline)
}
