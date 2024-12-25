package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/imamrpratama/mrt-schedules/common/response"
)


func main() {
	InitiateRouter()
}

func InitiateRouter() {
	// Corrected variable declarations and initialization
	router := gin.Default()
	api := router.Group("/v1/api")

	// Call the Initiate function from the station package
	station.Initiate(api)

	// Start the server on port 8080
	router.Run(":8080")
}
