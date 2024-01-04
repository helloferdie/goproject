package handler

import (
	"spun/pkg/liblocale"

	"github.com/labstack/echo/v4"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var localeBundle *i18n.Bundle
var localeInit = false

// initLocalization
func initLocalization() {
	if !localeInit {
		localeBundle, _ = liblocale.LoadBundle()
		localeInit = true
	}
}

// GetLocalizer
func GetLocalizer(c echo.Context) *i18n.Localizer {
	initLocalization()

	language := c.Request().Header.Get("Accept-Language")
	return i18n.NewLocalizer(localeBundle, language)
}
