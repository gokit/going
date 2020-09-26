package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/gokit/going"
	"github.com/gokit/going/bizapp"
)

// Context put database and request in r.Context
func Biz(biz bizapp.Biz) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			ctx.SetRequest(ctx.Request().WithContext(going.WithBiz(ctx.Request().Context(), biz)))
			return next(ctx)
		}
	}
}
