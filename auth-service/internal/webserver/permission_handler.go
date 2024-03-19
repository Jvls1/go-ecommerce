package webserver

import (
	"github.com/Jvls1/go-ecommerce/domain"
	"github.com/Jvls1/go-ecommerce/internal/service"
	"net/http"
)

func NewPermissionHandler(permissionService service.PermissionService) PermissionHandler {
	return &permissionHandler{permissionService}
}

type permissionHandler struct {
	permissionService service.PermissionService
}

func (permissionHandler *permissionHandler) CreatePermission(w http.ResponseWriter, r *http.Request) (*domain.Permission, error) {
	//TODO implement me
	panic("implement me")
}

func (permissionHandler *permissionHandler) GetPermissionById(w http.ResponseWriter, r *http.Request) (*domain.Permission, error) {
	//TODO implement me
	panic("implement me")
}

type PermissionHandler interface {
	CreatePermission(w http.ResponseWriter, r *http.Request) (*domain.Permission, error)
	GetPermissionById(w http.ResponseWriter, r *http.Request) (*domain.Permission, error)
}
