package mysql

import (
	"fmt"
	"spun/internal/model"
	"spun/internal/repository"
	"spun/pkg/libdb"

	"github.com/jmoiron/sqlx"
)

// MySQLCategoryRepository struct for the category repository with MySQL
type MySQLCategoryRepository struct {
	DB     *sqlx.DB
	Config *libdb.Config
}

// NewMySQLCategoryRepository creates a new instance of CategoryRepository for MySQL
func NewMySQLCategoryRepository(db *sqlx.DB) repository.CategoryRepository {
	return &MySQLCategoryRepository{
		DB:     db,
		Config: libdb.NewConfig(model.Category{}, "category", true),
	}
}

// Create inserts a new category into the database
func (repo *MySQLCategoryRepository) Create(category *model.Category) (*model.Category, error) {
	id, _, err := libdb.Exec(repo.DB, repo.Config.QueryInsert, category)
	if err != nil {
		return nil, err
	}
	return repo.GetByID(id)
}

// Update modifies an existing category
func (repo *MySQLCategoryRepository) Update(id int64, original, modified *model.Category) (*model.Category, map[string]interface{}, error) {
	changes := model.GetChanges(original, modified)
	querySet, queryValues := libdb.PrepareUpdateQuery(changes, map[string]interface{}{
		"id": id,
	})
	_, _, err := libdb.Exec(repo.DB, fmt.Sprintf(repo.Config.QueryUpdate, querySet, "AND id = :id"), queryValues)
	if err != nil {
		return nil, nil, err
	}

	category, err := repo.GetByID(id)
	return category, changes, err
}

// Delete removes a category from the database
func (repo *MySQLCategoryRepository) Delete(id int64) error {
	_, _, err := libdb.Exec(repo.DB, fmt.Sprintf(repo.Config.QueryDelete, "AND id = :id"), map[string]interface{}{
		"id": id,
	})
	return err
}

// GetByID retrieves a category from the database based on the given ID
func (repo *MySQLCategoryRepository) GetByID(id int64) (*model.Category, error) {
	category := new(model.Category)
	exist, err := libdb.Get(repo.DB, category, fmt.Sprintf(repo.Config.QuerySelect, "AND id = :id"), map[string]interface{}{
		"id": id,
	})

	if err == nil && exist {
		return category, nil
	}
	return nil, err
}
