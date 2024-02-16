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
	sql := "INSERT INTO products (id, name, description, image_url, price, quantity, department_id) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	_, err := productRepo.db.Exec(sql, product.ID, product.Name, product.Description, product.ImageURL, product.Price, product.Quantity, product.DepartmentID)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (productRepo *ProductRepo) FindProducts(page int, pageSize int) ([]*domain.Product, error) {
	offset := int(page-1) * pageSize
	limit := pageSize

	sql := "SELECT * FROM products LIMIT ? OFFSET ?"

	rows, err := productRepo.db.Query(sql, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*domain.Product

	for rows.Next() {
		var product domain.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Description)
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

func (productRepo *ProductRepo) FindProductById(id string) (*domain.Product, error) {
	var product domain.Product
	sql := "SELECT id, name, description, image_url, price, quantity, department_id FROM products WHERE id = ?"
	err := productRepo.db.QueryRow(sql, id).
		Scan(&product.ID, &product.Name, &product.Description, &product.ImageURL, &product.Price, &product.Quantity, &product.DepartmentID)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (productRepo *ProductRepo) FindProductByDepartmentId(departmentId string) ([]*domain.Product, error) {
	sql := "SELECT id, name, description, image_url, price, quantity, department_id FROM products WHERE department_id = ?"
	rows, err := productRepo.db.Query(sql, departmentId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*domain.Product
	for rows.Next() {
		var product domain.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.ImageURL, &product.Price, &product.Quantity, &product.DepartmentID)
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
