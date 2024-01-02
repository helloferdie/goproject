package libmiddleware

import (
	"errors"
	"fmt"
	"net/http"
	"spun/pkg/liblogger"
	"spun/pkg/libsession"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

type jwtClaims struct {
	libsession.SessionJWT
	jwt.RegisteredClaims
}

func JWT(secret []byte) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			headerToken := c.Request().Header.Get("Authorization")
			if headerToken == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "common.error.request.jwt.required")
			}

			splitToken := strings.Split(headerToken, " ")
			tokenString := splitToken[1]
			token, err := jwt.ParseWithClaims(tokenString, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return secret, nil
			})
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "common.error.request.jwt.invalid")
			}

			if token.Valid {
				claims, ok := token.Claims.(*jwtClaims)
				if !ok {
					return echo.NewHTTPError(http.StatusUnauthorized, "common.error.request.jwt.claims")
				}

				session, exist := libsession.FromContext(c.Request().Context())
				if !exist {
					// To avoid this error, make sure you have applied middleware 'Session' before 'JWT'
					return echo.NewHTTPError(http.StatusInternalServerError, "common.error.request.jwt.session")
				}

				// Assign claims to session
				session.SessionJWT = claims.SessionJWT
				return next(c)
			}

			if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
				return echo.NewHTTPError(http.StatusUnauthorized, "common.error.request.jwt.expired")
			}
			return echo.NewHTTPError(http.StatusUnauthorized, "common.error.request.jwt.invalid")
		}
	}
}
