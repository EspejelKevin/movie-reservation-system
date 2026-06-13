package showtime

import (
	"movie-reservation-system/src/domain/dto"
	"movie-reservation-system/src/domain/entities"
	"movie-reservation-system/src/domain/interfaces"
	"net/http"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type SaveShowTimeUseCase struct {
	repository interfaces.ShowTimeRepository
	validate   *validator.Validate
}

func NewSaveShowTimeUseCase(repository interfaces.ShowTimeRepository,
	validate *validator.Validate) *SaveShowTimeUseCase {
	return &SaveShowTimeUseCase{repository, validate}
}

func (usecase *SaveShowTimeUseCase) Execute(ctx *gin.Context) {
	var showDTO dto.ShowDTO

	if err := ctx.ShouldBindJSON(&showDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := usecase.validate.Struct(showDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	showDB, err := usecase.repository.GetShowByFolio(ctx.Request.Context(), showDTO.Folio)

	if err != gorm.ErrRecordNotFound && err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !reflect.DeepEqual(showDB, entities.ShowTime{}) {
		ctx.JSON(http.StatusConflict, gin.H{"error": "show ya existente"})
		return
	}

	newShow := entities.ShowTime{
		Folio:         showDTO.Folio,
		StartDate:     usecase.toTime(showDTO.StartDate),
		EndDate:       usecase.toTime(showDTO.EndDate),
		QuantitySeats: showDTO.QuantitySeats,
		MovieID:       showDTO.MovieID,
	}

	if err := usecase.repository.Save(ctx.Request.Context(), newShow); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "registro exitoso"})
}

func (usecase *SaveShowTimeUseCase) toTime(date string) time.Time {
	loc, _ := time.LoadLocation("America/Mexico_City")
	time, _ := time.ParseInLocation("2006-01-02 15:04:05", date, loc)
	return time.UTC()
}
