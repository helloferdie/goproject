package handler

import (
	"spun/internal/dto"
	"spun/internal/service"
	"spun/pkg/libsession"

	"github.com/labstack/echo/v4"
)

// CountryHandler handles HTTP requests related to country operations.
type CountryHandler struct {
	service *service.CountryService
}

// NewCountryHandler creates a new instance of CountryHandler
func NewCountryHandler(service *service.CountryService) *CountryHandler {
	return &CountryHandler{
		service: service,
	}
}

// ListCountry handles the HTTP request for deleting an existing country
func (h *CountryHandler) ListCountry(c echo.Context) error {
	req := new(service.ListCountryParam)
	if err := c.Bind(req); err != nil {
		return err
	}

	ctx := c.Request().Context()
	resp := new(Response)
	list, totalItems, totalPages, err := h.service.ListCountry(ctx, req)
	if err != nil {
		resp.Set(0, nil, err)
	} else {
		session, _ := libsession.FromContext(ctx)
		resp.Success(dto.CountryList(session, list, totalItems, totalPages), nil)
	}
	return FormatResponse(c, resp)
}

// ViewCountry handles the HTTP request for viewing an existing country
func (h *CountryHandler) ViewCountry(c echo.Context) error {
	req := new(service.ViewCountryParam)
	if err := c.Bind(req); err != nil {
		return err
	}

	ctx := c.Request().Context()
	resp := new(Response)
	country, err := h.service.ViewCountry(ctx, req)
	if err != nil {
		resp.Set(0, nil, err)
	} else {
		session, _ := libsession.FromContext(ctx)
		resp.Success(dto.Country(session, country), nil)
	}
	return FormatResponse(c, resp)
}
