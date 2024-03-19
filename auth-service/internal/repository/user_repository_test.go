package repository

import (
	"github.com/Jvls1/go-ecommerce/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserRepositoryWithContainer(t *testing.T) {
	userRepo := NewUserRepository(db)
	t.Run("Test user creation", func(t *testing.T) {
		user, _ := domain.NewUser("Test1", "1234", "test@test.com")
		userCreated, err := userRepo.StoreUser(user)
		assert.NotNil(t, userCreated)
		assert.NoError(t, err)
	})
	t.Run("Test fields consistency", func(t *testing.T) {
		user, _ := domain.NewUser("Test2", "1234", "test2@test.com")
		userCreated, _ := userRepo.StoreUser(user)
		require.Equal(t, user.ID, userCreated.ID)
		require.Equal(t, user.Name, userCreated.Name)
		require.Equal(t, user.Email, userCreated.Email)
		require.Equal(t, user.CreatedAt, userCreated.CreatedAt)
		require.Equal(t, user.UpdatedAt, userCreated.UpdatedAt)
		require.Equal(t, user.DeletedAt, userCreated.DeletedAt)
	})
	t.Run("Test find user by valid ID", func(t *testing.T) {
		usr, _ := domain.NewUser("Valid User", "password", "validemail@email.com")
		userCreated, _ := userRepo.StoreUser(usr)
		user, err := userRepo.GetUserById(userCreated.ID)
		assert.NotNil(t, user)
		assert.NoError(t, err)
	})
	t.Run("Test find user by random ID", func(t *testing.T) {
		randomUUID := uuid.New()
		_, err := userRepo.GetUserById(randomUUID)
		assert.Error(t, err)
	})
	t.Run("Test unique email constraint", func(t *testing.T) {
		usr, _ := domain.NewUser("Valid User", "password", "validemail@email.com")
		_, _ = userRepo.StoreUser(usr)
		_, err := userRepo.StoreUser(usr)
		assert.Error(t, err)
	})
	t.Run("Test add role to a user", func(t *testing.T) {
		user, role := setupValues()
		err := userRepo.AddRoleToUser(user.ID, role.ID)
		assert.NoError(t, err)
	})
	t.Run("Test add invalid role to a invalid user", func(t *testing.T) {
		err := userRepo.AddRoleToUser(uuid.New(), uuid.New())
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
	userCreated, _ := userRepo.StoreUser(user)
	return userCreated
}

func setupValidRoleToUser() *domain.Role {
	roleRepo := NewRoleRepository(db)
	role, _ := domain.NewRole("Role for user", "test")
	roleCreated, _ := roleRepo.StoreRole(role)
	return roleCreated
}
