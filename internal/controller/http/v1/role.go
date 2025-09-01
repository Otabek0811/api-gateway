package v1

import (
	"github.com/gofiber/fiber/v2"
)

// GetRole godoc
// @ID get-role
// @Security ApiKeyAuth
// @Router /api/v1/roles/{id} [GET]
// @Tags roles
// @Summary get role
// @Description get role
// @Accept json
// @Produce json
// @Param id path string true "ROLE ID"
// @Success 200 {object} entity.SuccessResponse{data=auth_service_entity.Role} "Success"
// @Response 400 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Response 422 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Failure 500 {object} entity.ErrorResponse{error=string} "Server Error"
func (r *V1) getRole(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}

// CreateRole godoc
// @Security ApiKeyAuth
// @ID create-role
// @Router /api/v1/roles [POST]
// @Tags roles
// @Summary create role
// @Description create role
// @Accept json
// @Produce json
// @Param role body auth_service_entity.Role true "role body"
// @Success 200 {object} entity.SuccessResponse{data=auth_service_entity.RolePK} "Success"
// @Response 400 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Response 422 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Failure 500 {object} entity.ErrorResponse{error=string} "Server Error"
func (r *V1) createRole(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}

// GetListRole godoc
// @ID get-list-role
// @Security ApiKeyAuth
// @Router /api/v1/roles [GET]
// @Tags roles
// @Summary get list roles
// @Description get list roles
// @Accept json
// @Produce json
// @Param find query auth_service_entity.GetListFilter false "filters"
// @Success 200 {object} entity.SuccessResponse{data=auth_service_entity.GetRoles} "Success"
// @Response 400 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Response 422 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Failure 500 {object} entity.ErrorResponse{error=string} "Server Error"
func (r *V1) getRoleList(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}

// UpdateRole godoc
// @ID update-role
// @Security ApiKeyAuth
// @Router /api/v1/roles/{id} [PUT]
// @Tags roles
// @Summary update role
// @Description update role
// @Accept json
// @Produce json
// @Param id path string true "ROLE ID"
// @Param role body auth_service_entity.UpdateRole true "body"
// @Success 200 {object} entity.SuccessResponse{data=auth_service_entity.RolePK} "Success"
// @Response 400 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Response 422 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Failure 500 {object} entity.ErrorResponse{error=string} "Server Error"
func (r *V1) updateRole(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}

// DeleteRole godoc
// @ID delete-role
// @Security ApiKeyAuth
// @Router /api/v1/roles/{id} [DELETE]
// @Tags roles
// @Summary delete role
// @Description delete role
// @Accept json
// @Produce json
// @Param id path string true "ROLE ID"
// @Success 200 {object} entity.SuccessResponse{data=string} "Success"
// @Response 400 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Response 422 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Failure 500 {object} entity.ErrorResponse{error=string} "Server Error"
func (r *V1) deleteRole(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}

// UpdateRoleField godoc
// @ID update-role-field
// @Security ApiKeyAuth
// @Router /api/v1/roles [PATCH]
// @Tags roles
// @Summary update role field
// @Description update role field
// @Accept json
// @Produce json
// @Param role body auth_service_entity.UpdateFieldRequest true "body"
// @Success 200 {object} entity.SuccessResponse{data=auth_service_entity.RolePK} "Success"
// @Response 400 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Response 422 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Failure 500 {object} entity.ErrorResponse{error=string} "Server Error"
func (r *V1) updateRoleField(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}
