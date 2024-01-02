package libmiddleware

import (
	"fmt"
	"net/http"
	"spun/pkg/liblogger"
	"spun/pkg/libsession"
	"time"

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

// Session -
func Session(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session := new(libsession.Session)

		header := c.Request().Header
		session.Language = header.Get("Accept-Language")

		tz := header.Get("Accept-Timezone")
		loc, err := time.LoadLocation(tz)
		if err != nil {
			loc, _ = time.LoadLocation("UTC")
		}
		session.Timezone = loc

		ip := c.Request().Header.Get("X-Real-Ip")
		if ip == "" {
			ip = c.Request().RemoteAddr
		}
		session.IPAddress = ip

		ctx := libsession.NewContext(c.Request().Context(), session)
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}

// JWT -
func JWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		headerToken := c.Request().Header.Get("Authorization")
		if headerToken == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "common.error.request.jwt.required")
		}

		fmt.Println("JWT")
		return next(c)
	}
}
