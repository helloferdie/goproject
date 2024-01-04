package service

import (
	"context"
	"fmt"
	"spun/internal/model"
	"spun/internal/repository"
	"spun/pkg/liberror"
	"spun/pkg/libvalidator"

	"github.com/jinzhu/copier"
)

// CategoryService provides methods to interact with the category repository
type CategoryService struct {
	repo repository.CategoryRepository
}

// NewCategoryService creates a new instance of CategoryService
func NewCategoryService(repo repository.CategoryRepository) *CategoryService {
	return &CategoryService{
		repo: repo,
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
		return nil, liberror.NewErrRepository()
	}
	return category, nil
}

// CheckCategoryExist
func (s *CategoryService) CheckCategoryExist(id int64) (*model.Category, *liberror.Error) {
	category, err := s.repo.GetByID(id)
	if err != nil {
		return nil, liberror.NewErrRepository()
	}

	if category == nil {
		return nil, liberror.NewErrNotFound("")
	}

	return category, nil
}

// ViewCategoryParam
type ViewCategoryParam struct {
	ID int64 `json:"id" loc:"common.id.other upper" validate:"required"`
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
	ID int64 `json:"id" loc:"common.id.other upper" validate:"required"`
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
		return nil, liberror.NewErrRepository()
	}

	fmt.Println(changes)
	return category, nil
}

// DeleteCategoryParam
type DeleteCategoryParam struct {
	ID int64 `json:"id" loc:"common.id.other upper" validate:"required"`
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
		return liberror.NewErrRepository()
	}
	return nil
}
