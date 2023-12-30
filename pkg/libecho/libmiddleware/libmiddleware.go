package libmiddleware

import (
	"spun/pkg/liblogger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Logger -
func Logger() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:      true,
		LogStatus:   true,
		LogMethod:   true,
		LogRemoteIP: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			liblogger.Infow("request", "uri", v.URI, "method", v.Method, "status", v.Status, "ip", c.Request().RemoteAddr,
				"ip_real", c.Request().Header.Get("X-Real-IP"), "ip_proxy", c.Request().Header.Get("X-Proxy-IP"))
			return nil
		},
	})
}
