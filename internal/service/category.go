package service

import (
	"context"
	"spun/internal/model"
	"spun/internal/repository"
	"spun/pkg/liberror"
	"spun/pkg/libvalidator"

	"github.com/jinzhu/copier"
)

// CategoryService provides methods to interact with the category repository
type CategoryService struct {
	repo     repository.CategoryRepository
	svcAudit *AuditTrailService
}

// NewCategoryService creates a new instance of CategoryService
func NewCategoryService(repo repository.CategoryRepository, svc *AuditTrailService) *CategoryService {
	return &CategoryService{
		repo:     repo,
		svcAudit: svc,
	}
}

// CreateCategoryParam
type CreateCategoryParam struct {
	Name        string `json:"name" loc:"common.name title" validate:"required"`
	Description string `json:"description" loc:"common.description title" validate:"required"`
}

// CreateCategory
func (s *CategoryService) CreateCategory(ctx context.Context, param *CreateCategoryParam) (*model.Category, *liberror.Error) {
	if errParam := libvalidator.Validate(param); errParam != nil {
		return nil, errParam
	}

	category := new(model.Category)
	category.Name = param.Name
	category.Description = param.Description
	category, err := s.repo.Create(category)
	if err != nil {
		return nil, liberror.NewServerError(err.Error())
	}
	logAuditTrail(s.svcAudit, ctx, "category", category.ID, "create", category, "")
	return category, nil
}

// CheckCategoryExist
func (s *CategoryService) CheckCategoryExist(id int64) (*model.Category, *liberror.Error) {
	category, err := s.repo.GetByID(id)
	if err != nil {
		return nil, liberror.NewServerError(err.Error())
	}

	if category == nil {
		return nil, liberror.NewErrNotFound("")
	}

	return category, nil
}

// ViewCategoryParam
type ViewCategoryParam struct {
	IDParam
}

// ViewCategory
func (s *CategoryService) ViewCategory(ctx context.Context, param *ViewCategoryParam) (*model.Category, *liberror.Error) {
	if errParam := libvalidator.Validate(param); errParam != nil {
		return nil, errParam
	}

	return s.CheckCategoryExist(param.ID)
}

// UpdateCategoryParam
type UpdateCategoryParam struct {
	IDParam
	CreateCategoryParam
}

// UpdateCategory
func (s *CategoryService) UpdateCategory(ctx context.Context, param *UpdateCategoryParam) (*model.Category, *liberror.Error) {
	if errParam := libvalidator.Validate(param); errParam != nil {
		return nil, errParam
	}

	category, errCheck := s.CheckCategoryExist(param.ID)
	if errCheck != nil {
		return nil, errCheck
	}

	newCategory := &model.Category{}
	copier.Copy(newCategory, category)
	newCategory.Name = param.Name
	newCategory.Description = param.Description

	category, changes, err := s.repo.Update(category.ID, category, newCategory)
	if err != nil {
		return nil, liberror.NewServerError(err.Error())
	}
	logAuditTrail(s.svcAudit, ctx, "category", category.ID, "update", changes, "")
	return category, nil
}

// DeleteCategoryParam
type DeleteCategoryParam struct {
	IDParam
}

// DeleteCategory
func (s *CategoryService) DeleteCategory(ctx context.Context, param *DeleteCategoryParam) *liberror.Error {
	if errParam := libvalidator.Validate(param); errParam != nil {
		return errParam
	}

	category, errCheck := s.CheckCategoryExist(param.ID)
	if errCheck != nil {
		return errCheck
	}

	err := s.repo.Delete(category.ID)
	if err != nil {
		return liberror.NewServerError(err.Error())
	}
	logAuditTrail(s.svcAudit, ctx, "category", category.ID, "delete", nil, "")
	return nil
}

// ListCategoryParam
type ListCategoryParam struct {
	PaginationParam
}

// ListCategory return list of categories from repository, totalItems, totalPages and error
func (s *CategoryService) ListCategory(ctx context.Context, param *ListCategoryParam) ([]*model.Category, int64, int64, *liberror.Error) {
	if errParam := libvalidator.Validate(param); errParam != nil {
		return nil, 0, 0, errParam
	}

	list, totalItems, err := s.repo.List(nil, param.PaginationParam.ToModel())
	if err != nil {
		return nil, 0, 0, liberror.NewServerError(err.Error())
	}

	return list, totalItems, GetTotalPages(totalItems, param.PageSize), nil
}
