package repository

import (
	"database/sql"
	"errors"
	"github.com/Jvls1/go-ecommerce/product-service/domain"
)

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}

type productRepository struct {
	db *sql.DB
}

type ProductRepository interface {
	CreateProduct(product *domain.Product) (*domain.Product, error)
	FindProducts(page int, pageSize int) ([]*domain.Product, error)
	FindProductById(id string) (*domain.Product, error)
	FindProductByDepartmentId(departmentId string) ([]*domain.Product, error)
}

func (productRepository *productRepository) CreateProduct(product *domain.Product) (*domain.Product, error) {
	stmt, err := productRepository.db.Prepare("INSERT INTO products (id, name, description, image_url, price, quantity, department_id) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.ID, product.Name, product.Description, product.ImageURL, product.Price, product.Quantity, product.DepartmentID)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (productRepository *productRepository) FindProducts(page int, pageSize int) ([]*domain.Product, error) {
	offset := page * pageSize
	limit := pageSize

	stmt, err := productRepository.db.Prepare("SELECT id, name, description, image_url, price, quantity, department_id FROM products WHERE deleted_at is null LIMIT $1 OFFSET $2")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*domain.Product

	if !rows.Next() {
		return products, errors.New("no rows found")
	}

	for rows.Next() {
		var product domain.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.ImageURL, &product.Price, &product.Quantity, &product.DepartmentID, &product.CreatedAt, &product.UpdatedAt, &product.DeletedAt)

		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (productRepository *productRepository) FindProductById(id string) (*domain.Product, error) {
	var product domain.Product
	sqlQuery := "SELECT id, name, description, image_url, price, quantity, department_id FROM products WHERE id = ?"
	err := productRepository.db.QueryRow(sqlQuery, id).
		Scan(&product.ID, &product.Name, &product.Description, &product.ImageURL, &product.Price, &product.Quantity, &product.DepartmentID, &product.CreatedAt, &product.UpdatedAt, &product.DeletedAt)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (productRepository *productRepository) FindProductByDepartmentId(departmentId string) ([]*domain.Product, error) {
	sqlQuery := "SELECT id, name, description, image_url, price, quantity, department_id FROM products WHERE department_id = ?"
	rows, err := productRepository.db.Query(sqlQuery, departmentId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*domain.Product
	for rows.Next() {
		var product domain.Product
		err = rows.Scan(&product.ID, &product.Name, &product.Description, &product.ImageURL, &product.Price, &product.Quantity, &product.DepartmentID, &product.CreatedAt, &product.UpdatedAt, &product.DeletedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
