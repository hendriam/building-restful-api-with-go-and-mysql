package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hendriam/movie-service/helper"
	"github.com/hendriam/movie-service/model/web"
	"github.com/hendriam/movie-service/service"
)

type MovieController struct {
	movieService service.MovieService
}

func NewControllerMovie(movieService service.MovieService) MovieController {
	return MovieController{movieService: movieService}
}

func (controller *MovieController) Create(c *gin.Context) {
	requestBody := web.MovieCreateRequestBody{}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		fmt.Println("[CONTROLLER][CREATE] error =>", err)
		helper.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	createResponse, err := controller.movieService.Create(c.Request.Context(), requestBody)
	if err != nil {
		fmt.Println("[CONTROLLER][CREATE] error =>", err)
		helper.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	helper.SuccessResponse(c, http.StatusCreated, createResponse)
}

func (controller *MovieController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("[CONTROLLER][UPDATE] error =>", err)
		helper.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	requestBody := web.MovieUpdateRequestBody{}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		fmt.Println("[CONTROLLER][UPDATE] error =>", err)
		helper.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	// This step is a function to check whether updated data exists or not.
	findById, err := controller.movieService.FindById(c.Request.Context(), id)
	if err != nil {
		fmt.Println("[CONTROLLER][FIND-BY-ID] error =>", err)
		helper.ErrorResponse(c, http.StatusNotFound, err)
		return
	}

	updateResponse, err := controller.movieService.Update(c.Request.Context(), requestBody, findById.ID)
	if err != nil {
		fmt.Println("[CONTROLLER][UPDATE] error =>", err)
		helper.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	helper.SuccessResponse(c, http.StatusOK, updateResponse)
}

func (controller *MovieController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("[CONTROLLER][DELETE] error =>", err)
		helper.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	// This step is a function to check whether delete data exists or not.
	findById, err := controller.movieService.FindById(c.Request.Context(), id)
	if err != nil {
		fmt.Println("[CONTROLLER][FIND-BY-ID] error =>", err)
		helper.ErrorResponse(c, http.StatusNotFound, err)
		return
	}

	err = controller.movieService.Delete(c.Request.Context(), findById.ID)
	if err != nil {
		fmt.Println("[CONTROLLER][DELETE] error =>", err)
		helper.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	helper.SuccessResponse(c, http.StatusOK, findById)
}
