package repository

import (
	"github.com/Jvls1/go-ecommerce/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRolesRepositoryWithContainer(t *testing.T) {
	userRolesRepo := NewUserRolesRepository(db)
	user, role := setupValues()
	t.Run("Test add role to a user", func(t *testing.T) {
		err := userRolesRepo.AddRoleToUser(user.ID, role.ID)
		assert.NoError(t, err)
	})
	t.Run("Test add invalid role to a invalid user", func(t *testing.T) {
		err := userRolesRepo.AddRoleToUser(uuid.New(), uuid.New())
		assert.Error(t, err)
	})
}

func setupValues() (*domain.User, *domain.Role) {
	user := setupValidUser()
	role := setupValidRoleToUser()
	return user, role
}

func setupValidUser() *domain.User {
	userRepo := NewUserRepository(db)
	user, _ := domain.NewUser("User with role", "1234", "userwithrole@email.com")
	userCreated, _ := userRepo.CreateUser(user)
	return userCreated
}

func setupValidRoleToUser() *domain.Role {
	roleRepo := NewRoleRepository(db)
	role, _ := domain.NewRole("Role for user", "test")
	roleCreated, _ := roleRepo.CreateRole(role)
	return roleCreated
}
