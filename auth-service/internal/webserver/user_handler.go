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

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{userService}
}

type userHandler struct {
	userService service.UserService
}

type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetUserById(w http.ResponseWriter, r *http.Request)
	AddRoleToUser(w http.ResponseWriter, r *http.Request)
}

func (uh *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userToSave *domain.User
	err := json.NewDecoder(r.Body).Decode(&userToSave)
	if err != nil {
		helpers.HandlerError(w, err, "Can't convert JSON", http.StatusBadRequest)
		return
	}
	user, err := uh.userService.SaveUser(userToSave)
	if err != nil {
		helpers.HandlerError(w, err, "Can't save the user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		helpers.HandlerError(w, err, "Error processing the request", http.StatusInternalServerError)
	}
}

func (uh *userHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	userUUID, err := uuid.Parse(userId)
	if err != nil {
		helpers.HandlerError(w, err, "Invalid UUID format", http.StatusBadRequest)
		return
	}
	user, err := uh.userService.RetrieveUserById(userUUID)
	if user == nil {
		helpers.HandlerError(w, err, "User not found", http.StatusNotFound)
		return
	}
	if err != nil {
		helpers.HandlerError(w, err, "Can't save the user", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		helpers.HandlerError(w, err, "Error processing the request", http.StatusInternalServerError)
	}
}

type AddRoleToUserRequest struct {
	UserId uuid.UUID `json:"userId"`
	RoleId uuid.UUID `json:"roleId"`
}

func (uh *userHandler) AddRoleToUser(w http.ResponseWriter, r *http.Request) {
	var addRoleToUserRequest AddRoleToUserRequest
	err := json.NewDecoder(r.Body).Decode(&addRoleToUserRequest)
	if err != nil {
		helpers.HandlerError(w, err, "Can't convert JSON", http.StatusBadRequest)
		return
	}
	err = uh.userService.AddRoleToUser(addRoleToUserRequest.UserId, addRoleToUserRequest.RoleId)
	if err != nil {
		helpers.HandlerError(w, err, "Can't add the permission to this role", http.StatusBadRequest)
		return
	}
}
