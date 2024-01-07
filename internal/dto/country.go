package dto

import (
	"spun/internal/model"
	"spun/pkg/libsession"
	"spun/pkg/libtime"
)

func Country(s *libsession.Session, country *model.Country) map[string]interface{} {
	return map[string]interface{}{
		"id":         country.ID,
		"iso2":       country.ISO2,
		"iso3":       country.ISO3,
		"callcode":   country.Callcode,
		"name":       country.Name,
		"created_at": libtime.NullFormat(s.Timezone, country.CreatedAt),
		"updated_at": libtime.NullFormat(s.Timezone, country.UpdatedAt),
		"deleted_at": libtime.NullFormat(s.Timezone, country.DeletedAt),
	}
}

func CountryList(s *libsession.Session, list []*model.Country, totalItems int64, totalPages int64) map[string]interface{} {
	items := make([]interface{}, len(list))
	for k, v := range list {
		items[k] = Country(s, v)
	}

	return map[string]interface{}{
		"total_items": totalItems,
		"total_pages": totalPages,
		"items":       items,
	}
}
