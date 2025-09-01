package v1

import (
	"github.com/gofiber/fiber/v2"
)

// GetListUser godoc
// @ID get-list-user
// @Security ApiKeyAuth
// @Router /api/v1/users [GET]
// @Tags users
// @Summary get list users
// @Description get list users
// @Accept json
// @Produce json
// @Param find query auth_service_entity.GetListFilter false "filters"
// @Success 200 {object} entity.SuccessResponse{data=auth_service_entity.GetUsers} "Success"
// @Response 400 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Response 422 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Failure 500 {object} entity.ErrorResponse{error=string} "Server Error"
func (r *V1) getUserList(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}

// GetUser godoc
// @ID get-user
// @Security ApiKeyAuth
// @Router /api/v1/users/{id} [GET]
// @Tags users
// @Summary get user
// @Description get user
// @Accept json
// @Produce json
// @Param id path string true "USER ID"
// @Success 200 {object} entity.SuccessResponse{data=auth_service_entity.User} "Success"
// @Response 400 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Response 422 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Failure 500 {object} entity.ErrorResponse{error=string} "Server Error"
func (r *V1) getUser(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}

// UpdateUser godoc
// @ID update-user
// @Security ApiKeyAuth
// @Router /api/v1/users/{id} [PUT]
// @Tags users
// @Summary update user
// @Description update user
// @Accept json
// @Produce json
// @Param id path string true "USER ID"
// @Param user body auth_service_entity.UpdateUser true "body"
// @Success 200 {object} entity.SuccessResponse{data=auth_service_entity.UserPK} "Success"
// @Response 400 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Response 422 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Failure 500 {object} entity.ErrorResponse{error=string} "Server Error"
func (r *V1) updateUser(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}

// UpdateUserLogin godoc
// @ID update-user-login
// @Security ApiKeyAuth
// @Router /api/v1/users/user-login/{id} [PUT]
// @Tags users
// @Summary update user login
// @Description update user login
// @Accept json
// @Produce json
// @Param id path string true "USER ID"
// @Param user body auth_service_entity.UserLogin true "body"
// @Success 200 {object} entity.SuccessResponse{data=auth_service_entity.UserPK} "Success"
// @Response 400 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Response 422 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Failure 500 {object} entity.ErrorResponse{error=string} "Server Error"
func (r *V1) updateUserLogin(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}

// DeleteUser godoc
// @ID delete-user
// @Security ApiKeyAuth
// @Router /api/v1/users/{id} [DELETE]
// @Tags users
// @Summary delete user
// @Description delete user
// @Accept json
// @Produce json
// @Param id path string true "USER ID"
// @Success 200 {object} entity.SuccessResponse{data=string} "Success"
// @Response 400 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Response 422 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Failure 500 {object} entity.ErrorResponse{error=string} "Server Error"
func (r *V1) deleteUser(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}

// DeActivateUser godoc
// @ID deactivate-user
// @Security ApiKeyAuth
// @Router /api/v1/users/deactivate/{id} [PUT]
// @Tags users
// @Summary deactivate user
// @Description deactivate user
// @Accept json
// @Produce json
// @Param id path string true "USER ID"
// @Success 200 {object} entity.SuccessResponse{data=auth_service_entity.UserPK} "Success"
// @Response 400 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Response 422 {object} entity.ErrorResponse{error=string} "Bad Request"
// @Failure 500 {object} entity.ErrorResponse{error=string} "Server Error"
func (r *V1) deActivateUser(ctx *fiber.Ctx) error {
	if err := r.MakeProxy(ctx, r.cfg.ServiceUrl.Auth, ctx.OriginalURL()); err != nil {
		r.log.Error("proxy error", err)
		return r.HandleErrorResponse(ctx, fiber.StatusInternalServerError, "proxy error", err)
	}
	return nil
}
