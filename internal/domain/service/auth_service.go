package service

import (
	"errors"
	"time"

	"sue_backend/internal/common/utils"
	"sue_backend/internal/domain/model"
	"sue_backend/internal/domain/repository"
	"sue_backend/internal/infra/auth"
	JWTManager "sue_backend/internal/infra/auth"
)

type AuthService struct {
	jwtManager *JWTManager.JWTManager
	userRepo   *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository, jwt *auth.JWTManager) *AuthService {
	return &AuthService{
		userRepo:   repo,
		jwtManager: jwt,
	}
}
func (a *AuthService) RegisterUser(user *model.User) (string, error) {
	existing, _ := a.userRepo.FindByEmail(*user.Email)
	if existing != nil {
		return "", errors.New("email already exists")
	}

	salt := utils.GenerateSalt()
	user.Salt = &salt
	hashed := utils.HashPassword(*user.Password, salt)
	user.Password = &hashed
	user.Status = "active"
	now := time.Now()
	user.CreatedAt, user.UpdatedAt = now, now

	if err := a.userRepo.Create(user); err != nil {
		return "", err
	}

	token, err := a.jwtManager.Generate(user.ID, string(user.Role))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (a *AuthService) Login(email, password string) (string, error) {
	user, err := a.userRepo.FindByEmail(email)
	if err != nil || user == nil {
		return "", errors.New("invalid email or password")
	}
	if utils.HashPassword(password, *user.Salt) != *user.Password {
		return "", errors.New("invalid email or password")
	}
	token, err := a.jwtManager.Generate(user.ID, string(user.Role))
	if err != nil {
		return "", err
	}
	return token, nil
}
