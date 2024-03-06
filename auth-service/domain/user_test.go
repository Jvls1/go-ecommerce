package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateValidUser(t *testing.T) {
	user, err := NewUser("John Doe", "1234", "johndoe@email.com")
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestCreateInvalidUser(t *testing.T) {
	invalidSizeString := "vmXSsLW2T280oUzkHe9Pjoq1p0dL91AtTztUYiNkaAWpadioHh6cKvrZVp9j2LUH5914hI3zMBvK9oZ992YyjidnvtJ9IEOWQPF5JJLzs6PAe4WDUdv6v89uXPFWZeV6bI2DBbG5WnNhE1b6AeqFqumwrpYjM9yyeNUoHNbHG8a1qHudhR6HsF4G8JYCY9bjTM4Vh6iBT9tfYlFAdeFEwDjz5Pjz6IV19xLzTvsP9jR6xo1GPWgtra6supLAqMaF"
	validName := "John Doe"
	validPassword := "1234"
	validEmail := "email@email.com"
	t.Run("empty name", func(t *testing.T) {
		user, err := NewUser("", validPassword, validEmail)
		assert.Error(t, err)
		assert.Nil(t, user)
	})
	t.Run("empty password", func(t *testing.T) {
		user, err := NewUser(validName, "", validEmail)
		assert.Error(t, err)
		assert.Nil(t, user)
	})
	t.Run("empty email", func(t *testing.T) {
		user, err := NewUser(validName, validPassword, "")
		assert.Error(t, err)
		assert.Nil(t, user)
	})
	t.Run("multiple spaces name", func(t *testing.T) {
		user, err := NewUser("          ", validPassword, validEmail)
		assert.Error(t, err)
		assert.Nil(t, user)
	})
	t.Run("multiple spaces password", func(t *testing.T) {
		user, err := NewUser(validName, "    ", validEmail)
		assert.Error(t, err)
		assert.Nil(t, user)
	})
	t.Run("multiple spaces email", func(t *testing.T) {
		user, err := NewUser(validName, validPassword, "    ")
		assert.Error(t, err)
		assert.Nil(t, user)
	})
	t.Run("all fields invalid", func(t *testing.T) {
		user, err := NewUser("", "", "")
		assert.Error(t, err)
		assert.Nil(t, user)
	})
	t.Run("name greater than 255 characters", func(t *testing.T) {
		user, err := NewUser(invalidSizeString, validPassword, validEmail)
		assert.Error(t, err)
		assert.Nil(t, user)
	})
	t.Run("plain password greater than 255 characters", func(t *testing.T) {
		user, err := NewUser(validName, invalidSizeString, validEmail)
		assert.Error(t, err)
		assert.Nil(t, user)
	})
	t.Run("email greater than 255 characters", func(t *testing.T) {
		user, err := NewUser(validName, validPassword, invalidSizeString)
		assert.Error(t, err)
		assert.Nil(t, user)
	})
}
