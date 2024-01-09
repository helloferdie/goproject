package service

import (
	"context"
	"goproj/internal/model"
	"goproj/internal/repository"
	"goproj/pkg/liberror"
	"goproj/pkg/libvalidator"
)

// UserService provides methods to interact with the user repository
type UserService struct {
	repo repository.UserRepository
}

// NewUserService creates a new instance of UserService
func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// ListUserParam
type ListUserParam struct {
	PaginationParam
}

// ListUser return list of users from repository, totalItems, totalPages and error
func (s *UserService) ListUser(ctx context.Context, param *ListUserParam) ([]*model.User, int64, int64, *liberror.Error) {
	if errParam := libvalidator.Validate(param); errParam != nil {
		return nil, 0, 0, errParam
	}

	list, totalItems, err := s.repo.List(nil, param.PaginationParam.ToModel())
	if err != nil {
		return nil, 0, 0, liberror.NewServerError(err.Error())
	}

	return list, totalItems, GetTotalPages(totalItems, param.PageSize), nil
}

// CheckUserExist
func (s *UserService) CheckUserExist(id int64) (*model.User, *liberror.Error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, liberror.NewServerError(err.Error())
	}

	if user == nil {
		return nil, liberror.NewErrNotFound("")
	}

	return user, nil
}

// ViewUserParam
type ViewUserParam struct {
	IDParam
}

// ViewUser
func (s *UserService) ViewUser(ctx context.Context, param *ViewUserParam) (*model.User, *liberror.Error) {
	if errParam := libvalidator.Validate(param); errParam != nil {
		return nil, errParam
	}

	return s.CheckUserExist(param.ID)
}
