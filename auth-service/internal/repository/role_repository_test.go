package repository

import (
	"github.com/Jvls1/go-ecommerce/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRoleRepositoryWithContainer(t *testing.T) {
	roleRepo := NewRoleRepository(db)
	t.Run("Test role creation", func(t *testing.T) {
		role, _ := domain.NewRole("Admin", "Admin role")
		roleCreated, err := roleRepo.CreateRole(role)
		assert.NotNil(t, roleCreated)
		assert.NoError(t, err)
	})
	t.Run("Test null fields", func(t *testing.T) {
		role, _ := domain.NewRole("TestTest", "test")
		roleCreated, err := roleRepo.CreateRole(role)
		assert.NoError(t, err)
		assert.Nil(t, roleCreated.DeletedAt)
	})
	t.Run("Test find role by valid ID", func(t *testing.T) {
		role, _ := domain.NewRole("Test2", "Test valid id")
		roleCreated, _ := roleRepo.CreateRole(role)
		role, err := roleRepo.FindById(roleCreated.ID)
		assert.NotNil(t, role)
		assert.NoError(t, err, err)
	})
	t.Run("Test find role by random ID", func(t *testing.T) {
		randomUUID := uuid.New()
		_, err := roleRepo.FindById(randomUUID)
		assert.Error(t, err)
	})
	t.Run("Test unique name constraint", func(t *testing.T) {
		role, _ := domain.NewRole("Unique", "Testing unique constraint")
		_, _ = roleRepo.CreateRole(role)
		roleCreated, err := roleRepo.CreateRole(role)
		assert.Error(t, err)
		assert.Nil(t, roleCreated)
	})
}