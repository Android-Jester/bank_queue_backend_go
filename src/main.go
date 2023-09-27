package main

import (
	"github.com/Android-Jester/bank_queue_backend_go/src/interface/analytics"
	"github.com/Android-Jester/bank_queue_backend_go/src/interface/client"
	"github.com/Android-Jester/bank_queue_backend_go/src/interface/server"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	// Basic Auth
	users := router.Group("/user")
	client.Routes(users)
	servers := router.Group("/server")
	server.Routes(servers)
	analytic := router.Group("/analytics")
	analytics.Routes(analytic)

	return router
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":3000")
}
