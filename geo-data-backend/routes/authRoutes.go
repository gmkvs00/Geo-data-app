package routes

import (
    "geo-data-backend/controllers"
    "github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
    r.POST("/signup", controllers.SignUp)
    r.POST("/login", controllers.Login)
}
