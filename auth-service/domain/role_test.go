package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateValidRole(t *testing.T) {
	role, err := NewRole("Admin", "Admin")
	assert.NoError(t, err)
	assert.NotNil(t, role)
}

func TestCreateInvalidRole(t *testing.T) {
	invalidSizeString := "vmXSsLW2T280oUzkHe9Pjoq1p0dL91AtTztUYiNkaAWpadioHh6cKvrZVp9j2LUH5914hI3zMBvK9oZ992YyjidnvtJ9IEOWQPF5JJLzs6PAe4WDUdv6v89uXPFWZeV6bI2DBbG5WnNhE1b6AeqFqumwrpYjM9yyeNUoHNbHG8a1qHudhR6HsF4G8JYCY9bjTM4Vh6iBT9tfYlFAdeFEwDjz5Pjz6IV19xLzTvsP9jR6xo1GPWgtra6supLAqMaF"
	t.Run("empty name", func(t *testing.T) {
		role, err := NewRole("", "Admin")
		assert.Error(t, err)
		assert.Nil(t, role)
	})
	t.Run("empty description", func(t *testing.T) {
		role, err := NewRole("Admin", "")
		assert.Error(t, err)
		assert.Nil(t, role)
	})
	t.Run("multiple spaces name", func(t *testing.T) {
		role, err := NewRole("          ", "Admin")
		assert.Error(t, err)
		assert.Nil(t, role)
	})
	t.Run("multiple spaces description", func(t *testing.T) {
		role, err := NewRole("Admin", "             ")
		assert.Error(t, err)
		assert.Nil(t, role)
	})
	t.Run("invalid name and description", func(t *testing.T) {
		role, err := NewRole("", "")
		assert.Error(t, err)
		assert.Nil(t, role)
	})
	t.Run("name greater than 255 characters", func(t *testing.T) {
		role, err := NewRole(invalidSizeString, "Admin")
		assert.Error(t, err)
		assert.Nil(t, role)
	})
	t.Run("description greater than 255 characters", func(t *testing.T) {
		role, err := NewRole("Admin", invalidSizeString)
		assert.Error(t, err)
		assert.Nil(t, role)
	})
}
