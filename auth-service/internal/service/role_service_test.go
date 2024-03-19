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

func TestSaveRole(t *testing.T) {
	roleRepo := createRoleRepoMock(t)
	roleValid := createValidRole()
	roleRepo.EXPECT().StoreRole(roleValid).Return(roleValid, nil)
	roleServiceTest := NewRoleService(roleRepo)
	role, err := roleServiceTest.SaveRole(roleValid)
	assert.NotNil(t, role)
	assert.NoError(t, err)
}

func TestSaveInvalidRole(t *testing.T) {
	roleRepo := createRoleRepoMock(t)
	roleServiceTest := NewRoleService(roleRepo)
	role, err := roleServiceTest.SaveRole(createInvalidRole())
	assert.Nil(t, role)
	assert.Error(t, err)
}

func TestRetrieveRoleById(t *testing.T) {
	roleRepo := createRoleRepoMock(t)
	validId := uuid.New()
	role := &domain.Role{
		ID:          validId,
		Name:        "Test",
		Description: "Test",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   nil,
	}
	roleRepo.EXPECT().GetRoleById(validId).Return(role, nil)
	roleServiceTest := NewRoleService(roleRepo)

	roleRetrieved, err := roleServiceTest.RetrieveRoleById(validId)

	assert.NotNil(t, roleRetrieved)
	assert.NoError(t, err)
	assert.Equal(t, validId, roleRetrieved.ID)
}

func TestRetrieveRoleById_NotFound(t *testing.T) {
	roleRepo := createRoleRepoMock(t)
	validId := uuid.New()

	roleRepo.EXPECT().GetRoleById(validId).Return(nil, nil)

	roleServiceTest := NewRoleService(roleRepo)

	roleRetrieved, err := roleServiceTest.RetrieveRoleById(validId)

	assert.Nil(t, roleRetrieved)
	assert.NoError(t, err)
}

func createRoleRepoMock(t *testing.T) *mocks.MockRoleRepository {
	ctrl := gomock.NewController(t)
	return mocks.NewMockRoleRepository(ctrl)
}

func createValidRole() *domain.Role {
	role, _ := domain.NewRole("Admin", "Admin")
	return role
}

func createInvalidRole() *domain.Role {
	return &domain.Role{
		ID:          uuid.UUID{},
		Name:        "",
		Description: "",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	}
}
