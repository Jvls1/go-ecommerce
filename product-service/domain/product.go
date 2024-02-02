package domain

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	ImageURL     string  `json:"image_url"`
	Price        float64 `json:"price"`
	Quantity     int32   `json:"quantity"`
	DepartmentID int32   `json:"department_id"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

func NewProduct(name, description, imageUrl string, price float64, quantity, departmentID int32) *Product {
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
