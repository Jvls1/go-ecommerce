package service

import (
	"errors"
	"github.com/Jvls1/go-ecommerce/domain"
	"github.com/Jvls1/go-ecommerce/internal/repository"
	"github.com/google/uuid"
	"strings"
)

func NewPermissionService(permissionRepository repository.PermissionRepository) PermissionService {
	return &permissionService{permissionRepository}
}

type permissionService struct {
	repo repository.PermissionRepository
}

type PermissionService interface {
	SavePermission(name, description string) (*domain.Permission, error)
	RetrievePermissionById(permissionId uuid.UUID) (*domain.Permission, error)
}

func (ps *permissionService) SavePermission(name, description string) (*domain.Permission, error) {
	err := validatePermissionValues(name, description)
	if err != nil {
		return nil, err
	}
	permission, err := domain.NewPermission(name, description)
	if err != nil {
		return nil, err
	}
	//TODO: Add validation for the unique constraint.
	permissionCreated, err := ps.repo.StorePermission(permission)
	if err != nil {
		return nil, err
	}
	return permissionCreated, err
}

func validatePermissionValues(name, description string) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("name cannot be empty")
	}
	if len(name) > 255 {
		return errors.New("name must be less than or equal to 255 characters")
	}
	if strings.TrimSpace(description) == "" {
		return errors.New("description cannot be empty")
	}
	if len(description) > 255 {
		return errors.New("description must be less than or equal to 255 characters")
	}
	return nil
}

func (ps *permissionService) RetrievePermissionById(permissionId uuid.UUID) (*domain.Permission, error) {
	permission, err := ps.repo.GetPermissionById(permissionId)
	if err != nil {
		return nil, err
	}
	return permission, nil
}
