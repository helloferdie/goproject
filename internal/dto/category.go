package dto

import (
	"goproj/internal/model"
	"goproj/pkg/libsession"
	"goproj/pkg/libtime"
)

func Category(s *libsession.Session, category *model.Category) map[string]interface{} {
	return map[string]interface{}{
		"id":          category.ID,
		"name":        category.Name,
		"description": category.Description,
		"is_active":   category.IsActive,
		"created_at":  libtime.NullFormat(s.Timezone, category.CreatedAt),
		"updated_at":  libtime.NullFormat(s.Timezone, category.UpdatedAt),
		"deleted_at":  libtime.NullFormat(s.Timezone, category.DeletedAt),
	}
}

func CategoryList(s *libsession.Session, list []*model.Category, totalItems int64, totalPages int64) map[string]interface{} {
	items := make([]interface{}, len(list))
	for k, v := range list {
		items[k] = Category(s, v)
	}

	return map[string]interface{}{
		"total_items": totalItems,
		"total_pages": totalPages,
		"items":       items,
	}
}
