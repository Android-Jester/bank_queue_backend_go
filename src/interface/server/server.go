package server

import (
	http "net/http"

	"github.com/gin-gonic/gin"
)

func Routes(group *gin.RouterGroup) {
	group.GET("/login", LoginRoute)
	group.GET("/dismiss", LoginRoute)
	group.GET("/complete", LoginRoute)

}

func LoginRoute(context *gin.Context) {
	data := make(map[string]interface{})
	data["Hello"] = "hi"
	context.JSON(http.StatusOK, data)
}
