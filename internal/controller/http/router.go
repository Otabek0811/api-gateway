// Package http implements routing paths. Each service is in its own file.
package http

import (
	"GolandProjects/api-gateway/config"
	"GolandProjects/api-gateway/internal/controller/http/middleware"
	v1 "GolandProjects/api-gateway/internal/controller/http/v1"
	"GolandProjects/api-gateway/pkg/logger"
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/swagger"
	"net/http"
	"time"
)

// NewRouter -.
// Swagger spec:
// @title       Go Api Gateway
// @description api gateway - go fiber
// @version     1.0
// @host        localhost:8080
// @BasePath    /api/v1
func NewRouter(app *fiber.App, cfg *config.Config, log logger.Interface) {
	permissions := config.NewPermissions()

	// Options
	app.Use(middleware.Logger(log))
	app.Use(middleware.Recovery(log))

	app.Use(limiter.New(limiter.Config{
		Max:               30,
		Expiration:        10 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))

	// Prometheus metrics
	if cfg.Metrics.Enabled {
		prometheus := fiberprometheus.New("Api Gateway")
		prometheus.RegisterAt(app, "/metrics")
		app.Use(prometheus.Middleware)
	}

	// Swagger
	if cfg.Swagger.Enabled {
		app.Get("/swagger/*", swagger.HandlerDefault)
	}

	// K8s
	app.Get("/health", func(ctx *fiber.Ctx) error { return ctx.SendStatus(http.StatusOK) })

	// Routers
	apiV1Group := app.Group("/api/v1")
	{
		v1.NewRoutes(apiV1Group, log, cfg, permissions)
	}
}
