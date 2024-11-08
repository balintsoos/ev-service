package routes

import (
	"ev-service/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/vehicles", controllers.CreateVehicle)
	r.GET("/vehicles", controllers.GetVehicles)
	r.GET("/vehicles/:id", controllers.GetVehicle)
	r.PUT("/vehicles/:id", controllers.UpdateVehicle)
	r.DELETE("/vehicles/:id", controllers.DeleteVehicle)
}
