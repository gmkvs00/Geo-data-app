package controllers

import (
    "geo-data-backend/config"
    "geo-data-backend/models"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    "net/http"
    "github.com/dgrijalva/jwt-go"
    "time"
)

func SignUp(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
    user.Password = string(hashedPassword)

    config.DB.Create(&user)
    c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func Login(c *gin.Context) {
    var input struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user models.User
    config.DB.Where("email = ?", input.Email).First(&user)

    if user.ID == 0 || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)) != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "userId": user.ID,
        "exp":    time.Now().Add(time.Hour * 24).Unix(),
    })

    tokenString, _ := token.SignedString([]byte("secret"))
    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
