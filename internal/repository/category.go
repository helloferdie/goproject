package repository

import "spun/internal/model"

type CategoryRepository interface {
	Create(category *model.Category) (*model.Category, error)
	GetByID(id int64) (*model.Category, error)
}
