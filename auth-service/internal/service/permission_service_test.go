package service

import (
	"github.com/Jvls1/go-ecommerce/domain"
	"github.com/Jvls1/go-ecommerce/test/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
	"time"
)

func TestSavePermission(t *testing.T) {
	permissionRepo := createPermissionRepoMock(t)
	permission := createValidPermission()
	permissionRepo.EXPECT().StorePermission(permission).Return(permission, nil)
	permissionServiceTest := NewPermissionService(permissionRepo)

	permissionSaved, err := permissionServiceTest.SavePermission(permission)

	assert.NoError(t, err)
	assert.NotNil(t, permissionSaved)
}

func TestSavePermissionInvalid(t *testing.T) {
	permissionRepo := createPermissionRepoMock(t)
	invalidPermission := createInvalidPermission()
	permissionServiceTest := NewPermissionService(permissionRepo)

	permissionSaved, err := permissionServiceTest.SavePermission(invalidPermission)

	assert.Nil(t, permissionSaved)
	assert.Error(t, err)
}

func TestRetrievePermissionById(t *testing.T) {
	permissionRepo := createPermissionRepoMock(t)
	validId := uuid.New()
	permission := &domain.Permission{
		ID:          validId,
		Name:        "Test",
		Description: "Test",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   nil,
	}
	permissionRepo.EXPECT().GetPermissionById(validId).Return(permission, nil)
	permissionServiceTest := NewPermissionService(permissionRepo)

	permissionRetrieved, err := permissionServiceTest.RetrievePermissionById(validId)

	assert.NotNil(t, permissionRetrieved)
	assert.NoError(t, err)
	assert.Equal(t, validId, permissionRetrieved.ID)
}

func TestRetrievePermissionById_NotFound(t *testing.T) {
	permissionRepo := createPermissionRepoMock(t)
	validId := uuid.New()

	permissionRepo.EXPECT().GetPermissionById(validId).Return(nil, nil)

	permissionServiceTest := NewPermissionService(permissionRepo)

	permissionRetrieved, err := permissionServiceTest.RetrievePermissionById(validId)

	assert.Nil(t, permissionRetrieved)
	assert.NoError(t, err)
}

func createPermissionRepoMock(t *testing.T) *mocks.MockPermissionRepository {
	ctrl := gomock.NewController(t)
	return mocks.NewMockPermissionRepository(ctrl)
}

func createValidPermission() *domain.Permission {
	permission, _ := domain.NewPermission("create-product", "allow product creation")
	return permission
}

func createInvalidPermission() *domain.Permission {
	return &domain.Permission{
		ID:          uuid.New(),
		Name:        "",
		Description: "",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   nil,
	}
}
