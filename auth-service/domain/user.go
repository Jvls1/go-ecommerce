package domain

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        string       `json:"id"`
	Name      string       `json:"name"`
	Password  string       `json:"password"`
	Email     string       `json:"email"`
	RoleID    string       `json:"role"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"-"`
}

func NewProduct(name, password, email, role string) *User {
	return &User{
		ID:        uuid.New().String(),
		Name:      name,
		Password:  password,
		Email:     email,
		RoleID:    role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
