package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hendriam/movie-service/controller"
	frameworks "github.com/hendriam/movie-service/framework"
	"github.com/hendriam/movie-service/repository"
	"github.com/hendriam/movie-service/service"
)

func main() {
	db, err := frameworks.LoadDatabase()
	if err != nil {
		panic(err)
	}

	movieRepository := repository.NewMovieRepository(db)
	movieService := service.NewMovieService(movieRepository)
	movieController := controller.NewControllerMovie(movieService)

	host := "localhost"
	port := 8080
	url := fmt.Sprintf("%s:%d", host, port)

	fmt.Println("[APP] Started at => ", url)

	gin.SetMode(gin.ReleaseMode)
	route := gin.Default()

	route.POST("/create", movieController.Create)

	route.Run(url)
}
