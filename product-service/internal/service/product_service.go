package service

import (
	"errors"
	"github.com/Jvls1/go-ecommerce/product-service/domain"
	"github.com/Jvls1/go-ecommerce/product-service/internal/repository"
	"github.com/google/uuid"
)

func NewProductService(productRepository repository.ProductRepository) ProductService {
	return &productService{productRepository}
}

type productService struct {
	ProductRepository repository.ProductRepository
}

type ProductService interface {
	CreateProduct(name, description, imageUrl string, price float64, quantity int32, departmentId string) (*domain.Product, error)
	FindProducts(page int, pageSize int) ([]*domain.Product, error)
	FindProductByID(id string) (*domain.Product, error)
	FindProductsByDepartmentId(departmentId string) ([]*domain.Product, error)
}

func (productService *productService) CreateProduct(name, description, imageUrl string, price float64, quantity int32, departmentId string) (*domain.Product, error) {
	product := domain.NewProduct(name, description, imageUrl, price, quantity, departmentId)
	if !product.IsValid() {
		return nil, errors.New("invalid product values")
	}
	_, err := productService.ProductRepository.CreateProduct(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (productService *productService) FindProducts(page int, pageSize int) ([]*domain.Product, error) {
	if page < 0 || pageSize <= 0 {
		return nil, errors.New("page must be zero or greater and pageSize must be positive integer")
	}
	products, err := productService.ProductRepository.FindProducts(page, pageSize)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (productService *productService) FindProductByID(id string) (*domain.Product, error) {
	err := uuid.Validate(id)
	if err != nil {
		return nil, err
	}
	product, err := productService.ProductRepository.FindProductById(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (productService *productService) FindProductsByDepartmentId(departmentId string) ([]*domain.Product, error) {
	err := uuid.Validate(departmentId)
	if err != nil {
		return nil, err
	}
	products, err := productService.ProductRepository.FindProductByDepartmentId(departmentId)
	if err != nil {
		return nil, err
	}
	return products, nil
}
