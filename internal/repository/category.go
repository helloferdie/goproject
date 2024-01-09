package repository

import "goproj/internal/model"

type CategoryRepository interface {
	Create(category *model.Category) (*model.Category, error)
	Update(id int64, original, modified *model.Category) (*model.Category, map[string]interface{}, error)
	Delete(id int64) error
	List(filter map[string]interface{}, pagination *model.Pagination) ([]*model.Category, int64, error)
	GetByID(id int64) (*model.Category, error)
}
