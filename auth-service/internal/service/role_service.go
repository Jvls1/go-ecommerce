package service

import (
	"errors"
	"github.com/Jvls1/go-ecommerce/domain"
	"github.com/Jvls1/go-ecommerce/internal/repository"
	"github.com/google/uuid"
	"strings"
)

func NewRoleService(roleRepository repository.RoleRepository) RoleService {
	return &roleService{roleRepository}
}

type roleService struct {
	roleRepo repository.RoleRepository
}

type RoleService interface {
	SaveRole(role *domain.Role) (*domain.Role, error)
	RetrieveRoleById(roleID uuid.UUID) (*domain.Role, error)
	AddPermissionToRole(roleID, permissionID uuid.UUID) error
}

func (rs *roleService) SaveRole(role *domain.Role) (*domain.Role, error) {
	err := validateRole(role)
	if err != nil {
		return nil, err
	}
	role, err = rs.roleRepo.StoreRole(role)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func validateRole(role *domain.Role) error {
	if strings.TrimSpace(role.Name) == "" {
		return errors.New("name cannot be empty")
	}
	if len(role.Name) > 255 {
		return errors.New("name must be less than or equal to 255 characters")
	}
	if strings.TrimSpace(role.Description) == "" {
		return errors.New("description cannot be empty")
	}
	if len(role.Description) > 255 {
		return errors.New("description must be less than or equal to 255 characters")
	}
	return nil
}

func (rs *roleService) RetrieveRoleById(roleId uuid.UUID) (*domain.Role, error) {
	return rs.roleRepo.GetRoleById(roleId)
}

func (rs *roleService) AddPermissionToRole(roleID, permissionID uuid.UUID) error {
	return rs.roleRepo.AddPermissionToRole(roleID, permissionID)
}
