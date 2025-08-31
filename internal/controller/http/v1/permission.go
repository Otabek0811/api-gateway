package v1

import (
	"github.com/gofiber/fiber/v2"
)

func (r *V1) getPermission(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}

func (r *V1) createPermission(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}

func (r *V1) getPermissionList(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}

func (r *V1) updatePermission(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}

func (r *V1) deletePermission(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}

func (r *V1) updatePermissionField(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}

func (r *V1) addRolePermission(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}

func (r *V1) deleteRolePermission(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}

func (r *V1) addUserPermission(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}

func (r *V1) deleteUserPermission(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}
