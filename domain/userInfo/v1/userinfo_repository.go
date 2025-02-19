package userInfo

import "your_project/common/models"

type userInfoRepository struct{}

func NewUserInfoRepository() *userInfoRepository {
    return &userInfoRepository{}
}

func (repo *userInfoRepository) GetUserInfo() *models.User {
    return &models.User{Email: "test@example.com", CreatedAt: 1700000000}
}
