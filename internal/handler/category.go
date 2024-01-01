package handler

import (
	"spun/internal/service"

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

func (h *CategoryHandler) ViewCategory(c echo.Context) error {
	req := new(requestViewDefault)
	if err := c.Bind(req); err != nil {
		return err
	}

	resp := new(Response)
	category, err := h.service.ViewCategory(req.ID)
	if err != nil {
		resp.Set(0, nil, err)
	} else {
		resp.Success(category, nil)
	}
	return FormatResponse(c, resp)
}
