package routes

import (
	"movie-reservation-system/src/application/usecases/reservation"
	"movie-reservation-system/src/infrastructure/middlewares"

	"github.com/gin-gonic/gin"
)

type RouterReservation struct {
	getReservations    *reservation.GetReservationsUseCase
	getReservationUser *reservation.GetReservationUserUseCase
	getReservation     *reservation.GetReservationUseCase
	saveReservation    *reservation.SaveReservationUseCase
	cancelReservation  *reservation.CancelReservationUseCase
	deleteReservation  *reservation.DeleteReservationUseCase

	middlewares *middlewares.Middlewares
}

func NewRouterReservation(getReservations *reservation.GetReservationsUseCase,
	getReservationUser *reservation.GetReservationUserUseCase,
	getReservation *reservation.GetReservationUseCase,
	saveReservation *reservation.SaveReservationUseCase,
	cancelReservation *reservation.CancelReservationUseCase,
	deleteReservation *reservation.DeleteReservationUseCase,
	middlewares *middlewares.Middlewares) *RouterReservation {
	return &RouterReservation{getReservations, getReservationUser, getReservation,
		saveReservation, cancelReservation, deleteReservation, middlewares}
}

func (router *RouterReservation) RegisterRoutesReservation(engine *gin.Engine) {
	reservation := engine.Group("/api/v1")
	reservation.Use(router.middlewares.Authentication())
	{
		reservation.GET("/me/reservations", router.middlewares.Authorization("ADMIN", "REGULAR_USER"),
			router.getReservationUser.Execute)
		reservation.GET("/reservations", router.middlewares.Authorization("ADMIN"),
			router.getReservations.Execute)
		reservation.GET("/reservations/:id", router.middlewares.Authorization("ADMIN"),
			router.getReservation.Execute)
		reservation.POST("/reservations", router.middlewares.Authorization("ADMIN", "REGULAR_USER"),
			router.saveReservation.Execute)
		reservation.PUT("/reservations/:id/cancel", router.middlewares.Authorization("ADMIN", "REGULAR_USER"),
			router.cancelReservation.Execute)
		reservation.DELETE("/reservations/:id", router.middlewares.Authorization("ADMIN"),
			router.deleteReservation.Execute)
	}
}
