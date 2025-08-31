package v1

import (
	"github.com/gofiber/fiber/v2"
)

func (r *V1) register(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}

func (r *V1) login(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}

func (r *V1) activate(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}

func (r *V1) resendActivation(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}
