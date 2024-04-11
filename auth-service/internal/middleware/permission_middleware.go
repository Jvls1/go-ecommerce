package middleware

import (
	"github.com/Jvls1/go-ecommerce/internal/service"
	"github.com/go-chi/jwtauth/v5"
	"net/http"
	"strings"
)

func NewPermissionMiddleware(permissionService service.PermissionService) PermissionMiddleware {
	return &permissionMiddleware{permissionService}
}

type permissionMiddleware struct {
	permissionService service.PermissionService
}

type PermissionMiddleware interface {
	RequirePermission(permission string) func(http.Handler) http.Handler
}

func (pm *permissionMiddleware) RequirePermission(permissionToCheck string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := jwtauth.TokenFromHeader(r)
			if token == "" {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}
			_, claims, err := jwtauth.FromContext(r.Context())
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			if rolesRaw, ok := claims["roles"]; ok {
				if rolesStr, ok := rolesRaw.(string); ok {
					if !pm.hasPermission(rolesStr, permissionToCheck) {
						http.Error(w, "You do not have permission to access this feature", http.StatusForbidden)
						return
					}
				} else {
					if rolesArr, ok := rolesRaw.([]interface{}); ok {
						var rolesStrArr []string
						for _, role := range rolesArr {
							if roleStr, ok := role.(string); ok {
								rolesStrArr = append(rolesStrArr, roleStr)
							}
						}
						rolesStr := strings.Join(rolesStrArr, ",")

						if !pm.hasPermission(rolesStr, permissionToCheck) {
							http.Error(w, "You do not have permission to access this feature", http.StatusForbidden)
							return
						}
					} else {
						http.Error(w, "Invalid roles", http.StatusForbidden)
						return
					}
				}
			} else {
				http.Error(w, "Missing roles", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func (pm *permissionMiddleware) hasPermission(roles, permissionToCheck string) bool {
	if strings.Contains(roles, "Admin") {
		return true
	}
	permissions, err := pm.permissionService.RetrievePermissionsByRoleName(roles)
	if err != nil {
		return false
	}
	var hasPermission bool
	for _, permission := range permissions {
		if permission == permissionToCheck {
			hasPermission = true
			break
		}
	}
	return hasPermission
}
