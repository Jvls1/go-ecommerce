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
	StoreRole(role *domain.Role) (*domain.Role, error)
	GetRoleById(id uuid.UUID) (*domain.Role, error)
	AddPermissionToRole(roleID, permissionID uuid.UUID) error
}

func (roleRepository *roleRepository) StoreRole(role *domain.Role) (*domain.Role, error) {
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

func (roleRepository *roleRepository) GetRoleById(id uuid.UUID) (*domain.Role, error) {
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
	permissions, err := roleRepository.getPermissionsByRoleId(role.ID)
	if err != nil {
		return nil, err
	}
	role.Permissions = permissions
	return role, nil
}

func (roleRepository *roleRepository) getPermissionsByRoleId(roleId uuid.UUID) ([]domain.Permission, error) {
	permissionsStmt, err := roleRepository.db.Prepare(`
		SELECT p.id, p.name, p.description, p.created_at, p.updated_at
		  FROM role_permissions rp
		 INNER JOIN permissions p ON rp.permission_id = p.id
		 WHERE rp.role_id = $1;
	`)
	if err != nil {
		return nil, err
	}
	rows, err := permissionsStmt.Query(roleId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []domain.Permission
	for rows.Next() {
		var permission domain.Permission
		err = rows.Scan(&permission.ID, &permission.Name, &permission.Description, &permission.CreatedAt, &permission.UpdatedAt)
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}
	return permissions, nil
}

func (roleRepository *roleRepository) AddPermissionToRole(roleID, permissionID uuid.UUID) error {
	stmt, err := roleRepository.db.Prepare(`
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
