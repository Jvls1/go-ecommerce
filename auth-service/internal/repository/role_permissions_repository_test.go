package repository

import (
	"database/sql"
	"github.com/Jvls1/go-ecommerce/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRolePermissionsRepositoryWithContainer(t *testing.T) {
	rolePermissionsRepo := NewRolePermissionsRepository(db)
	role, permission := setupRolePermissionsValues(db)
	t.Run("Test add permission to a role", func(t *testing.T) {
		err := rolePermissionsRepo.AddPermissionToRole(role.ID, permission.ID)
		assert.NoError(t, err)
	})
	t.Run("Test add invalid permission to a invalid role", func(t *testing.T) {
		err := rolePermissionsRepo.AddPermissionToRole(uuid.New(), uuid.New())
		assert.Error(t, err)
	})
}

func setupRolePermissionsValues(db *sql.DB) (*domain.Role, *domain.Permission) {
	role := setupValidRoleToPermission()
	permission := setupValidPermission(db)
	return role, permission
}

func setupValidRoleToPermission() *domain.Role {
	roleRepo := NewRoleRepository(db)
	roleCreated, _ := roleRepo.CreateRole(domain.NewRole("Role for permission", "test"))
	return roleCreated
}

func setupValidPermission(db *sql.DB) *domain.Permission {
	permissionRepo := NewPermissionRepository(db)
	permission, _ := permissionRepo.CreatePermission(domain.NewPermission("Permission valid", "test"))
	return permission
}
