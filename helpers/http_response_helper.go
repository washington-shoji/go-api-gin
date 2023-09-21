package helpers

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type ResponseError struct {
	Status int
	Error  []string
}

func WebResponseError(c *gin.Context, response ResponseError) {
	c.JSON(response.Status, map[string]interface{}{"error": strings.Join(response.Error, "; ")})
}

type ResponseSuccess struct {
	Status int
	Data   interface{}
}

func WebResponseSuccessHandler(c *gin.Context, response ResponseSuccess) {
	c.JSON(response.Status, response.Data)
}
