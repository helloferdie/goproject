package service

import (
	"context"
	"fmt"
	"spun/internal/model"
	"spun/internal/repository"
	"spun/pkg/liberror"
	"spun/pkg/libsession"
)

type CategoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

func (s *CategoryService) ViewCategory(ctx context.Context, id int64) (*model.Category, *liberror.Error) {
	fmt.Println("Read context - Start")
	fmt.Println(ctx)
	fmt.Println("Read context - Done")

	session, _ := libsession.FromContext(ctx)
	if session != nil {
		fmt.Println("Session found " + session.Timezone.String())
		fmt.Println("Session found " + session.Language)
	}

	if id == 0 {
		fieldErrors := []*liberror.Base{
			{Error: "common.error.validation.required",
				ErrorVars: map[string]interface{}{"field": "Field1"}},
			{Error: "common.error.validation.email",
				ErrorVars: map[string]interface{}{"field": "Field2"}},
		}
		return nil, liberror.NewErrValidation(fieldErrors...)
	}

	category, err := s.repo.GetByID(id)
	if err != nil {
		return nil, liberror.NewErrRepository()
	}

	if category == nil {
		return nil, liberror.NewErrNotFound()
	}

	return category, nil
}
