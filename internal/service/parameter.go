package service

import (
	"goproj/internal/model"
	"math"
)

// IDParam
type IDParam struct {
	ID int64 `json:"id" loc:"common.id.other upper" validate:"required"`
}

// IDStringParam
type IDStringParam struct {
	ID string `json:"id" loc:"common.id.other upper" validate:"required"`
}

// SortField
type SortField struct {
	Field     string `json:"field" loc:"common.sort.field title" validate:"required"`
	Direction string `json:"direction" loc:"common.sort.field title"`
}

// PaginationParam
type PaginationParam struct {
	Page       int64       `json:"page" loc:"common.page title" validate:"required,numeric,min=1"`
	PageSize   int64       `json:"page_size" loc:"common.page_size title" validate:"required,numeric,min=1,max=500"`
	SortFields []SortField `json:"sort" validate:"dive"`
}

// GetTotalPages calculate total pages
func GetTotalPages(totalItems int64, pageSize int64) int64 {
	fTotalItems := float64(totalItems)
	fPageSize := float64(pageSize)
	return int64(math.Ceil(fTotalItems / fPageSize))
}

// ToModel return pagination param that compatible with repository pagination
func (p *PaginationParam) ToModel() *model.Pagination {
	mPagination := new(model.Pagination)
	mPagination.Cursor = (p.Page - 1) * p.PageSize
	mPagination.Limit = p.PageSize

	sortFields := len(p.SortFields)
	if sortFields > 0 {
		listSorts := make([]model.SortField, sortFields)
		for k, v := range p.SortFields {
			listSorts[k] = model.SortField{
				Field:     v.Field,
				Direction: v.Direction,
			}
		}
		mPagination.SortFields = listSorts
	}
	return mPagination
}
