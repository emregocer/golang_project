// package database provides the db connection functionality.
package database

import (
	"fmt"

	"github.com/emregocer/golang_project/config"
	"github.com/jmoiron/sqlx"
)

// NewDB connects to a postgresql database with the given config values.
//
// Parameters:
//   - `dbConfig` : config.DatabaseConfig
func NewDB(dbConfig config.DatabaseConfig) (sqlx.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s",
		dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Password, dbConfig.Name)

	db, err := sqlx.Connect("postgres", psqlInfo)

	return *db, err
}
