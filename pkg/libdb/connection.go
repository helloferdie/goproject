package libdb

import (
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Connection -
type Connection struct {
	Driver string
	DSN    string
}

// setConnection - Set connection string
func setConnection(env string) (*Connection, error) {
	if env == "" {
		env = "db"
	}

	driver := os.Getenv(env + "_driver")
	host := os.Getenv(env + "_host")
	port := os.Getenv(env + "_port")
	user := os.Getenv(env + "_user")
	pass := os.Getenv(env + "_pass")
	dbname := os.Getenv(env + "_name")

	if driver == "mysql" {
		cfg := mysql.NewConfig()
		cfg.Net = "tcp"
		cfg.Addr = host + ":" + port
		cfg.User = user
		cfg.Passwd = pass
		cfg.DBName = dbname
		cfg.ParseTime = true
		cfg.Params = map[string]string{
			"charset": "utf8mb4",
		}

		return &Connection{
			Driver: driver,
			DSN:    cfg.FormatDSN(),
		}, nil
	}
	return nil, fmt.Errorf("Database driver for %s not supported.", driver)
}

// Open - Open connection with default retry parameter
func Open(env string) (*sqlx.DB, error) {
	return OpenWithRetry(env, 3)
}

// OpenWithRetry - Open connection with custom retry parameter
func OpenWithRetry(env string, maxRetry int) (*sqlx.DB, error) {
	conn, err := setConnection(env)
	if err != nil {
		return nil, fmt.Errorf("Error set connection string: %v", err)
	}

	if maxRetry < 0 {
		maxRetry = 0
	}

	db, err := sqlx.Connect(conn.Driver, conn.DSN)
	if err != nil {
		return nil, err
	}
	return db, err
}
