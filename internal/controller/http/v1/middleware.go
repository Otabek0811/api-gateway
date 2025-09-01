package v1

import (
	"GolandProjects/api-gateway/internal/entity/auth_service_entity"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io"
	"net/http"
	"strings"
	"time"
)

func (r *V1) AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {

		bearerToken := c.Get("Authorization")
		if bearerToken == "" {
			return r.HandleErrorResponse(c, fiber.StatusUnauthorized, "missing token", fmt.Errorf("authorization header not found"))
		}

		strArr := strings.Split(bearerToken, " ")
		if len(strArr) != 2 || !strings.EqualFold(strArr[0], "Bearer") {
			return r.HandleErrorResponse(c, fiber.StatusForbidden, "invalid token format", fmt.Errorf("wrong format"))
		}
		accessToken := strArr[1]
		permission, b := r.permissions.GetPermission(c.Method(), c.OriginalURL())
		if !b {
			r.log.Error("permission not found", c.OriginalURL())
			return r.HandleErrorResponse(c, fiber.StatusForbidden, "access denied", fmt.Errorf("no permission"))
		}
		data, statusCode, err := makeHasAccessRequest(r.cfg.ServiceUrl.Auth, accessToken, permission.Name)
		if err != nil {
			return r.HandleErrorResponse(c, fiber.StatusInternalServerError, "auth service error", err)
		}

		if statusCode >= 300 {
			var result auth_service_entity.HasErrorModel
			if err := json.Unmarshal(data, &result); err != nil {
				return r.HandleErrorResponse(c, fiber.StatusInternalServerError, "invalid response format", err)
			}
			return r.HandleErrorResponse(c, statusCode, "forbidden", fmt.Errorf("%s", result.Error))
		}

		var result auth_service_entity.HasAccessModel
		if err := json.Unmarshal(data, &result); err != nil {
			return r.HandleErrorResponse(c, fiber.StatusInternalServerError, "invalid response format", err)
		}

		if !result.HasAccess {
			return r.HandleErrorResponse(c, fiber.StatusForbidden, "access denied", fmt.Errorf("no permission"))
		}

		c.Request().Header.Set("user_id", result.UserID)
		c.Request().Header.Set("user_role", result.UserRole)
		c.Request().Header.Set("user_role_id", result.RoleID)

		return c.Next()
	}
}

func makeHasAccessRequest(authURL, accessToken, permName string) (data []byte, statusCode int, err error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/auth/has-access?perm_name=%s", authURL, permName), nil)
	if err != nil {
		return nil, 0, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	return data, resp.StatusCode, nil
}
