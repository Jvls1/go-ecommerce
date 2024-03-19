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

func TestSaveUser(t *testing.T) {
	userRepo := createUserRepoMock(t)
	userValid := createValidUser()
	userRepo.EXPECT().StoreUser(userValid).Return(userValid, nil)
	userServiceTest := NewUserService(userRepo)
	user, err := userServiceTest.SaveUser(userValid)
	assert.NotNil(t, user)
	assert.NoError(t, err)
}

func TestSaveInvalidUser(t *testing.T) {
	userRepo := createUserRepoMock(t)
	userServiceTest := NewUserService(userRepo)
	user, err := userServiceTest.SaveUser(createInvalidUser())
	assert.Nil(t, user)
	assert.Error(t, err)
}

func TestRetrieveUserById(t *testing.T) {
	userRepo := createUserRepoMock(t)
	validId := uuid.New()
	user := &domain.User{
		ID:        validId,
		Name:      "John Doe",
		Password:  "password",
		Email:     "test@test.com",
		Roles:     nil,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}
	userRepo.EXPECT().GetUserById(validId).Return(user, nil)
	userServiceTest := NewUserService(userRepo)

	userRetrieved, err := userServiceTest.RetrieveUserById(validId)

	assert.NotNil(t, userRetrieved)
	assert.NoError(t, err)
	assert.Equal(t, validId, userRetrieved.ID)
}

func TestRetrieveUserById_NotFound(t *testing.T) {
	userRepo := createUserRepoMock(t)
	validId := uuid.New()

	userRepo.EXPECT().GetUserById(validId).Return(nil, nil)

	userServiceTest := NewUserService(userRepo)

	userRetrieved, err := userServiceTest.RetrieveUserById(validId)

	assert.Nil(t, userRetrieved)
	assert.NoError(t, err)
}

func TestUserService_AddRoleToUser(t *testing.T) {
	userRepo := createUserRepoMock(t)
	userServiceTest := NewUserService(userRepo)
	user, _ := domain.NewUser("John Doe", "1234", "test@test.com")
	role, _ := domain.NewRole("Admin", "Admin")
	userRepo.EXPECT().AddRoleToUser(user.ID, role.ID).Return(nil)
	err := userServiceTest.AddRoleToUser(user.ID, role.ID)
	assert.NoError(t, err)
}

func createUserRepoMock(t *testing.T) *mocks.MockUserRepository {
	ctrl := gomock.NewController(t)
	return mocks.NewMockUserRepository(ctrl)
}

func createValidUser() *domain.User {
	user, _ := domain.NewUser("John Doe", "password", "test@test.com")
	return user
}

func createInvalidUser() *domain.User {
	return &domain.User{
		ID:        uuid.UUID{},
		Name:      "",
		Password:  "",
		Email:     "",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	}
}
