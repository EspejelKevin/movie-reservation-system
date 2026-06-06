package routes

import (
	"movie-reservation-system/src/application/usecases/genre"
	"movie-reservation-system/src/infrastructure/middlewares"

	"github.com/gin-gonic/gin"
)

type RouterGenre struct {
	getGenres   *genre.GetGenresUseCase
	getGenre    *genre.GetGenreUseCase
	saveGenre   *genre.SaveGenreUseCase
	updateGenre *genre.UpdateGenreUseCase
	deleteGenre *genre.DeleteGenreUseCase

	middlewares *middlewares.Middlewares
}

func NewRouterGenre(getGenres *genre.GetGenresUseCase, getGenre *genre.GetGenreUseCase,
	saveGenre *genre.SaveGenreUseCase, updateGenre *genre.UpdateGenreUseCase,
	deleteGenre *genre.DeleteGenreUseCase, middlewares *middlewares.Middlewares) *RouterGenre {
	return &RouterGenre{getGenres, getGenre, saveGenre, updateGenre, deleteGenre, middlewares}
}

func (router *RouterGenre) RegisterRoutesGenre(engine *gin.Engine) {
	genres := engine.Group("/api/v1")
	genres.Use(router.middlewares.Authentication())
	{
		genres.GET("/genres", router.middlewares.Authorization("ADMIN"), router.getGenres.Execute)
		genres.GET("/genres/:id", router.middlewares.Authorization("ADMIN"), router.getGenre.Execute)
		genres.POST("/genres", router.middlewares.Authorization("ADMIN"), router.saveGenre.Execute)
		genres.PUT("/genres/:id", router.middlewares.Authorization("ADMIN"), router.updateGenre.Execute)
		genres.DELETE("/genres/:id", router.middlewares.Authorization("ADMIN"), router.deleteGenre.Execute)
	}
}
