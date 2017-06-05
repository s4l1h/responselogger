package responselogger

import (
	"github.com/labstack/echo"
)

// EchoMiddleware echo middleware function
func EchoMiddleware(fn CallBackFN) echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			responseLogger := New(c.Response().Writer, fn)
			c.Response().Writer = responseLogger
			return h(c)
		}
	}
}
