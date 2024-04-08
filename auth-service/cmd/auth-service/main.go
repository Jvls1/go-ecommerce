package main

import (
	"database/sql"
	"github.com/Jvls1/go-ecommerce/internal/repository"
	"github.com/Jvls1/go-ecommerce/internal/service"
	"github.com/Jvls1/go-ecommerce/internal/webserver"
	"github.com/Jvls1/go-ecommerce/pkg"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

var tokenAuth *jwtauth.JWTAuth

var (
	permissionHandler webserver.PermissionHandler
	userHandler       webserver.UserHandler
	roleHandler       webserver.RoleHandler
	authHandler       webserver.AuthHandler
)

func init() {
	loadEnvVariables()
	tokenAuth = jwtauth.New("HS256", []byte(os.Getenv("JWT_SECRET")), nil)
}

func main() {
	db := pkg.NewDBConnection()
	permissionRepository, userRepository, roleRepository := createRepoInstances(db)
	permissionService, userService, roleService := createServicesInstance(permissionRepository, userRepository, roleRepository)
	createHandlersInstance(permissionService, userService, roleService)
	err := http.ListenAndServe(os.Getenv("PORT"), JSONMiddleware(router()))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func loadEnvVariables() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func createRepoInstances(db *sql.DB) (repository.PermissionRepository, repository.UserRepository, repository.RoleRepository) {
	permissionRepository := repository.NewPermissionRepository(db)
	userRepository := repository.NewUserRepository(db)
	roleRepository := repository.NewRoleRepository(db)
	return permissionRepository, userRepository, roleRepository
}

func createServicesInstance(permissionRepo repository.PermissionRepository, userRepo repository.UserRepository, roleRepo repository.RoleRepository) (service.PermissionService, service.UserService, service.RoleService) {
	permissionService := service.NewPermissionService(permissionRepo)
	userService := service.NewUserService(userRepo)
	roleService := service.NewRoleService(roleRepo)
	return permissionService, userService, roleService
}

func createHandlersInstance(permissionService service.PermissionService, userService service.UserService, roleService service.RoleService) {
	permissionHandler = webserver.NewPermissionHandler(permissionService)
	userHandler = webserver.NewUserHandler(userService)
	roleHandler = webserver.NewRoleHandler(roleService)
	authHandler = webserver.NewAuthHandler(userService, tokenAuth)
}

func router() http.Handler {
	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Post("/login", authHandler.Login)
	})
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator(tokenAuth))
		r.Route("/permissions", func(r chi.Router) {
			r.Get("/{permissionId}", permissionHandler.GetPermissionById)
			r.Post("/", permissionHandler.CreatePermission)
		})
		r.Route("/users", func(r chi.Router) {
			r.Get("/{userId}", userHandler.GetUserById)
			r.Post("/", userHandler.CreateUser)
			r.Post("/roles", userHandler.AddRoleToUser)
		})
		r.Route("/roles", func(r chi.Router) {
			r.Get("/{roleId}", roleHandler.GetRoleById)
			r.Post("/", roleHandler.CreateRole)
			r.Post("/permissions", roleHandler.AddPermissionToRole)
		})
	})
	return r
}

func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
