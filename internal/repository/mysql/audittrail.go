package mysql

import (
	"goproj/internal/model"
	"goproj/internal/repository"
	"goproj/pkg/libdb"

	"github.com/jmoiron/sqlx"
)

// MySQLAuditTrailRepository
type MySQLAuditTrailRepository struct {
	DB     *sqlx.DB
	Config *libdb.Config
}

// NewMySQLAuditTrailRepository
func NewMySQLAuditTrailRepository(db *sqlx.DB) repository.AuditTrailRepository {
	return &MySQLAuditTrailRepository{
		DB:     db,
		Config: libdb.NewConfig(model.AuditTrail{}, "audit_trail", true),
	}
}

// Create
func (repo *MySQLAuditTrailRepository) Create(audit *model.AuditTrail) error {
	_, _, err := libdb.Exec(repo.DB, repo.Config.QueryInsert, audit)
	return parseError(err)
}
