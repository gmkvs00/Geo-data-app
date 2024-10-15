// geo-data-backend/controllers/fileController.go
package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func UploadGeoData(c *gin.Context) {
    // Retrieve the file from the request
    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to retrieve file"})
        return
    }

    // Save the file to a specified directory
    err = c.SaveUploadedFile(file, "./uploads/"+file.Filename)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}
