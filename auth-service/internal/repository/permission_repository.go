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
	StorePermission(permission *domain.Permission) (*domain.Permission, error)
	GetPermissionById(id uuid.UUID) (*domain.Permission, error)
	GetPermissionsByRoleName(name string) ([]string, error)
}

func (permissionRepository *permissionRepository) StorePermission(permission *domain.Permission) (*domain.Permission, error) {
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

func (permissionRepository *permissionRepository) GetPermissionById(id uuid.UUID) (*domain.Permission, error) {
	stmt, err := permissionRepository.db.Prepare(`
		SELECT id, name, description, created_at, updated_at
		  FROM permissions r
		 WHERE r.id = $1
		   AND r.deleted_at IS NULL 
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	permission := &domain.Permission{}
	err = stmt.QueryRow(id).Scan(&permission.ID, &permission.Name, &permission.Description, &permission.CreatedAt, &permission.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return permission, nil
}

func (permissionRepository *permissionRepository) GetPermissionsByRoleName(name string) ([]string, error) {
	stmt, err := permissionRepository.db.Prepare(`
		SELECT p.name
  		  FROM permissions p
  		  JOIN role_permissions rp ON p.id = rp.permission_id
  		  JOIN roles r             ON rp.role_id = r.id
 		 WHERE r.name = $1
		   AND p.deleted_at IS NULL
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []string
	for rows.Next() {
		var permissionName string
		err = rows.Scan(&permissionName)
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, permissionName)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return permissions, nil
}
