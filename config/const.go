package config

import (
	"fmt"
	"strings"
)

type Permission struct {
	Name   string
	Method string
	Path   string
}

type Permissions struct {
	permissionsMap map[string]Permission
}

func NewPermissions() *Permissions {
	p := &Permissions{
		permissionsMap: make(map[string]Permission),
	}

	list := []Permission{

		// Permissions for Role
		{
			Name:   "CREATE_ROLE",
			Method: "POST",
			Path:   "/api/v1/roles",
		},
		{
			Name:   "GET_ALL_ROLES",
			Method: "GET",
			Path:   "/api/v1/roles",
		},
		{
			Name:   "GET_ROLE",
			Method: "GET",
			Path:   "/api/v1/roles/:id",
		},
		{
			Name:   "UPDATE_ROLE",
			Method: "PUT",
			Path:   "/api/v1/roles/:id",
		},
		{
			Name:   "DELETE_ROLE",
			Method: "DELETE",
			Path:   "/api/v1/roles/:id",
		},
		{
			Name:   "PATCH_ROLE",
			Method: "PATCH",
			Path:   "/api/v1/roles",
		},

		// Permissions for Permission
		{
			Name:   "CREATE_PERMISSION",
			Method: "POST",
			Path:   "/api/v1/permissions",
		},
		{
			Name:   "GET_ALL_PERMISSIONS",
			Method: "GET",
			Path:   "/api/v1/permissions",
		},
		{
			Name:   "GET_PERMISSION",
			Method: "GET",
			Path:   "/api/v1/permissions/:id",
		},
		{
			Name:   "UPDATE_PERMISSION",
			Method: "PUT",
			Path:   "/api/v1/permissions/:id",
		},
		{
			Name:   "DELETE_PERMISSION",
			Method: "DELETE",
			Path:   "/api/v1/permissions/:id",
		},
		{
			Name:   "PATCH_PERMISSION",
			Method: "PATCH",
			Path:   "/api/v1/permissions",
		},
		{
			Name:   "ADD_ROLE_PERMISSION",
			Method: "POST",
			Path:   "/api/v1/permissions/add-role-permission",
		},
		{
			Name:   "DELETE_ROLE_PERMISSION",
			Method: "DELETE",
			Path:   "/api/v1/permissions/delete-role-permission/:id",
		},
		{
			Name:   "ADD_USER_PERMISSION",
			Method: "POST",
			Path:   "/api/v1/permissions/add-user-permission",
		},
		{
			Name:   "DELETE_USER_PERMISSION",
			Method: "DELETE",
			Path:   "/api/v1/permissions/delete-user-permission/:id",
		},

		// Permissions for User
		{
			Name:   "GET_ALL_USERS",
			Method: "GET",
			Path:   "/api/v1/users",
		},
		{
			Name:   "GET_USER",
			Method: "GET",
			Path:   "/api/v1/users/:id",
		},
		{
			Name:   "UPDATE_USER",
			Method: "PUT",
			Path:   "/api/v1/users/:id",
		},
		{
			Name:   "DELETE_USER",
			Method: "DELETE",
			Path:   "/api/v1/users/:id",
		},
		{
			Name:   "UPDATE_USER_LOGIN",
			Method: "PUT",
			Path:   "/api/v1/users/user-login",
		},
		{
			Name:   "DEACTIVATE_USER",
			Method: "PUT",
			Path:   "/api/v1/users/deactivate/:id",
		},
	}

	for _, perm := range list {
		key := makeKey(perm.Method, perm.Path)
		p.permissionsMap[key] = perm
	}

	return p
}

func makeKey(method, path string) string {
	return method + " " + path
}

//func (p *Permissions) GetPermission(method, path string) (Permission, bool) {
//	perm, ok := p.permissionsMap[makeKey(method, path)]
//	return perm, ok
//}

func matchPath(routePath, reqPath string) bool {
	routeParts := strings.Split(strings.Trim(routePath, "/"), "/")
	reqParts := strings.Split(strings.Trim(reqPath, "/"), "/")

	if len(routeParts) != len(reqParts) {
		return false
	}

	for i := range routeParts {
		if strings.HasPrefix(routeParts[i], ":") {
			continue
		}
		if routeParts[i] != reqParts[i] {
			return false
		}
	}

	return true
}

func (p *Permissions) GetPermission(method, reqPath string) (Permission, bool) {
	reqPath = strings.Split(reqPath, "?")[0]
	fmt.Println(reqPath)
	for _, perm := range p.permissionsMap {
		if strings.EqualFold(perm.Method, method) && matchPath(perm.Path, reqPath) {
			return perm, true
		}
	}
	return Permission{}, false
}
