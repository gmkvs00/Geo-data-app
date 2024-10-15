// geo-data-backend/routes/fileRoutes.go
package routes

import (
    "geo-data-backend/controllers"
    "github.com/gin-gonic/gin"
)

func FileRoutes(router *gin.Engine) {
    router.POST("/upload", controllers.UploadGeoData)
}
