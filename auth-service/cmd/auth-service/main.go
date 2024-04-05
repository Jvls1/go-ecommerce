package main

import (
	"database/sql"
	"github.com/Jvls1/go-ecommerce/internal/repository"
	"github.com/Jvls1/go-ecommerce/internal/service"
	"github.com/Jvls1/go-ecommerce/internal/webserver"
	"github.com/Jvls1/go-ecommerce/pkg"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	loadEnvVariables()
	db := pkg.NewDBConnection()
	permissionRepository, userRepository, roleRepository := createRepoInstances(db)
	permissionService, userService, roleService := createServicesInstance(permissionRepository, userRepository, roleRepository)
	permissionHandler, userHandler, roleHandler := createHandlersInstance(permissionService, userService, roleService)
	defineRoutes(permissionHandler, userHandler, roleHandler)
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

func createHandlersInstance(permissionService service.PermissionService, userService service.UserService, roleService service.RoleService) (webserver.PermissionHandler, webserver.UserHandler, webserver.RoleHandler) {
	permissionHandler := webserver.NewPermissionHandler(permissionService)
	userHandler := webserver.NewUserHandler(userService)
	roleHandler := webserver.NewRoleHandler(roleService)
	return permissionHandler, userHandler, roleHandler
}

func defineRoutes(permissionHandler webserver.PermissionHandler, userHandler webserver.UserHandler, roleHandler webserver.RoleHandler) {
	r := chi.NewRouter()
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
	err := http.ListenAndServe(os.Getenv("PORT"), JSONMiddleware(r))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
