package handler

import (
	"spun/internal/dto"
	"spun/internal/service"
	"spun/pkg/libsession"

	"github.com/labstack/echo/v4"
)

// CategoryHandler handles HTTP requests related to category operations.
type CategoryHandler struct {
	service *service.CategoryService
}

// NewCategoryHandler creates a new instance of CategoryHandler
func NewCategoryHandler(service *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		service: service,
	}
}

// CreateCategory handles the HTTP request for creating a new category
func (h *CategoryHandler) CreateCategory(c echo.Context) error {
	req := new(service.CreateCategoryParam)
	if err := c.Bind(req); err != nil {
		return err
	}

	ctx := c.Request().Context()
	resp := new(Response)
	category, err := h.service.CreateCategory(ctx, req)
	if err != nil {
		resp.Set(0, nil, err)
	} else {
		session, _ := libsession.FromContext(ctx)
		resp.Success(dto.Category(session, category), nil)
	}

	return FormatResponse(c, resp)
}

// ViewCategory handles the HTTP request for viewing an existing category
func (h *CategoryHandler) ViewCategory(c echo.Context) error {
	req := new(service.ViewCategoryParam)
	if err := c.Bind(req); err != nil {
		return err
	}

	ctx := c.Request().Context()
	resp := new(Response)
	category, err := h.service.ViewCategory(ctx, req)
	if err != nil {
		resp.Set(0, nil, err)
	} else {
		session, _ := libsession.FromContext(ctx)
		resp.Success(dto.Category(session, category), nil)
	}
	return FormatResponse(c, resp)
}

// UpdateCategory handles the HTTP request for updating an existing category
func (h *CategoryHandler) UpdateCategory(c echo.Context) error {
	req := new(service.UpdateCategoryParam)
	if err := c.Bind(req); err != nil {
		return err
	}

	ctx := c.Request().Context()
	resp := new(Response)
	category, err := h.service.UpdateCategory(ctx, req)
	if err != nil {
		resp.Set(0, nil, err)
	} else {
		session, _ := libsession.FromContext(ctx)
		resp.Success(dto.Category(session, category), nil)
	}
	return FormatResponse(c, resp)
}

// DeleteCategory handles the HTTP request for deleting an existing category
func (h *CategoryHandler) DeleteCategory(c echo.Context) error {
	req := new(service.DeleteCategoryParam)
	if err := c.Bind(req); err != nil {
		return err
	}

	ctx := c.Request().Context()
	resp := new(Response)
	err := h.service.DeleteCategory(ctx, req)
	if err != nil {
		resp.Set(0, nil, err)
	} else {
		resp.Success(nil, nil)
	}
	return FormatResponse(c, resp)
}
