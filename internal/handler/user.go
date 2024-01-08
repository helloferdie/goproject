package handler

import (
	"spun/internal/dto"
	"spun/internal/service"
	"spun/pkg/libsession"

	"github.com/labstack/echo/v4"
)

// UserHandler handles HTTP requests related to user operations.
type UserHandler struct {
	service *service.UserService
}

// NewUserHandler creates a new instance of UserHandler
func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

// ListUser handles the HTTP request for deleting an existing user
func (h *UserHandler) ListUser(c echo.Context) error {
	req := new(service.ListUserParam)
	if err := c.Bind(req); err != nil {
		return err
	}

	ctx := c.Request().Context()
	resp := new(Response)
	list, totalItems, totalPages, err := h.service.ListUser(ctx, req)
	if err != nil {
		resp.Set(0, nil, err)
	} else {
		session, _ := libsession.FromContext(ctx)
		resp.Success(dto.UserList(session, list, totalItems, totalPages), nil)
	}
	return FormatResponse(c, resp)
}

// ViewUser handles the HTTP request for viewing an existing user
func (h *UserHandler) ViewUser(c echo.Context) error {
	req := new(service.ViewUserParam)
	if err := c.Bind(req); err != nil {
		return err
	}

	ctx := c.Request().Context()
	resp := new(Response)
	user, err := h.service.ViewUser(ctx, req)
	if err != nil {
		resp.Set(0, nil, err)
	} else {
		session, _ := libsession.FromContext(ctx)
		resp.Success(dto.User(session, user), nil)
	}
	return FormatResponse(c, resp)
}
