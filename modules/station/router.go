package station

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"your_project_path/response" // Update the import path
)

func Initiate(router *gin.RouterGroup) {
	stationService := NewService() // Initialize your service here.

	station := router.Group("/stations")
	station.GET("", func(c *gin.Context) {
		GetAllStation(c, stationService)
	})
}

func GetAllStation(c *gin.Context, service Service) {
	datas, err := service.GetAllStation()
	if err != nil {
		// Handle the error
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, response.APIResponse{
		Success: true,
		Message: "Successfully fetched all stations",
		Data:    datas,
	})
}
