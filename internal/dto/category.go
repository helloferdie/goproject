package dto

import (
	"spun/internal/model"
	"spun/pkg/libsession"
	"spun/pkg/libtime"
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
