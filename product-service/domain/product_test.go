package domain

import (
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	name := "Test Product"
	description := "Test Description"
	imageURL := "http://example.com/image.jpg"
	price := 19.99
	quantity := 100
	departmentID := "123"

	product := NewProduct(name, description, imageURL, price, int32(quantity), departmentID)

	assert.Equal(t, name, product.Name, "Product name is not set correctly")
	assert.Equal(t, description, product.Description, "Product description is not set correctly")
	assert.Equal(t, imageURL, product.ImageURL, "Product imageURL is not set correctly")
	assert.Equal(t, price, product.Price, "Product price is not set correctly")
	assert.Equal(t, int32(quantity), product.Quantity, "Product quantity is not set correctly")
	assert.Equal(t, departmentID, product.DepartmentID, "Product department ID is not set correctly")

	assert.NotEmpty(t, product.ID, "Product ID should not be empty")
	assert.False(t, product.CreatedAt.IsZero(), "Product CreatedAt should not be zero")
	assert.False(t, product.UpdatedAt.IsZero(), "Product UpdatedAt should not be zero")
}

func TestProductNullFields(t *testing.T) {
	product := NewProduct("Test Product", "Test Description", "http://example.com/image.jpg", 19.99, 100, "123")
	product.DeletedAt = sql.NullTime{Time: time.Now(), Valid: true}

	assert.True(t, product.DeletedAt.Valid, "DeletedAt should be valid")
}

func TestProductValidation(t *testing.T) {
	t.Run("EmptyDescription", func(t *testing.T) {
		product := NewProduct("123", "", "Description", 10.0, 5, "123")
		if !product.IsValid() {
			t.Error("Expected product with empty description to be valid")
		}
	})

	t.Run("EmptyName", func(t *testing.T) {
		product := NewProduct("", "Valid Name", "Description", 10.0, 5, "123")
		if product.IsValid() {
			t.Error("Expected product with empty name to be invalid")
		}
	})

	t.Run("NegativePrice", func(t *testing.T) {
		product := NewProduct("123", "Valid Name", "Description", -10.0, 5, "123")
		if product.IsValid() {
			t.Error("Expected product with negative price to be invalid")
		}
	})

	t.Run("NegativeQuantity", func(t *testing.T) {
		product := NewProduct("123", "Valid Name", "Description", 10.0, -5, "123")
		if product.IsValid() {
			t.Error("Expected product with negative quantity to be invalid")
		}
	})

	t.Run("EmptyDepartmentID", func(t *testing.T) {
		product := NewProduct("123", "Valid Name", "Description", 10.0, 5, "")
		if product.IsValid() {
			t.Error("Expected product with empty department ID to be invalid")
		}
	})

}
