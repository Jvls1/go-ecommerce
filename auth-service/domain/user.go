package domain

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type User struct {
	ID        uuid.UUID    `json:"id"`
	Name      string       `json:"name"`
	Password  string       `json:"password,omitempty"`
	Email     string       `json:"email"`
	Roles     []Role       `json:"roles"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"-"`
}

func NewUser(name, plainPassword, email string) (*User, error) {
	err := validatedPlainPassword(plainPassword)
	if err != nil {
		return nil, err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &User{
		ID:        uuid.New(),
		Name:      name,
		Password:  string(hashedPassword),
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = user.isValid()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func validatedPlainPassword(plainPassword string) error {
	if strings.TrimSpace(plainPassword) == "" {
		return errors.New("password cannot be empty")
	}
	if len(plainPassword) > 255 {
		return errors.New("password cannot be empty")
	}
	return nil
}

func (user *User) isValid() error {
	if strings.TrimSpace(user.Name) == "" {
		return errors.New("name cannot be empty")
	}
	if len(user.Name) > 255 {
		return errors.New("name must be less than or equal to 255 characters")
	}
	if strings.TrimSpace(user.Email) == "" {
		return errors.New("email cannot be empty")
	}
	if len(user.Email) > 255 {
		return errors.New("description must be less than or equal to 255 characters")
	}
	return nil
}
