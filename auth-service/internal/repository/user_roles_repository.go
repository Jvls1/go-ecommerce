package repository

import (
	"database/sql"
	"github.com/google/uuid"
)

func NewUserRolesRepository(db *sql.DB) UserRolesRepository {
	return &userRolesRepository{db: db}
}

type userRolesRepository struct {
	db *sql.DB
}

type UserRolesRepository interface {
	AddRoleToUser(userID, roleID uuid.UUID) error
}

func (userRolesRepository *userRolesRepository) AddRoleToUser(userID, roleID uuid.UUID) error {
	stmt, err := userRolesRepository.db.Prepare(`
		INSERT INTO user_roles (user_id, role_id) VALUES ($1, $2)
	`)
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(userID, roleID)
	if err != nil {
		return err
	}
	return nil
}
