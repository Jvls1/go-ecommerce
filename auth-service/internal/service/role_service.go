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
	SaveRole(name, description string) (*domain.Role, error)
	RetrieveRoleById(roleID uuid.UUID) (*domain.Role, error)
	AddPermissionToRole(roleID, permissionID uuid.UUID) error
}

func (rs *roleService) SaveRole(name, description string) (*domain.Role, error) {
	err := validateRole(name, description)
	if err != nil {
		return nil, err
	}
	role, _ := domain.NewRole(name, description)
	role, err = rs.roleRepo.StoreRole(role)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func validateRole(name, description string) error {
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

func (rs *roleService) RetrieveRoleById(roleId uuid.UUID) (*domain.Role, error) {
	return rs.roleRepo.GetRoleById(roleId)
}

func (rs *roleService) AddPermissionToRole(roleID, permissionID uuid.UUID) error {
	return rs.roleRepo.AddPermissionToRole(roleID, permissionID)
}
