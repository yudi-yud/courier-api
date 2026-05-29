package services

import (
	"courier-api/config"
	"courier-api/models"
	"courier-api/repositories"
	"courier-api/utils"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(username, password string) (string, error)
	Register(username, password, role string) error
}

type authService struct {
	userRepo repositories.UserRepository
}

func NewAuthService() AuthService {
	return &authService{userRepo: repositories.NewUserRepository()}
}

func (s *authService) Login(username, password string) (string, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return "", errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *authService) Register(username, password, role string) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := &models.User{
		Username: username,
		Password: string(hashedPassword),
		Role:     role,
	}
	return s.userRepo.Create(user)
}

func InitAdmin() {
	repo := repositories.NewUserRepository()
	if _, err := repo.FindByUsername("admin"); err != nil {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
		admin := &models.User{
			Username: "admin",
			Password: string(hashedPassword),
			Role:     "admin",
		}
		config.DB.Create(admin)
	}
}
