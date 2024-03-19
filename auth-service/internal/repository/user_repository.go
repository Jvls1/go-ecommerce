package repository

import (
	"database/sql"
	"github.com/Jvls1/go-ecommerce/domain"
	"github.com/google/uuid"
)

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

type userRepository struct {
	db *sql.DB
}

type UserRepository interface {
	StoreUser(user *domain.User) (*domain.User, error)
	GetUserById(id uuid.UUID) (*domain.User, error)
	AddRoleToUser(userID, roleID uuid.UUID) error
}

func (userRepository *userRepository) StoreUser(user *domain.User) (*domain.User, error) {
	stmt, err := userRepository.db.Prepare(`
		INSERT INTO users (id, name, password, email, created_at, updated_at, deleted_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`)
	defer stmt.Close()
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(user.ID, user.Name, user.Password, user.Email, user.CreatedAt, user.UpdatedAt, user.DeletedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (userRepository *userRepository) GetUserById(id uuid.UUID) (*domain.User, error) {
	stmt, err := userRepository.db.Prepare(`
		SELECT id, name, email, created_at, updated_at
		FROM users u
		WHERE u.id = $1
		AND deleted_at IS NULL
	`)
	defer stmt.Close()
	if err != nil {
		return nil, err
	}
	var user domain.User
	err = stmt.QueryRow(id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	roles, err := userRepository.getRolesByUserId(id)
	if err != nil {
		return nil, err
	}
	user.Roles = roles
	return &user, nil
}

func (userRepository *userRepository) getRolesByUserId(userId uuid.UUID) ([]domain.Role, error) {
	stmt, err := userRepository.db.Prepare(`SELECT role_id FROM user_roles ur WHERE ur.user_id = $1`)
	defer stmt.Close()
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []domain.Role
	for rows.Next() {
		var role domain.Role
		err = rows.Scan(&role.ID, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt, &role.DeletedAt)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	return roles, nil
}

func (userRepository *userRepository) AddRoleToUser(userID, roleID uuid.UUID) error {
	stmt, err := userRepository.db.Prepare(`
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
