package service

import (
	"context"
	"spun/internal/model"
	"spun/internal/repository"
	"spun/pkg/liberror"
	"spun/pkg/libvalidator"
)

// CountryService provides methods to interact with the country repository
type CountryService struct {
	repo repository.CountryRepository
}

// NewCountryService creates a new instance of CountryService
func NewCountryService(repo repository.CountryRepository) *CountryService {
	return &CountryService{
		repo: repo,
	}
}

// ListCountryParam
type ListCountryParam struct {
	PaginationParam
}

// ListCountry return list of countries from repository, totalItems, totalPages and error
func (s *CountryService) ListCountry(ctx context.Context, param *ListCountryParam) ([]*model.Country, int64, int64, *liberror.Error) {
	if errParam := libvalidator.Validate(param); errParam != nil {
		return nil, 0, 0, errParam
	}

	list, totalItems, err := s.repo.List(nil, param.PaginationParam.ToModel())
	if err != nil {
		return nil, 0, 0, liberror.NewServerError(err.Error())
	}

	return list, totalItems, GetTotalPages(totalItems, param.PageSize), nil
}

// CheckCountryExist
func (s *CountryService) CheckCountryExist(id int64) (*model.Country, *liberror.Error) {
	country, err := s.repo.GetByID(id)
	if err != nil {
		return nil, liberror.NewServerError(err.Error())
	}

	if country == nil {
		return nil, liberror.NewErrNotFound("")
	}

	return country, nil
}

// ViewCountryParam
type ViewCountryParam struct {
	IDParam
}

// ViewCountry
func (s *CountryService) ViewCountry(ctx context.Context, param *ViewCountryParam) (*model.Country, *liberror.Error) {
	if errParam := libvalidator.Validate(param); errParam != nil {
		return nil, errParam
	}

	return s.CheckCountryExist(param.ID)
}
