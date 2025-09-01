package v1

import (
	"github.com/gofiber/fiber/v2"
)

// GetPermission godoc
// @ID get-permission
// @Security ApiKeyAuth
// @Router /api/v1/permissions/{id} [GET]
// @Tags permissions
// @Summary get permission
// @Description get permission
// @Accept json
// @Produce json
// @Param id path string true "PERMISSION ID"
// @Success 200 {object} entity.SuccessResponse{data=auth_service_entity.Permission} "Success"
// @Response 400 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Response 422 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Failure 500 {object} entity.ErrorResponse{error=string} "Server Error"
func (r *V1) getPermission(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}

// CreatePermission godoc
// @Security ApiKeyAuth
// @ID create-permission
// @Router /api/v1/permissions [POST]
// @Tags permissions
// @Summary create permission
// @Description create permission
// @Accept json
// @Produce json
// @Param permission body auth_service_entity.Permission true "permission body"
// @Success 200 {object} entity.SuccessResponse{data=auth_service_entity.PermissionPK} "Success"
// @Response 400 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Response 422 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Failure 500 {object} entity.ErrorResponse{error=string} "Server Error"
func (r *V1) createPermission(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}

// GetListPermission godoc
// @ID get-list-permission
// @Security ApiKeyAuth
// @Router /api/v1/permissions [GET]
// @Tags permissions
// @Summary get list permissions
// @Description get list permissions
// @Accept json
// @Produce json
// @Param find query auth_service_entity.GetListFilter false "filters"
// @Success 200 {object} entity.SuccessResponse{data=auth_service_entity.GetPermissions} "Success"
// @Response 400 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Response 422 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Failure 500 {object} entity.ErrorResponse{error=string} "Server Error"
func (r *V1) getPermissionList(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}

// UpdatePermission godoc
// @ID update-permission
// @Security ApiKeyAuth
// @Router /api/v1/permissions/{id} [PUT]
// @Tags permissions
// @Summary update permission
// @Description update permission
// @Accept json
// @Produce json
// @Param id path string true "PERMISSION ID"
// @Param permission body auth_service_entity.UpdatePermission true "body"
// @Success 200 {object} entity.SuccessResponse{data=auth_service_entity.PermissionPK} "Success"
// @Response 400 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Response 422 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Failure 500 {object} entity.ErrorResponse{error=string} "Server Error"
func (r *V1) updatePermission(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}

// DeletePermission godoc
// @ID delete-permission
// @Security ApiKeyAuth
// @Router /api/v1/permissions/{id} [DELETE]
// @Tags permissions
// @Summary delete permission
// @Description delete permission
// @Accept json
// @Produce json
// @Param id path string true "PERMISSION ID"
// @Success 200 {object} entity.SuccessResponse{data=string} "Success"
// @Response 400 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Response 422 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Failure 500 {object} entity.ErrorResponse{error=string} "Server Error"
func (r *V1) deletePermission(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}

// UpdatePermissionField godoc
// @ID update-permission-field
// @Security ApiKeyAuth
// @Router /api/v1/permissions [PATCH]
// @Tags permissions
// @Summary update permission field
// @Description update permission field
// @Accept json
// @Produce json
// @Param role body auth_service_entity.UpdateFieldRequest true "body"
// @Success 200 {object} entity.SuccessResponse{data=auth_service_entity.PermissionPK} "Success"
// @Response 400 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Response 422 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Failure 500 {object} entity.ErrorResponse{error=string} "Server Error"
func (r *V1) updatePermissionField(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}

// AddRolePermission godoc
// @ID add-role-permission
// @Security ApiKeyAuth
// @Router /api/v1/permissions/add-role-permission [POST]
// @Tags permissions
// @Summary add role permission
// @Description add role permission
// @Accept json
// @Produce json
// @Param role body auth_service_entity.RolePermission true "body"
// @Success 200 {object} entity.SuccessResponse{data=auth_service_entity.PermissionPK} "Success"
// @Response 400 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Response 422 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Failure 500 {object} entity.ErrorResponse{error=string} "Server Error"
func (r *V1) addRolePermission(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}

// DeleteRolePermission godoc
// @ID delete-role-permission
// @Security ApiKeyAuth
// @Router /api/v1/permissions/delete-role-permission/{id} [DELETE]
// @Tags permissions
// @Summary delete role permission
// @Description delete role permission
// @Accept json
// @Produce json
// @Param id path string true "ROLE PERMISSION ID"
// @Success 200 {object} entity.SuccessResponse{data=string} "Success"
// @Response 400 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Response 422 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Failure 500 {object} entity.ErrorResponse{error=string} "Server Error"
func (r *V1) deleteRolePermission(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}

// AddUserPermission godoc
// @ID add-user-permission
// @Security ApiKeyAuth
// @Router /api/v1/permissions/add-user-permission [POST]
// @Tags permissions
// @Summary add user permission
// @Description add user permission
// @Accept json
// @Produce json
// @Param role body auth_service_entity.UserPermission true "body"
// @Success 200 {object} entity.SuccessResponse{data=auth_service_entity.PermissionPK} "Success"
// @Response 400 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Response 422 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Failure 500 {object} entity.ErrorResponse{error=string} "Server Error"
func (r *V1) addUserPermission(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}

// DeleteUserPermission godoc
// @ID delete-user-permission
// @Security ApiKeyAuth
// @Router /api/v1/permissions/delete-user-permission/{id} [DELETE]
// @Tags permissions
// @Summary delete user permission
// @Description delete user permission
// @Accept json
// @Produce json
// @Param id path string true "USER PERMISSION ID"
// @Success 200 {object} entity.SuccessResponse{data=string} "Success"
// @Response 400 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Response 422 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Failure 500 {object} entity.ErrorResponse{error=string} "Server Error"
func (r *V1) deleteUserPermission(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}
