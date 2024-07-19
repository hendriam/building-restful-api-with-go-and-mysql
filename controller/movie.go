package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hendriam/movie-service/framework"
	"github.com/hendriam/movie-service/helper"
	"github.com/hendriam/movie-service/model/web"
	"github.com/hendriam/movie-service/service"
)

type MovieController struct {
	movieService service.MovieService
	Logging      framework.Logging
}

func NewControllerMovie(movieService service.MovieService, logging framework.Logging) MovieController {
	return MovieController{movieService: movieService, Logging: logging}
}

func (controller *MovieController) Create(c *gin.Context) {
	requestBody := web.MovieCreateRequestBody{}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		controller.Logging.Error().Msgf("[CONTROLLER][CREATE] error => %s", err.Error())
		helper.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	createResponse, err := controller.movieService.Create(c.Request.Context(), requestBody)
	if err != nil {
		controller.Logging.Error().Msgf("[CONTROLLER][CREATE] error => %s", err.Error())
		helper.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	controller.Logging.Info().Msgf("[CONTROLLER][CREATE] success => %v", createResponse)

	helper.SuccessResponse(c, http.StatusCreated, createResponse)
}

func (controller *MovieController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		controller.Logging.Error().Msgf("[CONTROLLER][UPDATE] error => %s", err.Error())
		helper.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	requestBody := web.MovieUpdateRequestBody{}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		controller.Logging.Error().Msgf("[CONTROLLER][UPDATE] error => %s", err.Error())
		helper.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	// This step is a function to check whether updated data exists or not.
	ReadById, err := controller.movieService.ReadById(c.Request.Context(), id)
	if err != nil {
		controller.Logging.Error().Msgf("[CONTROLLER][FIND-BY-ID] error => %s", err.Error())
		helper.ErrorResponse(c, http.StatusNotFound, err)
		return
	}

	updateResponse, err := controller.movieService.Update(c.Request.Context(), requestBody, ReadById.ID)
	if err != nil {
		controller.Logging.Error().Msgf("[CONTROLLER][UPDATE] error => %s", err.Error())
		helper.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	controller.Logging.Info().Msgf("[CONTROLLER][UPDATE] success => %v", updateResponse)

	helper.SuccessResponse(c, http.StatusOK, updateResponse)
}

func (controller *MovieController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		controller.Logging.Error().Msgf("[CONTROLLER][DELETE] error => %s", err.Error())
		helper.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	// This step is a function to check whether delete data exists or not.
	ReadById, err := controller.movieService.ReadById(c.Request.Context(), id)
	if err != nil {
		controller.Logging.Error().Msgf("[CONTROLLER][FIND-BY-ID] error => %s", err.Error())
		helper.ErrorResponse(c, http.StatusNotFound, err)
		return
	}

	err = controller.movieService.Delete(c.Request.Context(), ReadById.ID)
	if err != nil {
		controller.Logging.Error().Msgf("[CONTROLLER][DELETE] error => %s", err.Error())
		helper.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	controller.Logging.Info().Msgf("[CONTROLLER][DELETE] success => %v", ReadById)

	helper.SuccessResponse(c, http.StatusOK, ReadById)
}

func (controller *MovieController) FindById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		controller.Logging.Error().Msgf("[CONTROLLER][FIND-BY-ID] error => %s", err.Error())
		helper.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	// This step is a function to check whether delete data exists or not.
	ReadById, err := controller.movieService.ReadById(c.Request.Context(), id)
	if err != nil {
		controller.Logging.Error().Msgf("[CONTROLLER][FIND-BY-ID] error => %s", err.Error())
		helper.ErrorResponse(c, http.StatusNotFound, err)
		return
	}

	controller.Logging.Info().Msgf("[CONTROLLER][FIND-BY-ID] success => %v", ReadById)

	helper.SuccessResponse(c, http.StatusOK, ReadById)
}

func (controller *MovieController) FindAll(c *gin.Context) {
	readAll, err := controller.movieService.ReadAll(c.Request.Context())
	if err != nil {
		controller.Logging.Error().Msgf("[CONTROLLER][FIND-ALL] error => %s", err.Error())
		helper.ErrorResponse(c, http.StatusNotFound, err)
		return
	}

	controller.Logging.Info().Msgf("[CONTROLLER][FIND-ALL] success => %v", readAll)

	helper.SuccessResponse(c, http.StatusOK, readAll)
}
