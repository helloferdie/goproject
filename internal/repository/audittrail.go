package repository

import "goproj/internal/model"

type AuditTrailRepository interface {
	Create(audit *model.AuditTrail) error
}
