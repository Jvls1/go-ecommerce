package domain

import (
	"errors"
	"github.com/google/uuid"
	"strings"
	"time"
)

type Permission struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"-"`
}

func NewPermission(name, description string) (*Permission, error) {
	permission := &Permission{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   nil,
	}
	err := permission.isValid()
	if err != nil {
		return nil, err
	}
	return permission, nil
}

func (permission *Permission) isValid() error {
	if strings.TrimSpace(permission.Name) == "" {
		return errors.New("name cannot be empty")
	}
	if len(permission.Name) > 255 {
		return errors.New("name must be less than or equal to 255 characters")
	}
	if strings.TrimSpace(permission.Description) == "" {
		return errors.New("description cannot be empty")
	}
	if len(permission.Description) > 255 {
		return errors.New("description must be less than or equal to 255 characters")
	}
	return nil
}
