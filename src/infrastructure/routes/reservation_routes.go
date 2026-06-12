package routes

import (
	"movie-reservation-system/src/application/usecases/reservation"
	"movie-reservation-system/src/infrastructure/middlewares"

	"github.com/gin-gonic/gin"
)

type RouterReservation struct {
	getReservation  *reservation.GetReservationUseCase
	saveReservation *reservation.SaveReservationUseCase

	middlewares *middlewares.Middlewares
}

func NewRouterReservation(getReservation *reservation.GetReservationUseCase,
	saveReservation *reservation.SaveReservationUseCase,
	middlewares *middlewares.Middlewares) *RouterReservation {
	return &RouterReservation{getReservation, saveReservation, middlewares}
}

func (router *RouterReservation) RegisterRoutesReservation(engine *gin.Engine) {
	reservation := engine.Group("/api/v1")
	reservation.Use(router.middlewares.Authentication())
	{
		reservation.GET("/reservation/:id", router.middlewares.Authorization("ADMIN"),
			router.getReservation.Execute)
		reservation.POST("/reservation", router.middlewares.Authorization("ADMIN", "REGULAR_USER"),
			router.saveReservation.Execute)
	}
}
