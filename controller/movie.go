package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	api "github.com/hendriam/movie-service/model/web"
	"github.com/hendriam/movie-service/service"
)

type MovieController struct {
	movieService service.MovieService
}

func NewControllerMovie(movieService service.MovieService) MovieController {
	return MovieController{movieService: movieService}
}

func (controller *MovieController) Create(c *gin.Context) {
	movieBody := api.MovieCreateRequestBody{}
	if err := c.ShouldBindJSON(&movieBody); err != nil {
		fmt.Println("[CONTROLLER] error =>", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	movieResponse, err := controller.movieService.CreateMovie(c.Request.Context(), movieBody)
	if err != nil {
		fmt.Println("[CONTROLLER] error =>", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, movieResponse)
}
