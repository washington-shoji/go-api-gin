package helpers

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type HttpResponse struct {
	Status interface{} `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

type Response struct {
	Status  int
	Message []string
	Error   []string
}

func WebResponse(status interface{}, data interface{}) *HttpResponse {
	res := &HttpResponse{
		Status: status,
		Data:   data,
	}
	return res
}

func WebResponseHandler(c *gin.Context, response Response) {
	if len(response.Message) > 0 {
		c.JSON(response.Status, map[string]interface{}{"message": strings.Join(response.Message, "; ")})
	} else if len(response.Error) > 0 {
		c.JSON(response.Status, map[string]interface{}{"error": strings.Join(response.Error, "; ")})
	}
}
