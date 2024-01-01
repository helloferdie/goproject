package mysql

import (
	"fmt"
	"spun/internal/model"
	"spun/internal/repository"
	"spun/pkg/libdb"

	"github.com/jmoiron/sqlx"
)

type MySQLCategoryRepository struct {
	DB     *sqlx.DB
	Config *libdb.Config
}

func NewMySQLCategoryRepository(db *sqlx.DB) repository.CategoryRepository {
	return &MySQLCategoryRepository{
		DB:     db,
		Config: libdb.NewConfig(model.Category{}, "category", true),
	}
}

func (*MySQLCategoryRepository) Create(category *model.Category) (*model.Category, error) {
	panic("unimplemented")
}

func (repo *MySQLCategoryRepository) GetByID(id int64) (*model.Category, error) {
	category := new(model.Category)
	exist, err := libdb.Get(repo.DB, category, fmt.Sprintf(repo.Config.QuerySelect, "AND id = :id"), map[string]interface{}{
		"id": id,
	})

	if err == nil && exist {
		return category, nil
	}

	return nil, err

	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(exist)
	// return category, err
}
