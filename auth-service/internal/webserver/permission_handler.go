package webserver

import (
	"encoding/json"
	"github.com/Jvls1/go-ecommerce/domain"
	"github.com/Jvls1/go-ecommerce/helpers"
	"github.com/Jvls1/go-ecommerce/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
)

func NewPermissionHandler(permissionService service.PermissionService) PermissionHandler {
	return &permissionHandler{permissionService}
}

type permissionHandler struct {
	permissionService service.PermissionService
}

type PermissionHandler interface {
	CreatePermission(w http.ResponseWriter, r *http.Request)
	GetPermissionById(w http.ResponseWriter, r *http.Request)
}

func (ph *permissionHandler) CreatePermission(w http.ResponseWriter, r *http.Request) {
	var permissionToSave domain.Permission
	err := json.NewDecoder(r.Body).Decode(&permissionToSave)
	if err != nil {
		helpers.HandlerError(w, err, "Can't convert JSON", http.StatusBadRequest)
		return
	}
	permission, err := ph.permissionService.SavePermission(permissionToSave.Name, permissionToSave.Description)
	if err != nil {
		helpers.HandlerError(w, err, "Can't save the permission", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(permission)
	if err != nil {
		helpers.HandlerError(w, err, "Error processing the request", http.StatusInternalServerError)
	}
}

func (ph *permissionHandler) GetPermissionById(w http.ResponseWriter, r *http.Request) {
	permissionId := chi.URLParam(r, "permissionId")
	permissionUUID, err := uuid.Parse(permissionId)
	if err != nil {
		helpers.HandlerError(w, err, "Invalid UUID format", http.StatusBadRequest)
		return
	}
	permission, err := ph.permissionService.RetrievePermissionById(permissionUUID)
	if permission == nil {
		helpers.HandlerError(w, err, "Permission not found", http.StatusNotFound)
		return
	}
	if err != nil {
		helpers.HandlerError(w, err, "Can't save the permission", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(permission)
	if err != nil {
		helpers.HandlerError(w, err, "Error processing the request", http.StatusInternalServerError)
	}
}
