package routes

import (
	"movie-reservation-system/src/application/usecases/showtime"
	"movie-reservation-system/src/infrastructure/middlewares"

	"github.com/gin-gonic/gin"
)

type RouterShowTime struct {
	getShows     *showtime.GetShowTImesUseCase
	saveShowTime *showtime.SaveShowTimeUseCase

	middlewares *middlewares.Middlewares
}

func NewRouterShowTime(getShows *showtime.GetShowTImesUseCase, saveShowTime *showtime.SaveShowTimeUseCase,
	middlewares *middlewares.Middlewares) *RouterShowTime {
	return &RouterShowTime{getShows, saveShowTime, middlewares}
}

func (router *RouterShowTime) RegisterRoutesShowTime(engine *gin.Engine) {
	showtime := engine.Group("/api/v1")
	showtime.Use(router.middlewares.Authentication())
	{
		showtime.GET("/showtimes", router.middlewares.Authorization("ADMIN"), router.getShows.Execute)
		showtime.POST("/showtimes", router.middlewares.Authorization("ADMIN"), router.saveShowTime.Execute)
	}
}
