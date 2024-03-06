package service

import (
	"github.com/Jvls1/go-ecommerce/domain"
	"github.com/Jvls1/go-ecommerce/internal/repository"
)

func NewPermissionService(permissionRepository repository.PermissionRepository) PermissionService {
	return &permissionService{permissionRepository}
}

type permissionService struct {
	permissionRepository repository.PermissionRepository
}

type PermissionService interface {
	CreatePermission(permission *domain.Permission) (*domain.Permission, error)
}

func (permissionService *permissionService) CreatePermission(permission *domain.Permission) (*domain.Permission, error) {
	permissionCreated, err := permissionService.permissionRepository.CreatePermission(permission)
	if err != nil {
		return nil, err
	}
	return permissionCreated, err
}
