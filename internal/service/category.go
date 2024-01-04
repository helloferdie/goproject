package service

import (
	"context"
	"fmt"
	"spun/internal/model"
	"spun/internal/repository"
	"spun/pkg/liberror"
	"spun/pkg/libsession"
	"spun/pkg/libvalidator"
)

type CategoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

type CreateCategoryParam struct {
	Name        string `json:"name" loc:"common.name title" validate:"required"`
	Description string `json:"description" loc:"common.description title" validate:"required"`
}

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

type ViewCategoryParam struct {
	ID int64 `json:"id"`
}

func (s *CategoryService) ViewCategory(ctx context.Context, param *ViewCategoryParam) (*model.Category, *liberror.Error) {
	fmt.Println("Read context - Start")
	fmt.Println(ctx)
	fmt.Println("Read context - Done")

	session, _ := libsession.FromContext(ctx)
	if session != nil {
		fmt.Println("Session found " + session.Timezone.String())
		fmt.Println("Session found " + session.Language)
		fmt.Printf("Session found user ID %v\n", session.UserID)
		fmt.Printf("Session found role ID %v", session.RoleID)
	}

	if param.ID == 0 {
		fieldErrors := []*liberror.Base{
			{Error: "common.error.validation.required",
				Field:     "id",
				ErrorVars: map[string]string{"field": "{{common.id.other upper}}"}},
			{Error: "common.error.validation.email",
				Field:     "email",
				ErrorVars: map[string]string{"field": "{{common.email title}}"}},
		}
		return nil, liberror.NewErrValidation(fieldErrors...)
	}

	category, err := s.repo.GetByID(param.ID)
	if err != nil {
		return nil, liberror.NewErrRepository()
	}

	if category == nil {
		return nil, liberror.NewErrNotFound("")
	}

	return category, nil
}
