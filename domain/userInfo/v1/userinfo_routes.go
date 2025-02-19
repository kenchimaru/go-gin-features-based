package userInfo

import (
    "github.com/gin-gonic/gin"
    "your_project/middleware"
)

func UserInfoRoutes(r *gin.Engine) {
    repo := NewUserInfoRepository()
    controller := NewUserInfoController(repo)

    user := r.Group("/api/v1/user")
    user.Use(middleware.AuthMiddleware())
    {
        user.GET("/info", controller.GetUserInfo)
    }
}
