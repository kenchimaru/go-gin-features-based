package routes

import (
    "github.com/gin-gonic/gin"
    "your_project/domain/auth/v1"
    "your_project/domain/userInfo/v1"
    "your_project/domain/health/v1"
)

func RegisterRoutes(r *gin.Engine) {
    auth.AuthRoutes(r)
    userInfo.UserInfoRoutes(r)
    health.HealthRoutes(r)
}
