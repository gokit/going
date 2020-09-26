package middlewares

import (
	"github.com/gokit/going"
	"github.com/gokit/going/db"
	"github.com/labstack/echo/v4"
)

// Context put database and request in r.Context
func Database(d *db.Database) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			ctx.SetRequest(ctx.Request().WithContext(going.WithDatabase(ctx.Request().Context(), d)))
			return next(ctx)
		}
	}
}
