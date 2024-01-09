package repository

import "goproj/internal/model"

type CountryRepository interface {
	List(filter map[string]interface{}, pagination *model.Pagination) ([]*model.Country, int64, error)
	GetByID(id int64) (*model.Country, error)
}
