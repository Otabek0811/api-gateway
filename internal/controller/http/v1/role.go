package v1

import (
	"github.com/gofiber/fiber/v2"
)

func (r *V1) getRole(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}

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
// @Tags role
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

func (r *V1) updateRole(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}

func (r *V1) deleteRole(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}

func (r *V1) updateRoleField(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}
