package middlewares

import (
	"github.com/gokit/going"
	"github.com/gokit/going/logger"
	"github.com/labstack/echo/v4"
)

// Context put database and request in r.Context
func Logger(logger *logger.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			ctx.SetRequest(ctx.Request().WithContext(going.WithLogger(ctx.Request().Context(), logger)))
			return next(ctx)
		}
	}
}
