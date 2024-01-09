package dto

import (
	"goproj/internal/model"
	"goproj/pkg/libsession"
	"goproj/pkg/libtime"
)

func User(s *libsession.Session, user *model.User) map[string]interface{} {
	return map[string]interface{}{
		"id":         user.ID,
		"uuid":       user.UUID,
		"account_no": user.AccountNo,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"email":      user.Email,
		"phone":      user.Phone,
		"created_at": libtime.NullFormat(s.Timezone, user.CreatedAt),
		"updated_at": libtime.NullFormat(s.Timezone, user.UpdatedAt),
		"deleted_at": libtime.NullFormat(s.Timezone, user.DeletedAt),
	}
}

func UserList(s *libsession.Session, list []*model.User, totalItems int64, totalPages int64) map[string]interface{} {
	items := make([]interface{}, len(list))
	for k, v := range list {
		items[k] = User(s, v)
	}

	return map[string]interface{}{
		"total_items": totalItems,
		"total_pages": totalPages,
		"items":       items,
	}
}
