package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/hendriam/movie-service/model/web"
)

func SuccessResponse(c *gin.Context, httpCode int, movie any) {
	c.JSON(httpCode, web.Response{
		Code:    httpCode,
		Message: "Oke",
		Data:    movie,
	})
}
