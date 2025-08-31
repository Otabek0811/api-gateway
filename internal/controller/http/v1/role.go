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
