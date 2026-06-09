package routes

import (
	"movie-reservation-system/src/application/usecases/movie"
	"movie-reservation-system/src/infrastructure/middlewares"

	"github.com/gin-gonic/gin"
)

type RouterMovie struct {
	getMovie    *movie.GetMovieUseCase
	getMovies   *movie.GetMoviesUseCase
	saveMovie   *movie.SaveMovieUseCase
	updateImage *movie.UpdateImageMovieUseCase
	updateMovie *movie.UpdateMovieUseCase
	deleteMovie *movie.DeleteMovieUseCase

	middlewares *middlewares.Middlewares
}

func NewRouterMovie(getMovie *movie.GetMovieUseCase, getMovies *movie.GetMoviesUseCase,
	saveMovie *movie.SaveMovieUseCase, updateImage *movie.UpdateImageMovieUseCase,
	updateMovie *movie.UpdateMovieUseCase, deleteMovie *movie.DeleteMovieUseCase,
	middlewares *middlewares.Middlewares) *RouterMovie {
	return &RouterMovie{getMovie, getMovies, saveMovie, updateImage, updateMovie, deleteMovie, middlewares}
}

func (router *RouterMovie) RegisterRoutesMovie(engine *gin.Engine) {
	movies := engine.Group("/api/v1")
	movies.Use(router.middlewares.Authentication())
	{
		movies.GET("/movies", router.middlewares.Authorization("ADMIN", "REGULAR_USER"), router.getMovies.Execute)
		movies.POST("/movies", router.middlewares.Authorization("ADMIN"), router.saveMovie.Execute)
		movies.PATCH("/movies/:id/image", router.middlewares.Authorization("ADMIN"), router.updateImage.Execute)
		movies.GET("/movies/:id", router.middlewares.Authorization("ADMIN"), router.getMovie.Execute)
		movies.DELETE("/movies/:id", router.middlewares.Authorization("ADMIN"), router.deleteMovie.Execute)
		movies.PUT("/movies/:id", router.middlewares.Authorization("ADMIN"), router.updateMovie.Execute)
	}
}
