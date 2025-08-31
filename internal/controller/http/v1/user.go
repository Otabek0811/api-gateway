package v1

import (
	"github.com/gofiber/fiber/v2"
)

func (r *V1) getUserList(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}

func (r *V1) getUser(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}

func (r *V1) updateUser(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}

func (r *V1) updateUserLogin(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}

func (r *V1) deleteUser(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}
