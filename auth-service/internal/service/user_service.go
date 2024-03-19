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
	SaveUser(user *domain.User) (*domain.User, error)
	RetrieveUserById(userID uuid.UUID) (*domain.User, error)
	AddRoleToUser(userID, roleID uuid.UUID) error
}

func (us *userService) SaveUser(user *domain.User) (*domain.User, error) {
	err := validateUser(user)
	if err != nil {
		return nil, err
	}
	storedUser, err := us.repo.StoreUser(user)
	if err != nil {
		return nil, err
	}
	return storedUser, nil
}

func validateUser(user *domain.User) error {
	if strings.TrimSpace(user.Name) == "" {
		return errors.New("name cannot be empty")
	}
	if len(user.Name) > 255 {
		return errors.New("name must be less than or equal to 255 characters")
	}
	if strings.TrimSpace(user.Email) == "" {
		return errors.New("email cannot be empty")
	}
	if len(user.Email) > 255 {
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
