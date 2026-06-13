package routes

import (
	"movie-reservation-system/src/application/usecases/reservation"
	"movie-reservation-system/src/infrastructure/middlewares"

	"github.com/gin-gonic/gin"
)

type RouterReservation struct {
	getReservationUser *reservation.GetReservationUserUseCase
	getReservation     *reservation.GetReservationUseCase
	saveReservation    *reservation.SaveReservationUseCase
	cancelReservation  *reservation.CancelReservationUseCase

	middlewares *middlewares.Middlewares
}

func NewRouterReservation(getReservationUser *reservation.GetReservationUserUseCase,
	getReservation *reservation.GetReservationUseCase,
	saveReservation *reservation.SaveReservationUseCase,
	cancelReservation *reservation.CancelReservationUseCase,
	middlewares *middlewares.Middlewares) *RouterReservation {
	return &RouterReservation{getReservationUser, getReservation, saveReservation, cancelReservation, middlewares}
}

func (router *RouterReservation) RegisterRoutesReservation(engine *gin.Engine) {
	reservation := engine.Group("/api/v1")
	reservation.Use(router.middlewares.Authentication())
	{
		reservation.GET("/me/reservations", router.middlewares.Authorization("ADMIN", "REGULAR_USER"),
			router.getReservationUser.Execute)
		reservation.GET("/reservations/:id", router.middlewares.Authorization("ADMIN"),
			router.getReservation.Execute)
		reservation.POST("/reservations", router.middlewares.Authorization("ADMIN", "REGULAR_USER"),
			router.saveReservation.Execute)
		reservation.PUT("/reservations/:id/cancel", router.middlewares.Authorization("ADMIN", "REGULAR_USER"),
			router.cancelReservation.Execute)
	}
}
