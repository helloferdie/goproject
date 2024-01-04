package handler

import (
	"spun/internal/dto"
	"spun/internal/service"
	"spun/pkg/libsession"

	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	service *service.CategoryService
}

func NewCategoryHandler(service *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		service: service,
	}
}

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
