package service

import (
	"github.com/Jvls1/go-ecommerce/product-service/domain"
	"github.com/Jvls1/go-ecommerce/product-service/test/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
	"time"
)

func TestCreateProductValid(t *testing.T) {
	ctrl := gomock.NewController(t)
	productRepository := mocks.NewMockProductRepository(ctrl)

	product := domain.NewProduct("Test", "Test", "Test", 20.99, 100, "4321")

	productRepository.EXPECT().CreateProduct(gomock.Any()).Return(product, nil)

	productService := NewProductService(productRepository)
	result, err := productService.CreateProduct("Test", "Test", "Test", 20.99, 100, "4321")
	if err != nil {
		assert.Error(t, err)
	}
	if result == nil {
		assert.Fail(t, "Expect the product is created")
	}
}

func TestCreateProductInvalid(t *testing.T) {
	ctrl := gomock.NewController(t)
	productRepository := mocks.NewMockProductRepository(ctrl)

	productService := NewProductService(productRepository)
	result, err := productService.CreateProduct("", "Test", "Test", -20.22, 0, "")
	if err == nil {
		assert.Fail(t, "Expect an error because the product is invalid.")
	}
	if result != nil {
		assert.Fail(t, "The product is invalid and was created.")
	}
}

func TestFindProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	productRepository := mocks.NewMockProductRepository(ctrl)

	productRepository.EXPECT().FindProducts(0, 5).Return([]*domain.Product{}, nil)

	productService := NewProductService(productRepository)
	results, err := productService.FindProducts(0, 5)
	if err != nil {
		assert.Error(t, err)
	}
	if results == nil {
		assert.Fail(t, "Expect at least an empty array")
	}
}

func TestFindProductsWithInvalidValues(t *testing.T) {
	ctrl := gomock.NewController(t)
	productRepository := mocks.NewMockProductRepository(ctrl)

	productService := NewProductService(productRepository)
	results, err := productService.FindProducts(-1, 0)
	if err == nil {
		assert.Fail(t, "Expect error")
	}
	if results != nil {
		assert.Fail(t, "Unexpected result, invalid parameters was passed")
	}
}

func TestFindProductByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	productRepository := mocks.NewMockProductRepository(ctrl)

	productId := uuid.New().String()
	product := &domain.Product{
		ID:           productId,
		Name:         "Test",
		Description:  "Test",
		ImageURL:     "Test",
		Price:        20,
		Quantity:     1,
		DepartmentID: "1234",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	productRepository.EXPECT().FindProductById(productId).Return(product, nil)

	productService := NewProductService(productRepository)
	result, err := productService.FindProductByID(productId)
	if err != nil {
		assert.Error(t, err)
	}
	if result == nil {
		assert.Fail(t, "Expect a product")
	}
	if result.ID != productId {
		assert.Fail(t, "Find product with wrong ID")
	}
}

func TestFindProductByIdWithInvalidId(t *testing.T) {
	ctrl := gomock.NewController(t)
	productRepository := mocks.NewMockProductRepository(ctrl)

	productId := "1234"

	productService := NewProductService(productRepository)
	result, err := productService.FindProductByID(productId)
	if err == nil {
		assert.Fail(t, "Expect an error, the ID is invalid")
	}
	if result != nil {
		assert.Fail(t, "Unexpected product")
	}
}

func TestFindProductByDepartmentId(t *testing.T) {
	ctrl := gomock.NewController(t)
	productRepository := mocks.NewMockProductRepository(ctrl)

	departmentId := uuid.New().String()
	productRepository.EXPECT().FindProductByDepartmentId(departmentId).Return([]*domain.Product{}, nil)

	productService := NewProductService(productRepository)
	result, err := productService.FindProductsByDepartmentId(departmentId)
	if err != nil {
		assert.Error(t, err)
	}
	if result == nil {
		assert.Fail(t, "Expect at least an empty array")
	}
}

func TestFindProductByDepartmentIdWithInvalidId(t *testing.T) {
	ctrl := gomock.NewController(t)
	productRepository := mocks.NewMockProductRepository(ctrl)

	departmentId := "1234"

	productService := NewProductService(productRepository)
	result, err := productService.FindProductByID(departmentId)
	if err == nil {
		assert.Fail(t, "Expect an error, the ID is invalid")
	}
	if result != nil {
		assert.Fail(t, "Unexpected results")
	}
}
