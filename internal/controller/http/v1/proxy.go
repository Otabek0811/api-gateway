package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func (r *V1) MakeProxy(c *fiber.Ctx, proxyURL, endpoint string) error {
	fullURL := proxyURL + endpoint

	if err := proxy.Do(c, fullURL); err != nil {
		r.log.Error("proxy error", err)
		return err
	}

	c.Request().Header.Set("X-Forwarded-Host", string(c.Request().Host()))

	return nil
}
