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
		userCreated, err := userRepo.CreateUser(user)
		assert.NotNil(t, userCreated)
		assert.NoError(t, err)
	})
	t.Run("Test fields consistency", func(t *testing.T) {
		user, _ := domain.NewUser("Test2", "1234", "test2@test.com")
		userCreated, _ := userRepo.CreateUser(user)
		require.Equal(t, user.ID, userCreated.ID)
		require.Equal(t, user.Name, userCreated.Name)
		require.Equal(t, user.Email, userCreated.Email)
		require.Equal(t, user.CreatedAt, userCreated.CreatedAt)
		require.Equal(t, user.UpdatedAt, userCreated.UpdatedAt)
		require.Equal(t, user.DeletedAt, userCreated.DeletedAt)
	})
	t.Run("Test find user by valid ID", func(t *testing.T) {
		usr, _ := domain.NewUser("Valid User", "password", "validemail@email.com")
		userCreated, _ := userRepo.CreateUser(usr)
		user, err := userRepo.FindUserById(userCreated.ID)
		assert.NotNil(t, user)
		assert.NoError(t, err)
	})
	t.Run("Test find user by random ID", func(t *testing.T) {
		randomUUID := uuid.New()
		_, err := userRepo.FindUserById(randomUUID)
		assert.Error(t, err)
	})

}
