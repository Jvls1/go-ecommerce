package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateValidPermission(t *testing.T) {
	permission, err := NewPermission("create-product", "create-product")
	assert.NoError(t, err)
	assert.NotNil(t, permission)
}

func TestCreateInvalidPermission(t *testing.T) {
	invalidSizeString := "vmXSsLW2T280oUzkHe9Pjoq1p0dL91AtTztUYiNkaAWpadioHh6cKvrZVp9j2LUH5914hI3zMBvK9oZ992YyjidnvtJ9IEOWQPF5JJLzs6PAe4WDUdv6v89uXPFWZeV6bI2DBbG5WnNhE1b6AeqFqumwrpYjM9yyeNUoHNbHG8a1qHudhR6HsF4G8JYCY9bjTM4Vh6iBT9tfYlFAdeFEwDjz5Pjz6IV19xLzTvsP9jR6xo1GPWgtra6supLAqMaF"
	t.Run("empty name", func(t *testing.T) {
		permission, err := NewPermission("", "create-product")
		assert.Error(t, err)
		assert.Nil(t, permission)
	})
	t.Run("empty description", func(t *testing.T) {
		permission, err := NewPermission("create-product", "")
		assert.Error(t, err)
		assert.Nil(t, permission)
	})
	t.Run("multiple spaces name", func(t *testing.T) {
		permission, err := NewPermission("          ", "create-product")
		assert.Error(t, err)
		assert.Nil(t, permission)
	})
	t.Run("multiple spaces description", func(t *testing.T) {
		permission, err := NewPermission("create-product", "             ")
		assert.Error(t, err)
		assert.Nil(t, permission)
	})
	t.Run("invalid name and description", func(t *testing.T) {
		permission, err := NewPermission("", "")
		assert.Error(t, err)
		assert.Nil(t, permission)
	})
	t.Run("name greater than 255 characters", func(t *testing.T) {
		permission, err := NewPermission(invalidSizeString, "create-product")
		assert.Error(t, err)
		assert.Nil(t, permission)
	})
	t.Run("description greater than 255 characters", func(t *testing.T) {
		permission, err := NewPermission("create-product", invalidSizeString)
		assert.Error(t, err)
		assert.Nil(t, permission)
	})
}
