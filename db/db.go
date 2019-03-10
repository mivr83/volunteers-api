package db

import (
	"database/sql"
	"fmt"
	"volunteers-api/settings"
)

var PostgresDb *sql.DB

// connects to database and creates creates globally accessible DB variable
func OpenDbConnection(settings *settings.ConnSettings) error {
	var err error
	PostgresDb, err = sql.Open("postgres", createDbParamString(settings))
	return err
}

func createDbParamString(settings *settings.ConnSettings) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		settings.DbHost,
		settings.DbPort,
		settings.DbUserName,
		settings.DbPassword,
		settings.DbName)
}
