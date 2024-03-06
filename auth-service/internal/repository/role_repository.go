package repository

import (
	"database/sql"
	"github.com/Jvls1/go-ecommerce/domain"
	"github.com/google/uuid"
)

func NewRoleRepository(db *sql.DB) RoleRepository {
	return &roleRepository{db: db}
}

type roleRepository struct {
	db *sql.DB
}

type RoleRepository interface {
	CreateRole(role *domain.Role) (*domain.Role, error)
	FindById(id uuid.UUID) (*domain.Role, error)
}

func (roleRepository *roleRepository) CreateRole(role *domain.Role) (*domain.Role, error) {
	stmt, err := roleRepository.db.Prepare(`
		INSERT INTO roles (id, name, description, created_at, updated_at, deleted_at)
		VALUES ($1, $2, $3, $4, $5, $6) 
	`)
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(role.ID, role.Name, role.Description, role.CreatedAt, role.UpdatedAt, nil)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (roleRepository *roleRepository) FindById(id uuid.UUID) (*domain.Role, error) {
	stmt, err := roleRepository.db.Prepare(`
		SELECT id, name, description, created_at, updated_at
		  FROM roles r
		 WHERE r.id = $1
		   AND r.deleted_at IS NULL 
	`)
	if err != nil {
		return nil, err
	}
	role := &domain.Role{}
	err = stmt.QueryRow(id).Scan(&role.ID, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return role, nil
}
