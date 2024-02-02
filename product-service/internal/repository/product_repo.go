package repository

import (
	"database/sql"

	"github.com/Jvls1/go-ecommerce/product-service/domain"
)

type ProductRepo struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) *ProductRepo {
	return &ProductRepo{db: db}
}

func (productRepo *ProductRepo) CreateProduct(product *domain.Product) (*domain.Product, error) {
	_, err := productRepo.db.Exec("INSERT INTO products (id, name, description, image_url, price, quantity, department_id) VALUES (?, ?, ?, ?, ?, ?, ?)",
		product.ID, product.Name, product.Description, product.ImageURL, product.Price, product.Quantity, product.DepartmentID)
	if err != nil {
		return nil, err
	}
	return product, nil
}
