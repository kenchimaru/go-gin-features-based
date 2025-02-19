package auth

import (
    "errors"
    "your_project/utils"
    "your_project/common/models"
    "your_project/middleware"
)

type authService struct {
    repo *authRepository
}

func NewAuthService(repo *authRepository) *authService {
    return &authService{repo: repo}
}

// Register new user
func (s *authService) RegisterUser(email, password string) error {
    existingUser, _ := s.repo.FindUserByEmail(email)
    if existingUser != nil {
        return errors.New("user already exists")
    }

    hashedPassword, _ := utils.HashPassword(password)
    newUser := &models.User{Email: email, Password: hashedPassword}

    _, err := s.repo.SaveUser(newUser)
    return err
}

// Authenticate user and return JWT token
func (s *authService) LoginUser(email, password string) (string, error) {
    user, err := s.repo.FindUserByEmail(email)
    if err != nil {
        return "", errors.New("user not found")
    }

    if !utils.CheckPasswordHash(password, user.Password) {
        return "", errors.New("incorrect password")
    }

    // Generate JWT token
    token, err := middleware.GenerateToken(user.Email)
    if err != nil {
        return "", err
    }

    return token, nil
}
