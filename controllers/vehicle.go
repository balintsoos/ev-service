package controllers

import (
	"context"
	"ev-service/db"
	"ev-service/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create a new vehicle
func CreateVehicle(c *gin.Context) {
	var vehicle models.ElectricVehicle
	if err := c.ShouldBindJSON(&vehicle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `INSERT INTO electric_vehicles (make, model, year, battery_capacity, range_km, price) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err := db.DB.QueryRow(context.Background(), query, vehicle.Make, vehicle.Model, vehicle.Year, vehicle.BatteryCapacity, vehicle.RangeKm, vehicle.Price).Scan(&vehicle.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create vehicle"})
		return
	}
	c.JSON(http.StatusCreated, vehicle)
}

// Retrieve all vehicles
func GetVehicles(c *gin.Context) {
	rows, err := db.DB.Query(context.Background(), `SELECT id, make, model, year, battery_capacity, range_km, price FROM electric_vehicles`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch vehicles"})
		return
	}
	defer rows.Close()

	var vehicles []models.ElectricVehicle
	for rows.Next() {
		var vehicle models.ElectricVehicle
		if err := rows.Scan(&vehicle.ID, &vehicle.Make, &vehicle.Model, &vehicle.Year, &vehicle.BatteryCapacity, &vehicle.RangeKm, &vehicle.Price); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning vehicle data"})
			return
		}
		vehicles = append(vehicles, vehicle)
	}
	c.JSON(http.StatusOK, vehicles)
}

// Retrieve a single vehicle by ID
func GetVehicle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid vehicle ID"})
		return
	}

	var vehicle models.ElectricVehicle
	query := `SELECT id, make, model, year, battery_capacity, range_km, price FROM electric_vehicles WHERE id = $1`
	err = db.DB.QueryRow(context.Background(), query, id).Scan(&vehicle.ID, &vehicle.Make, &vehicle.Model, &vehicle.Year, &vehicle.BatteryCapacity, &vehicle.RangeKm, &vehicle.Price)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vehicle not found"})
		return
	}
	c.JSON(http.StatusOK, vehicle)
}

// Update a vehicle by ID
func UpdateVehicle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid vehicle ID"})
		return
	}

	var vehicle models.ElectricVehicle
	if err := c.ShouldBindJSON(&vehicle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	vehicle.ID = id

	query := `UPDATE electric_vehicles SET make=$1, model=$2, year=$3, battery_capacity=$4, range_km=$5, price=$6 WHERE id=$7`
	_, err = db.DB.Exec(context.Background(), query, vehicle.Make, vehicle.Model, vehicle.Year, vehicle.BatteryCapacity, vehicle.RangeKm, vehicle.Price, vehicle.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update vehicle"})
		return
	}
	c.JSON(http.StatusOK, vehicle)
}

// Delete a vehicle by ID
func DeleteVehicle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid vehicle ID"})
		return
	}

	_, err = db.DB.Exec(context.Background(), `DELETE FROM electric_vehicles WHERE id = $1`, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete vehicle"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
