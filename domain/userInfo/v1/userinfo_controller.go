package userInfo

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type UserInfoController struct {
    repo *userInfoRepository
}

func NewUserInfoController(repo *userInfoRepository) *UserInfoController {
    return &UserInfoController{repo: repo}
}

func (ctr *UserInfoController) GetUserInfo(c *gin.Context) {
    user := ctr.repo.GetUserInfo()
    c.JSON(http.StatusOK, gin.H{
        "email":     user.Email,
        "createdAt": user.CreatedAt,
    })
}
