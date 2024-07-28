package middleware

import (
	"context"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
)

type CustomContext struct {
	echo.Context
	StartTime time.Time
	ctx       context.Context
}

func (cc *CustomContext) Deadline() (deadline time.Time, ok bool) {
	return cc.ctx.Deadline()
}

func (cc *CustomContext) Done() <-chan struct{} {
	return cc.ctx.Done()
}

func (cc *CustomContext) Err() error {
	return cc.ctx.Err()
}

func (cc *CustomContext) Value(key interface{}) interface{} {
	return cc.ctx.Value(key)
}

func CustomContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &CustomContext{
			Context:   c,
			StartTime: time.Now(),
			ctx:       c.Request().Context(),
		}
		return next(cc)
	}
}

func LogMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc, ok := c.(*CustomContext)
		if !ok {
			cc = &CustomContext{Context: c, StartTime: time.Now()}
		}

		err := next(cc)

		elapsedTime := time.Since(cc.StartTime)
		formattedTime := time.Now().Format("2006-01-02 15:04:05")
		fmt.Printf("[%s] [%s] [%s] [%s]\n", formattedTime, cc.Request().Method, cc.Request().URL.Path, elapsedTime)

		return err
	}
}

func ParseFormMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := c.Request().ParseForm(); err != nil {
			return err
		}
		return next(c)
	}
}

func ParseMultipartFormMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := c.Request().ParseMultipartForm(10 << 20); err != nil {
			return err
		}
		return next(c)
	}
}
