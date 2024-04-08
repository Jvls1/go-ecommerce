package service

import (
	"errors"
	"github.com/Jvls1/go-ecommerce/domain"
	"github.com/Jvls1/go-ecommerce/internal/repository"
	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/mail"
	"strings"
	"time"
)

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository}
}

type userService struct {
	repo repository.UserRepository
}

type UserService interface {
	SaveUser(name, password, email string) (*domain.User, error)
	RetrieveUserById(userID uuid.UUID) (*domain.User, error)
	RetrieveUserByEmail(email string) (*domain.User, error)
	AddRoleToUser(userID, roleID uuid.UUID) error
	Login(email, password string, tokenAuth *jwtauth.JWTAuth) (string, error)
}

func (us *userService) SaveUser(name, password, email string) (*domain.User, error) {
	err := validateUser(name, email)
	if err != nil {
		return nil, err
	}
	user, err := domain.NewUser(name, password, email)
	if err != nil {
		return nil, err
	}
	storedUser, err := us.repo.StoreUser(user)
	if err != nil {
		return nil, err
	}
	return storedUser, nil
}

func validateUser(name, email string) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("name cannot be empty")
	}
	if len(name) > 255 {
		return errors.New("name must be less than or equal to 255 characters")
	}
	if strings.TrimSpace(email) == "" {
		return errors.New("email cannot be empty")
	}
	if len(email) > 255 {
		return errors.New("description must be less than or equal to 255 characters")
	}
	return nil
}

func (us *userService) RetrieveUserById(userID uuid.UUID) (*domain.User, error) {
	return us.repo.GetUserById(userID)
}

func (us *userService) RetrieveUserByEmail(email string) (*domain.User, error) {
	if validateEmail(email) {
		return nil, errors.New("invalid email")
	}
	user, err := us.repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func validateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err != nil
}

func (us *userService) AddRoleToUser(userID, roleID uuid.UUID) error {
	// TODO Validate if the role is not add already
	return us.repo.AddRoleToUser(userID, roleID)
}

func (us *userService) Login(email, password string, tokenAuth *jwtauth.JWTAuth) (string, error) {
	user, err := us.RetrieveUserByEmail(email)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("user not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid password")
	}
	roles, err := us.repo.GetRolesByUserId(user.ID)
	if err != nil {
		return "", err
	}
	var rolesName []string
	for _, role := range roles {
		rolesName = append(rolesName, role.Name)
	}
	rolesString := strings.Join(rolesName, ",")
	tokenString, err := MakeToken(email, rolesString, tokenAuth)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func MakeToken(email, roles string, tokenAuth *jwtauth.JWTAuth) (string, error) {
	_, tokenString, err := tokenAuth.Encode(map[string]interface{}{"email": email, "roles": roles, "exp": time.Now().Add(time.Hour * 24).Unix()})
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
