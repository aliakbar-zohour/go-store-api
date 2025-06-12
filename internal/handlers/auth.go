package handlers

import (
    "context"
    "log"
    "net/http"
    "store/internal/models"
    "store/internal/repository"
    "store/pkg"

    "github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
    var user models.User

    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    existingUser, err := repository.FindUserByEmail(context.TODO(), user.Email)
    if err == nil && existingUser.Email != "" {
        c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
        return
    }

    hashed, err := pkg.HashPassword(user.Password)
    if err != nil {
        log.Println("Hashing password failed:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
        return
    }

    user.Password = hashed
    user.Role = "user"

    if err := repository.CreateUser(context.TODO(), &user); err != nil {
        log.Println("User creation failed:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Registration failed"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Registered successfully"})
}

func Login(c *gin.Context) {
    var input struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    user, err := repository.FindUserByEmail(context.TODO(), input.Email)
    if err != nil || user.Email == "" {
        log.Println("User not found:", err)
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    if !pkg.CheckPasswordHash(input.Password, user.Password) {
        log.Println("Password mismatch for:", input.Email)
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    token, err := pkg.GenerateJWT(user.Email)
    if err != nil {
        log.Println("JWT generation failed:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Login successful",
        "token":   token,
        "user": gin.H{
            "id":    user.ID.Hex(),
            "email": user.Email,
            "role":  user.Role,
        },
    })
}


func Me(c *gin.Context) {
    userRaw, exists := c.Get("user")
    if !exists {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "User context missing"})
        return
    }

    user := userRaw.(*models.User)
    c.JSON(http.StatusOK, gin.H{"user": user})
}
