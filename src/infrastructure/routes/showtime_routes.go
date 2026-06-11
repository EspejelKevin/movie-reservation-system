package routes

import (
	"movie-reservation-system/src/application/usecases/showtime"
	"movie-reservation-system/src/infrastructure/middlewares"

	"github.com/gin-gonic/gin"
)

type RouterShowTime struct {
	getShow    *showtime.GetShowTimeUseCase
	getShows   *showtime.GetShowTimesUseCase
	saveShow   *showtime.SaveShowTimeUseCase
	updateShow *showtime.UpdateShowTimeUseCase
	deleteShow *showtime.DeleteShowTimeUseCase

	middlewares *middlewares.Middlewares
}

func NewRouterShowTime(getShow *showtime.GetShowTimeUseCase, getShows *showtime.GetShowTimesUseCase,
	saveShow *showtime.SaveShowTimeUseCase, updateShow *showtime.UpdateShowTimeUseCase,
	deleteShow *showtime.DeleteShowTimeUseCase, middlewares *middlewares.Middlewares) *RouterShowTime {
	return &RouterShowTime{getShow, getShows, saveShow, updateShow, deleteShow, middlewares}
}

func (router *RouterShowTime) RegisterRoutesShowTime(engine *gin.Engine) {
	showtime := engine.Group("/api/v1")
	showtime.Use(router.middlewares.Authentication())
	{
		showtime.GET("/showtimes", router.middlewares.Authorization("ADMIN"), router.getShows.Execute)
		showtime.GET("/showtimes/:id", router.middlewares.Authorization("ADMIN"), router.getShow.Execute)
		showtime.POST("/showtimes", router.middlewares.Authorization("ADMIN"), router.saveShow.Execute)
		showtime.PUT("/showtimes/:id", router.middlewares.Authorization("ADMIN"), router.updateShow.Execute)
		showtime.DELETE("/showtimes/:id", router.middlewares.Authorization("ADMIN"), router.deleteShow.Execute)
	}
}
