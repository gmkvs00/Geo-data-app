package routes

import (
    "geo-data-backend/controllers"
    "github.com/gin-gonic/gin"
)

// GeoDataRoutes sets up the routes for geospatial data
func GeoDataRoutes(r *gin.Engine) {
    r.POST("/upload", controllers.UploadGeoDataController)
}
