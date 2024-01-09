package repository

import "goproj/internal/model"

type UserRepository interface {
	List(filter map[string]interface{}, pagination *model.Pagination) ([]*model.User, int64, error)
	GetByID(id int64) (*model.User, error)
}
