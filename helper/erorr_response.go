package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/hendriam/movie-service/model/web"
)

func ErrorResponse(c *gin.Context, httpCode int, err error) {
	c.JSON(httpCode, web.Response{
		Code:    httpCode,
		Message: err.Error(),
	})
}
