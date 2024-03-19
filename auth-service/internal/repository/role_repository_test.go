package repository

import (
	"database/sql"
	"github.com/Jvls1/go-ecommerce/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRoleRepositoryWithContainer(t *testing.T) {
	roleRepo := NewRoleRepository(db)
	t.Run("Test role creation", func(t *testing.T) {
		role, _ := domain.NewRole("Admin", "Admin role")
		roleCreated, err := roleRepo.StoreRole(role)
		assert.NotNil(t, roleCreated)
		assert.NoError(t, err)
	})
	t.Run("Test null fields", func(t *testing.T) {
		role, _ := domain.NewRole("TestTest", "test")
		roleCreated, err := roleRepo.StoreRole(role)
		assert.NoError(t, err)
		assert.Nil(t, roleCreated.DeletedAt)
	})
	t.Run("Test find role by valid ID", func(t *testing.T) {
		role, _ := domain.NewRole("Test2", "Test valid id")
		roleCreated, _ := roleRepo.StoreRole(role)
		role, err := roleRepo.GetRoleById(roleCreated.ID)
		assert.NotNil(t, role)
		assert.NoError(t, err, err)
	})
	t.Run("Test find role by random ID", func(t *testing.T) {
		randomUUID := uuid.New()
		_, err := roleRepo.GetRoleById(randomUUID)
		assert.Error(t, err)
	})
	t.Run("Test unique name constraint", func(t *testing.T) {
		role, _ := domain.NewRole("Unique", "Testing unique constraint")
		_, _ = roleRepo.StoreRole(role)
		roleCreated, err := roleRepo.StoreRole(role)
		assert.Error(t, err)
		assert.Nil(t, roleCreated)
	})
	t.Run("Test add permission to a role", func(t *testing.T) {
		role, permission := setupRolePermissionsValues(db)
		err := roleRepo.AddPermissionToRole(role.ID, permission.ID)
		assert.NoError(t, err)
	})
	t.Run("Test add invalid permission to a invalid role", func(t *testing.T) {
		err := roleRepo.AddPermissionToRole(uuid.New(), uuid.New())
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
	role, _ := domain.NewRole("Role for permission", "test")
	roleCreated, _ := roleRepo.StoreRole(role)
	return roleCreated
}

func setupValidPermission(db *sql.DB) *domain.Permission {
	permissionRepo := NewPermissionRepository(db)
	permission, _ := domain.NewPermission("Permission valid", "test")
	permissionCreated, _ := permissionRepo.StorePermission(permission)
	return permissionCreated
}
