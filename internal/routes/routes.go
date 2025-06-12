package routes

import (
	"store/internal/middleware"
    "github.com/gin-gonic/gin"
    "store/internal/handlers"
)

func SetupRoutes(r *gin.Engine) {
    v1 := r.Group("/api/v1")
{
    v1.POST("/register", handlers.Register)
    v1.POST("/login", handlers.Login)

    protected := v1.Group("/")
    protected.Use(middleware.AuthMiddleware())
    {
        protected.GET("/me", handlers.Me)
    }
}

}
