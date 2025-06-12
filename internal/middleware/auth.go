package middleware

import (
    "net/http"
    "store/pkg"
    "store/internal/repository"
    "context"
    "github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
            c.Abort()
            return
        }

        // Expected format: "Bearer <token>"
        const prefix = "Bearer "
        if len(authHeader) < len(prefix) || authHeader[:len(prefix)] != prefix {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
            c.Abort()
            return
        }

        token := authHeader[len(prefix):]

        email, err := pkg.ParseJWT(token)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        user, err := repository.FindUserByEmail(context.TODO(), email)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
            c.Abort()
            return
        }

        c.Set("user", user)
        c.Next()
    }
}
