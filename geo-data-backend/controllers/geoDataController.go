package controllers

import (
    "geo-data-backend/config"
    "geo-data-backend/models"
    "github.com/gin-gonic/gin"
    "net/http"
)

// UploadGeoDataController handles uploading GeoJSON or KML files
func UploadGeoDataController(c *gin.Context) {
    var geoData models.GeoData
    if err := c.ShouldBindJSON(&geoData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    config.DB.Create(&geoData)
    c.JSON(http.StatusOK, gin.H{"message": "GeoData uploaded successfully"})
}
