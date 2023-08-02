package database

import (
	"database/sql"

	"github.com/pterm/pterm"
	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB

func SetupDatabase() {
	var err error
	database, err = sql.Open("sqlite3", "database.db")

	if err != nil {
		pterm.Fatal.WithFatal(true).Println(err)
	}

	database.Exec("CREATE TABLE IF NOT EXISTS recommended_apps(app_name VARCHAR(50), app_location VARCHAR(255), app_visits BIGINT)")
}