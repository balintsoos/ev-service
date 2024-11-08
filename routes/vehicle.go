package routes

import (
	"ev-service/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/api/v0/vehicles", controllers.CreateVehicle)
	r.GET("/api/v0/vehicles", controllers.GetVehicles)
	r.GET("/api/v0/vehicles/:id", controllers.GetVehicle)
	r.PUT("/api/v0/vehicles/:id", controllers.UpdateVehicle)
	r.DELETE("/api/v0/vehicles/:id", controllers.DeleteVehicle)
}
