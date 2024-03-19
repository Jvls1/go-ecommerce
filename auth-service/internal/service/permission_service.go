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
	SavePermission(permission *domain.Permission) (*domain.Permission, error)
	RetrievePermissionById(permissionId uuid.UUID) (*domain.Permission, error)
}

func (ps *permissionService) SavePermission(permission *domain.Permission) (*domain.Permission, error) {
	err := validatePermission(permission)
	if err != nil {
		return nil, err
	}
	permissionCreated, err := ps.repo.StorePermission(permission)
	if err != nil {
		return nil, err
	}
	return permissionCreated, err
}

func validatePermission(permission *domain.Permission) error {
	if strings.TrimSpace(permission.Name) == "" {
		return errors.New("name cannot be empty")
	}
	if len(permission.Name) > 255 {
		return errors.New("name must be less than or equal to 255 characters")
	}
	if strings.TrimSpace(permission.Description) == "" {
		return errors.New("description cannot be empty")
	}
	if len(permission.Description) > 255 {
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
