package v1

import (
	"GolandProjects/api-gateway/config"
	"GolandProjects/api-gateway/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// NewRoutes -.
func NewRoutes(apiV1Group fiber.Router, log logger.Interface, cfg *config.Config, permissions *config.Permissions) {
	r := &V1{log: log, validator: validator.New(validator.WithRequiredStructEnabled()), cfg: cfg, permissions: permissions}

	authGroup := apiV1Group.Group("/auth")
	{
		authGroup.Post("/register", r.register)
		authGroup.Get("/activate/:activation_token", r.activate)
		authGroup.Get("/resend-activation/:activation_token", r.resendActivation)
		authGroup.Post("/login", r.login)
	}

	apiV1Group.Use(r.AuthMiddleware())

	userGroup := apiV1Group.Group("/users")

	{
		userGroup.Get("/", r.getUserList)
		userGroup.Get("/:id", r.getUser)
		userGroup.Put("/:id", r.updateUser)
		userGroup.Delete("/:id", r.deleteUser)
		userGroup.Put("/user-login", r.updateUserLogin)
	}
	roleGroup := apiV1Group.Group("/roles")
	{
		roleGroup.Post("/", r.createRole)
		roleGroup.Get("/", r.getRoleList)
		roleGroup.Get("/:id", r.getRole)
		roleGroup.Put("/:id", r.updateRole)
		roleGroup.Delete("/:id", r.deleteRole)
		roleGroup.Patch("/", r.updateRoleField)

	}

	permissionGroup := apiV1Group.Group("/permissions")
	{
		permissionGroup.Post("/", r.createPermission)
		permissionGroup.Get("/", r.getPermissionList)
		permissionGroup.Get("/:id", r.getPermission)
		permissionGroup.Put("/:id", r.updatePermission)
		permissionGroup.Delete("/:id", r.deletePermission)
		permissionGroup.Patch("/", r.updatePermissionField)
		permissionGroup.Post("/add-role-permission", r.addRolePermission)
		permissionGroup.Delete("/delete-role-permission/:id", r.deleteRolePermission)
		permissionGroup.Post("/add-user-permission", r.addUserPermission)
		permissionGroup.Delete("/delete-user-permission/:id", r.deleteUserPermission)
	}
}
