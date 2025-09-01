package v1

import (
	"GolandProjects/api-gateway/config"
	"GolandProjects/api-gateway/internal/entity"
	"GolandProjects/api-gateway/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// V1 -.
type V1 struct {
	log         logger.Interface
	validator   *validator.Validate
	cfg         *config.Config
	permissions *config.Permissions
}

func (r *V1) HandleSuccessResponse(ctx *fiber.Ctx, code int, data interface{}) error {
	return ctx.Status(code).JSON(entity.SuccessResponse{
		Code:    code,
		Success: true,
		Data:    data,
	})
}

func (r *V1) HandleErrorResponse(ctx *fiber.Ctx, code int, msg string, err error) error {
	return ctx.Status(code).JSON(entity.ErrorResponse{
		Success: false,
		Message: msg,
		Error:   err.Error(),
		Code:    code,
	})
}
