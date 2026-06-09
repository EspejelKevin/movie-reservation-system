package routes

import (
	"movie-reservation-system/src/application/usecases/movie"
	"movie-reservation-system/src/infrastructure/middlewares"

	"github.com/gin-gonic/gin"
)

type RouterMovie struct {
	getMovies   *movie.GetMoviesUseCase
	saveMovie   *movie.SaveMovieUseCase
	updateImage *movie.UpdateImageMovieUseCase

	middlewares *middlewares.Middlewares
}

func NewRouterMovie(getMovies *movie.GetMoviesUseCase, saveMovie *movie.SaveMovieUseCase,
	updateImage *movie.UpdateImageMovieUseCase, middlewares *middlewares.Middlewares) *RouterMovie {
	return &RouterMovie{getMovies, saveMovie, updateImage, middlewares}
}

func (router *RouterMovie) RegisterRoutesMovie(engine *gin.Engine) {
	movies := engine.Group("/api/v1")
	movies.Use(router.middlewares.Authentication())
	{
		movies.GET("/movies", router.middlewares.Authorization("ADMIN", "REGULAR_USER"), router.getMovies.Execute)
		movies.POST("/movies", router.middlewares.Authorization("ADMIN"), router.saveMovie.Execute)
		movies.PATCH("/movies/:id/image", router.middlewares.Authorization("ADMIN"), router.updateImage.Execute)
	}
}
