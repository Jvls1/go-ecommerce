package webserver

import (
	"encoding/json"
	"github.com/Jvls1/go-ecommerce/domain"
	"github.com/Jvls1/go-ecommerce/helpers"
	"github.com/Jvls1/go-ecommerce/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"net/http"
)

func NewRoleHandler(roleService service.RoleService) RoleHandler {
	return &roleHandler{roleService}
}

type roleHandler struct {
	roleService service.RoleService
}

type RoleHandler interface {
	CreateRole(w http.ResponseWriter, r *http.Request)
	GetRoleById(w http.ResponseWriter, r *http.Request)
	AddPermissionToRole(w http.ResponseWriter, r *http.Request)
}

func (rh *roleHandler) CreateRole(w http.ResponseWriter, r *http.Request) {
	var roleToSave *domain.Role
	err := render.DecodeJSON(r.Body, &roleToSave)
	if err != nil {
		helpers.HandlerError(w, err, "Can't convert JSON", http.StatusBadRequest)
		return
	}
	role, err := rh.roleService.SaveRole(roleToSave.Name, roleToSave.Description)
	if err != nil {
		helpers.HandlerError(w, err, "Can't save the role", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(role)
	if err != nil {
		helpers.HandlerError(w, err, "Error processing the request", http.StatusInternalServerError)
	}
}

func (rh *roleHandler) GetRoleById(w http.ResponseWriter, r *http.Request) {
	roleId := chi.URLParam(r, "roleId")
	roleUUID, err := uuid.Parse(roleId)
	if err != nil {
		helpers.HandlerError(w, err, "Invalid UUID format", http.StatusBadRequest)
		return
	}
	role, err := rh.roleService.RetrieveRoleById(roleUUID)
	if role == nil {
		helpers.HandlerError(w, err, "Role not found", http.StatusNotFound)
		return
	}
	if err != nil {
		helpers.HandlerError(w, err, "Can't save the role", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(role)
	if err != nil {
		helpers.HandlerError(w, err, "Error processing the request", http.StatusInternalServerError)
	}
}

type AddPermissionToRoleRequest struct {
	RoleId       uuid.UUID `json:"roleId"`
	PermissionId uuid.UUID `json:"permissionId"`
}

func (rh *roleHandler) AddPermissionToRole(w http.ResponseWriter, r *http.Request) {
	var addPermissionToRole AddPermissionToRoleRequest
	err := json.NewDecoder(r.Body).Decode(&addPermissionToRole)
	if err != nil {
		helpers.HandlerError(w, err, "Can't convert JSON", http.StatusBadRequest)
		return
	}
	err = rh.roleService.AddPermissionToRole(addPermissionToRole.RoleId, addPermissionToRole.PermissionId)
	if err != nil {
		helpers.HandlerError(w, err, "Can't add the permission to this role", http.StatusBadRequest)
		return
	}
}
