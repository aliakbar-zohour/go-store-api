package main

import (
    "store/config"
    "store/internal/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    config.ConnectDB()
    routes.SetupRoutes(r)

    r.Run(":8080")
}
