package mysql

import (
	"fmt"
	"spun/internal/model"
	"spun/internal/repository"
	"spun/pkg/libdb"

	"github.com/jmoiron/sqlx"
)

// MySQLCountryRepository struct for the country repository with MySQL
type MySQLCountryRepository struct {
	DB     *sqlx.DB
	Config *libdb.Config
}

// NewMySQLCountryRepository creates a new instance of CountryRepository for MySQL
func NewMySQLCountryRepository(db *sqlx.DB) repository.CountryRepository {
	return &MySQLCountryRepository{
		DB:     db,
		Config: libdb.NewConfig(model.Country{}, "country", true),
	}
}

// List returns list of countries include with pagination
func (repo *MySQLCountryRepository) List(filter map[string]interface{}, pagination *model.Pagination) ([]*model.Country, int64, error) {
	// Parse condition
	condition := ""
	values := map[string]interface{}{}

	// Query total
	total := new(model.Total)
	queryTotal := fmt.Sprintf(repo.Config.QueryTotal, condition)
	_, err := libdb.Get(repo.DB, total, queryTotal, values)
	if err != nil {
		return nil, 0, err
	}

	// Pagination
	querySelect := fmt.Sprintf(repo.Config.QuerySelect, condition)
	querySelect, values = applySortFields(querySelect, values, pagination, "ORDER BY created_at DESC")

	// Query list
	list := []*model.Country{}
	err = libdb.Select(repo.DB, &list, querySelect, values)
	return list, total.Total, err
}

// GetByID retrieves a country from the database based on the given ID
func (repo *MySQLCountryRepository) GetByID(id int64) (*model.Country, error) {
	country := new(model.Country)
	querySelect := fmt.Sprintf(repo.Config.QuerySelect, "AND id = :id ") + " LIMIT 1"
	exist, err := libdb.Get(repo.DB, country, querySelect, map[string]interface{}{
		"id": id,
	})

	if err == nil && exist {
		return country, nil
	}
	return nil, err
}
