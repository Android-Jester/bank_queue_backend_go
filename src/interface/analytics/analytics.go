package analytics

import (
	http "net/http"

	"github.com/gin-gonic/gin"
)

func Routes(group *gin.RouterGroup) {
	group.GET("/", LoginRoute)
	group.GET("/print", LoginRoute)
}

func LoginRoute(context *gin.Context) {
	data := make(map[string]interface{})
	data["Hello"] = "hi"
	context.JSON(http.StatusOK, data)
}
