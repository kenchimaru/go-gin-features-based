package auth

import (
    "github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
    repo := NewAuthRepository()
    service := NewAuthService(repo)
    controller := NewAuthController(service)

    auth := r.Group("/api/v1/auth")
    {
        auth.POST("/register", controller.Register)
        auth.POST("/login", controller.Login)
    }
}
