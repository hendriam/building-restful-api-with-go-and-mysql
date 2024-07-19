package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hendriam/movie-service/controller"
	"github.com/hendriam/movie-service/framework"
	"github.com/hendriam/movie-service/repository"
	"github.com/hendriam/movie-service/service"
)

func main() {
	cfg := framework.LoadConfig()
	logging := framework.LoadLogging()
	db, err := framework.LoadDatabase()
	if err != nil {
		panic(err)
	}

	movieRepository := repository.NewMovieRepository(db)
	movieService := service.NewMovieService(movieRepository)
	movieController := controller.NewControllerMovie(movieService, logging)

	url := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)

	logging.Info().Msgf("starting web server at http://%s/", url)

	gin.SetMode(gin.ReleaseMode)
	route := gin.Default()

	route.POST("movie/create", movieController.Create)
	route.PUT("movie/update/:id", movieController.Update)
	route.DELETE("movie/delete/:id", movieController.Delete)
	route.GET("movie/:id", movieController.FindById)
	route.GET("movie", movieController.FindAll)

	route.Run(url)
}
