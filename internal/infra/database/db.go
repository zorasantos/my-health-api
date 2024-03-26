package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/zorasantos/my-health/config"
)

func ConnectDB() (*sql.DB, error) {
	var connStr = config.GetEnvVars().DBSource

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return db, nil
}
