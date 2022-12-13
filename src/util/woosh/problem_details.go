package woosh

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProblemDetails struct {
	Title    string `json:"title" example:"a error title" format:"string"`
	Status   int    `json:"status" example:"400" format:"int"`
	Detail   string `json:"detail" example:"a error detail" format:"string"`
	Instance string `json:"instance" example:"/api/to/path" format:"string"`
}

func NewProblemDetails(title string, status int, detail string, instance string) ProblemDetails {
	return ProblemDetails{
		title,
		status,
		detail,
		instance,
	}
}

func HandleException(detail string, ctx *gin.Context) {
	title := "服务器无法处理该请求"
	status := http.StatusBadRequest
	instance := ctx.Request.URL.Path

	problemDetails := NewProblemDetails(title, status, detail, instance)

	ctx.JSON(status, problemDetails)
}
