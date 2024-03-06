package repository

import (
	"database/sql"
	"github.com/Jvls1/go-ecommerce/domain"
	"github.com/Jvls1/go-ecommerce/testhelpers"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

var pgContainer *testhelpers.PostgresContainer
var db *sql.DB

func TestMain(m *testing.M) {
	pgContainer = testhelpers.SetupTestDatabase()
	defer func() {
		pgContainer.TearDown()
	}()
	db = setupDb(pgContainer.ConnectionString)
	os.Exit(m.Run())
}

func TestPermissionRepositoryWithContainer(t *testing.T) {
	permissionRepo := NewPermissionRepository(db)
	t.Run("Test permission creation", func(t *testing.T) {
		permission, _ := domain.NewPermission("Admin", "Admin permission")
		roleCreated, err := permissionRepo.CreatePermission(permission)
		assert.NotNil(t, roleCreated)
		assert.NoError(t, err)
	})
	t.Run("Test null fields", func(t *testing.T) {
		permission, _ := domain.NewPermission("TestTest", "test")
		roleCreated, err := permissionRepo.CreatePermission(permission)
		assert.NoError(t, err)
		assert.Nil(t, roleCreated.DeletedAt)
	})
	t.Run("Test find permission by valid ID", func(t *testing.T) {
		permission, _ := domain.NewPermission("Test2", "Test valid id")
		roleCreated, _ := permissionRepo.CreatePermission(permission)
		role, err := permissionRepo.FindPermissionById(roleCreated.ID)
		assert.NotNil(t, role)
		assert.NoError(t, err, err)
	})
	t.Run("Test find permission by random ID", func(t *testing.T) {
		randomUUID := uuid.New()
		_, err := permissionRepo.FindPermissionById(randomUUID)
		assert.Error(t, err)
	})
	t.Run("Test unique name constraint", func(t *testing.T) {
		permission, _ := domain.NewPermission("Unique", "Testing unique constraint")
		_, _ = permissionRepo.CreatePermission(permission)
		permissionCreated, err := permissionRepo.CreatePermission(permission)
		assert.Error(t, err)
		assert.Nil(t, permissionCreated)
	})
}

func setupDb(connectionString string) *sql.DB {
	database, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("failed to open connection with the database: %s", err)
	}
	return database
}
