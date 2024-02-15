package service

import (
	"github.com/Jvls1/go-ecommerce/product-service/domain"
	"github.com/Jvls1/go-ecommerce/product-service/internal/repository"
)

type ProductService struct {
	ProductRepo repository.ProductRepo
}

func (productService *ProductService) CreateProduct(name, description, imageUrl string, price float64, quantity, departmentId int32) (*domain.Product, error) {
	product := domain.NewProduct(name, description, imageUrl, price, quantity, departmentId)
	_, err := productService.ProductRepo.CreateProduct(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (productService *ProductService) FindProducts(page int8, pageSize int) ([]*domain.Product, error) {
	products, err := productService.ProductRepo.FindProducts(page, pageSize)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (productService *ProductService) FindProductByID(id string) (*domain.Product, error) {
	product, err := productService.ProductRepo.FindProductById(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (productService *ProductService) FindProductByDepartmentId(departmentId string) ([]*domain.Product, error) {
	products, err := productService.ProductRepo.FindProductByDepartmentId(departmentId)
	if err != nil {
		return nil, err
	}
	return products, nil
}
