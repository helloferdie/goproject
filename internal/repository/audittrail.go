package repository

import "spun/internal/model"

type AuditTrailRepository interface {
	Create(audit *model.AuditTrail) error
}
