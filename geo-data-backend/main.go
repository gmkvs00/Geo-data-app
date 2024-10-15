package main

import (
    "geo-data-backend/config"
    "geo-data-backend/routes"
    "github.com/gin-gonic/gin"
    "net/http" // Import the net/http package
)

func main() {
    r := gin.Default()

    // CORS middleware
    r.Use(func(c *gin.Context) {
        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Methods", "POST, OPTIONS")
        c.Header("Access-Control-Allow-Headers", "Content-Type")
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(http.StatusNoContent)
            return
        }
        c.Next()
    })

    config.ConnectDatabase()

    // Set up routes for geospatial data
    routes.FileRoutes(r)

    r.Run(":8080") // Run on port 8080
}
