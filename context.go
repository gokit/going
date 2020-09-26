package going

import (
	"context"
	"github.com/gokit/going/bizapp"
	"github.com/gokit/going/db"
	"github.com/gokit/going/logger"
	"github.com/labstack/echo/v4"
)

const (
	keyBiz      string = "going.biz"
	keyLogger   string = "going.logger"
	keyDatabase string = "going.database"
)

type WebContext echo.Context

type Ctx interface {
	context.Context
	WebContext() WebContext
	Logger() *logger.Logger
	DB() *db.Database
	Biz() bizapp.Biz
}

type appContext struct {
	context.Context
	ctx echo.Context
}

func (c *appContext) WebContext() WebContext {
	return c.ctx
}

func (c *appContext) Logger() *logger.Logger {
	return c.Value(keyLogger).(*logger.Logger)
}

func (c *appContext) DB() *db.Database {
	return c.Value(keyDatabase).(*db.Database)
}

func (c *appContext) Biz() bizapp.Biz {
	return c.Value(keyBiz).(bizapp.Biz)
}

func Context(ctx WebContext) Ctx {
	return &appContext{ctx.Request().Context(), ctx}
}

// WithLogger put logger in context
func WithLogger(ctx context.Context, logger *logger.Logger) context.Context {
	return context.WithValue(ctx, keyLogger, logger)
}

// WithDatabase put database in context
func WithDatabase(ctx context.Context, database *db.Database) context.Context {
	return context.WithValue(ctx, keyDatabase, database)
}

// WithDatabase put biz in context
func WithBiz(ctx context.Context, biz bizapp.Biz) context.Context {
	return context.WithValue(ctx, keyBiz, biz)
}
