package main

import (
	"context"
	"ev-service/db"
	"ev-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDatabase()
	defer db.CloseDatabase(context.Background())

	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
