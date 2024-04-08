package webserver

import (
	"encoding/json"
	"github.com/Jvls1/go-ecommerce/helpers"
	"github.com/Jvls1/go-ecommerce/internal/service"
	"github.com/go-chi/jwtauth/v5"
	"net/http"
)

func NewAuthHandler(userService service.UserService, tokenAuth *jwtauth.JWTAuth) AuthHandler {
	return &authHandler{userService, tokenAuth}
}

type authHandler struct {
	userService service.UserService
	tokenAuth   *jwtauth.JWTAuth
}

type AuthHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (ah *authHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		helpers.HandlerError(w, err, "Can't convert JSON", http.StatusBadRequest)
		return
	}
	token, err := ah.userService.Login(loginRequest.Email, loginRequest.Password, ah.tokenAuth)
	if err != nil {
		helpers.HandlerError(w, err, "Invalid email or password", http.StatusBadRequest)
		return
	}
	if token == "" {
		helpers.HandlerError(w, err, "Invalid email or password", http.StatusNotFound)
	}
	err = json.NewEncoder(w).Encode(token)
}
