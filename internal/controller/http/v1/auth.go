package v1

import (
	"github.com/gofiber/fiber/v2"
)

// CreateUser godoc
// @ID create-user
// @Router /api/v1/auth/register [POST]
// @Tags auth
// @Summary create user
// @Description create user
// @Accept json
// @Produce json
// @Param user body auth_service_entity.User true "user body"
// @Success 200 {object} entity.SuccessResponse{data=string} "Success"
// @Response 400 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Response 422 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Failure 500 {object} entity.ErrorResponse{error=string} "Server Error"
func (r *V1) register(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}

// Login godoc
// @ID login
// @Router /api/v1/auth/login [POST]
// @Tags auth
// @Summary login
// @Description login
// @Accept json
// @Produce json
// @Param login body auth_service_entity.Login true "login"
// @Success 200 {object} entity.SuccessResponse{data=auth_service_entity.LoginResponse} "Success"
// @Response 400 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Response 422 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Failure 500 {object} entity.ErrorResponse{error=string} "Server Error"
func (r *V1) login(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}

// Activate godoc
// @ID activate
// @Router /api/v1/auth/activate/{activation_token} [GET]
// @Tags auth
// @Summary activate
// @Description activate
// @Accept json
// @Produce text/html
// @Param activation_token path string true "ACTIVATION TOKEN"
// @Success 200 {object} entity.SuccessResponse{data=string} "Success"
// @Response 400 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Response 422 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Failure 500 {object} entity.ErrorResponse{error=string} "Server Error"
func (r *V1) activate(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}

// ResendActivation godoc
// @ID resend-activation
// @Router /api/v1/auth/resend-activation/{activation_token} [GET]
// @Tags auth
// @Summary resend-activation
// @Description resend-activation
// @Accept json
// @Produce text/html
// @Param activation_token path string true "ACTIVATION TOKEN"
// @Success 200 {object} entity.SuccessResponse{data=string} "Success"
// @Response 400 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Response 422 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Failure 500 {object} entity.ErrorResponse{error=string} "Server Error"
func (r *V1) resendActivation(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}
