package reservation

import (
	"movie-reservation-system/src/domain/dto"
	"movie-reservation-system/src/domain/entities"
	"movie-reservation-system/src/domain/interfaces"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type SaveReservationUseCase struct {
	reservationRepo interfaces.ReservationRepository
	showtimeRepo    interfaces.ShowTimeRepository
	validate        *validator.Validate
}

func NewSaveReservationUseCase(reservationRepo interfaces.ReservationRepository,
	showtimeRepo interfaces.ShowTimeRepository, validate *validator.Validate) *SaveReservationUseCase {
	return &SaveReservationUseCase{reservationRepo, showtimeRepo, validate}
}

func (usecase *SaveReservationUseCase) Execute(ctx *gin.Context) {
	var reservationDTO dto.ReservationDTO

	if err := ctx.ShouldBindJSON(&reservationDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := usecase.validate.Struct(reservationDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reservationDB, err := usecase.reservationRepo.GetReservationByFolio(ctx.Request.Context(), reservationDTO.Folio)

	if err != gorm.ErrRecordNotFound && err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !reflect.DeepEqual(reservationDB, entities.Reservation{}) {
		ctx.JSON(http.StatusConflict, gin.H{"error": "reservation ya existente"})
		return
	}

	showDB, err := usecase.showtimeRepo.GetShowByID(ctx.Request.Context(), reservationDTO.ShowTimeID)

	if err != gorm.ErrRecordNotFound && err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err == gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "showtime no existente"})
		return
	}

	availableQuantitySeats := showDB.QuantitySeats - *showDB.UnavailableSeats

	if reservationDTO.ReservedSeats > availableQuantitySeats {
		ctx.JSON(http.StatusConflict, gin.H{
			"error": "cantidad de asientos no disponible para reservar",
			"seats": availableQuantitySeats})
		return
	}

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

	newReservation := entities.Reservation{
		Folio:      reservationDTO.Folio,
		ShowTimeID: reservationDTO.ShowTimeID,
		UserID:     uint(userID),
	}

	if err := usecase.reservationRepo.Save(ctx.Request.Context(), newReservation); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	unavailableSeats := reservationDTO.ReservedSeats + *showDB.UnavailableSeats

	_, err = usecase.showtimeRepo.UpdateUnavailableSeats(ctx, reservationDTO.ShowTimeID,
		unavailableSeats)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "registro exitoso"})
}
