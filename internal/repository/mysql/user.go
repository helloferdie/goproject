package mysql

import (
	"fmt"
	"spun/internal/model"
	"spun/internal/repository"
	"spun/pkg/libdb"

	"github.com/jmoiron/sqlx"
)

// MySQLUserRepository struct for the user repository with MySQL
type MySQLUserRepository struct {
	DB     *sqlx.DB
	Config *libdb.Config
}

// NewMySQLUserRepository creates a new instance of UserRepository for MySQL
func NewMySQLUserRepository(db *sqlx.DB) repository.UserRepository {
	return &MySQLUserRepository{
		DB:     db,
		Config: libdb.NewConfig(model.User{}, "user", true),
	}
}

// List returns list of users include with pagination
func (repo *MySQLUserRepository) List(filter map[string]interface{}, pagination *model.Pagination) ([]*model.User, int64, error) {
	// Parse condition
	condition := ""
	values := map[string]interface{}{}

	// Query total
	total := new(model.Total)
	queryTotal := fmt.Sprintf(repo.Config.QueryTotal, condition)
	_, err := libdb.Get(repo.DB, total, queryTotal, values)
	if err != nil {
		return nil, 0, parseError(err)
	}

	//repo.Config.

	// Pagination
	querySelect := fmt.Sprintf(repo.Config.QuerySelect, condition)
	querySelect, values = applySortFields(querySelect, values, pagination, "ORDER BY created_at DESC")

	// Query list
	list := []*model.User{}
	err = libdb.Select(repo.DB, &list, querySelect, values)
	return list, total.Total, parseError(err)
}

// GetByID retrieves a user from the database based on the given ID
func (repo *MySQLUserRepository) GetByID(id int64) (*model.User, error) {
	user := new(model.User)
	querySelect := fmt.Sprintf(repo.Config.QuerySelect, "AND id = :id ") + " LIMIT 1"
	exist, err := libdb.Get(repo.DB, user, querySelect, map[string]interface{}{
		"id": id,
	})

	if err == nil && exist {
		return user, nil
	}
	return nil, parseError(err)
}
