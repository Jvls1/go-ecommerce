package repository

import (
	"database/sql"
	"github.com/Jvls1/go-ecommerce/domain"
	"github.com/google/uuid"
)

func NewPermissionRepository(db *sql.DB) PermissionRepository {
	return &permissionRepository{db: db}
}

type permissionRepository struct {
	db *sql.DB
}

type PermissionRepository interface {
	CreatePermission(permission *domain.Permission) (*domain.Permission, error)
	FindPermissionById(id uuid.UUID) (*domain.Permission, error)
}

func (permissionRepository *permissionRepository) CreatePermission(permission *domain.Permission) (*domain.Permission, error) {
	stmt, err := permissionRepository.db.Prepare(`
		INSERT INTO permissions (id, name, description, created_at, updated_at, deleted_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`)
	defer stmt.Close()
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(permission.ID, permission.Name, permission.Description, permission.CreatedAt, permission.UpdatedAt, nil)
	if err != nil {
		return nil, err
	}
	return permission, nil
}

func (permissionRepository *permissionRepository) FindPermissionById(id uuid.UUID) (*domain.Permission, error) {
	stmt, err := permissionRepository.db.Prepare(`
		SELECT id, name, description, created_at, updated_at
		  FROM permissions r
		 WHERE r.id = $1
		   AND r.deleted_at IS NULL 
	`)
	if err != nil {
		return nil, err
	}
	permission := &domain.Permission{}
	err = stmt.QueryRow(id).Scan(&permission.ID, &permission.Name, &permission.Description, &permission.CreatedAt, &permission.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return permission, nil
}
