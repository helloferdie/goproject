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

func (h *CategoryHandler) ViewCategory(c echo.Context) error {
	req := new(requestViewDefault)
	if err := c.Bind(req); err != nil {
		return err
	}

	// ctx := c.Request().Context()
	// fmt.Println(ctx)

	// ctx = libsession.NewContext(c.Request().Context())
	// fmt.Println(ctx)
	//context.WithValue()

	//ctx :=
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	ctx := c.Request().Context()
	resp := new(Response)
	category, err := h.service.ViewCategory(ctx, req.ID)
	if err != nil {
		resp.Set(0, nil, err)
	} else {
		session, _ := libsession.FromContext(ctx)
		resp.Success(dto.Category(session, category), nil)
	}
	return FormatResponse(c, resp)
}
