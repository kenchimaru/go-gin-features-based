package main

import (
    "your_project/config"
    "your_project/routes"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "log"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }
    config.ConnectDB()
    r := gin.Default()
    routes.RegisterRoutes(r)
    log.Println("🚀 Server running on port 8080")
    r.Run(":8080")
}
