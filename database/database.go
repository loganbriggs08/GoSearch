package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pterm/pterm"
)

var database *sql.DB

func SetupDatabase() bool {
	var err error
	database, err = sql.Open("sqlite3", "database.db")

	if err != nil {
		pterm.Fatal.WithFatal(true).Println(err)
		return false
	}

	_, databaseTableCreationError := database.Exec("CREATE TABLE IF NOT EXISTS recommended_apps(app_name VARCHAR(50), app_location VARCHAR(255), app_visits BIGINT)")

	if databaseTableCreationError != nil {
		pterm.Fatal.WithFatal(true).Println(err)
		return false
	} else {
		return true
	}
}