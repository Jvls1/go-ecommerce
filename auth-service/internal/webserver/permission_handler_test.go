package webserver

import (
	"github.com/Jvls1/go-ecommerce/test/mocks"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestCreatePermission(t *testing.T) {
	//permissionService := createPermissionServiceMock(t)
	//handler := NewPermissionHandler(permissionService)
	//handler.CreatePermission()
}

func createPermissionServiceMock(t *testing.T) *mocks.MockPermissionService {
	ctrl := gomock.NewController(t)
	return mocks.NewMockPermissionService(ctrl)
}
