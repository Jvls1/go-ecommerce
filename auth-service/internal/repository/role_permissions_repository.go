package repository

import (
	"database/sql"
	"github.com/google/uuid"
)

func NewRolePermissionsRepository(db *sql.DB) RolePermissionsRepository {
	return &rolePermissionsRepository{db: db}
}

type rolePermissionsRepository struct {
	db *sql.DB
}

type RolePermissionsRepository interface {
	AddPermissionToRole(permissionID, roleID uuid.UUID) error
}

func (rolePermissionsRepository *rolePermissionsRepository) AddPermissionToRole(roleID, permissionID uuid.UUID) error {
	stmt, err := rolePermissionsRepository.db.Prepare(`
		INSERT INTO role_permissions (role_id, permission_id) VALUES ($1, $2)
	`)
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(roleID, permissionID)
	if err != nil {
		return err
	}
	return nil
}
