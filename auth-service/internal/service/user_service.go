package service

import (
	"errors"
	"github.com/Jvls1/go-ecommerce/domain"
	"github.com/Jvls1/go-ecommerce/internal/repository"
	"github.com/google/uuid"
	"strings"
)

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository}
}

type userService struct {
	repo repository.UserRepository
}

type UserService interface {
	SaveUser(name, password, email string) (*domain.User, error)
	RetrieveUserById(userID uuid.UUID) (*domain.User, error)
	AddRoleToUser(userID, roleID uuid.UUID) error
}

func (us *userService) SaveUser(name, password, email string) (*domain.User, error) {
	err := validateUser(name, password, email)
	if err != nil {
		return nil, err
	}
	user, _ := domain.NewUser(name, password, email)
	storedUser, err := us.repo.StoreUser(user)
	if err != nil {
		return nil, err
	}
	return storedUser, nil
}

func validateUser(name, password, email string) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("name cannot be empty")
	}
	if len(name) > 255 {
		return errors.New("name must be less than or equal to 255 characters")
	}
	if strings.TrimSpace(email) == "" {
		return errors.New("email cannot be empty")
	}
	if len(email) > 255 {
		return errors.New("description must be less than or equal to 255 characters")
	}
	return nil
}

func (us *userService) RetrieveUserById(userID uuid.UUID) (*domain.User, error) {
	return us.repo.GetUserById(userID)
}

func (us *userService) AddRoleToUser(userID, roleID uuid.UUID) error {
	// TODO Validate if the role is not add already
	return us.repo.AddRoleToUser(userID, roleID)
}
