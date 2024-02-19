package domain

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	ImageURL     string       `json:"image_url"`
	Price        float64      `json:"price"`
	Quantity     int32        `json:"quantity"`
	DepartmentID string       `json:"department_id"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
	DeletedAt    sql.NullTime `json:"-"`
}

func NewProduct(name, description, imageUrl string, price float64, quantity int32, departmentID string) *Product {
	return &Product{
		ID:           uuid.New().String(),
		Name:         name,
		Description:  description,
		ImageURL:     imageUrl,
		Price:        price,
		Quantity:     quantity,
		DepartmentID: departmentID,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func (p *Product) IsValid() bool {
	if p.ID == "" || p.Name == "" || p.Price < 0 || p.Quantity < 0 || p.DepartmentID == "" {
		return false
	}

	return true
}
