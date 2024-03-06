package domain

import (
	"errors"
	"github.com/google/uuid"
	"strings"
	"time"
)

type Role struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"-"`
}

func NewRole(name, description string) (*Role, error) {
	role := &Role{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   nil,
	}
	err := role.isValid()
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (role *Role) isValid() error {
	if strings.TrimSpace(role.Name) == "" {
		return errors.New("name cannot be empty")
	}
	if len(role.Name) > 255 {
		return errors.New("name must be less than or equal to 255 characters")
	}
	if strings.TrimSpace(role.Description) == "" {
		return errors.New("description cannot be empty")
	}
	if len(role.Description) > 255 {
		return errors.New("description must be less than or equal to 255 characters")
	}
	return nil
}
