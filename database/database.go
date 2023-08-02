package database

import (
	"database/sql"
	"github.com/NotKatsu/GoSearch/modules/os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pterm/pterm"
)

type RecommendedAppStruct struct {
	Name string
	Location string
	Visits uint16
}

var database *sql.DB

func SetupDatabase() bool {
	var err error

	location, createAppDataFolderError := os.CreateAppDataFolder()

	if createAppDataFolderError != nil {
		pterm.Fatal.WithFatal(true).Println(err)
		return false
	}

	databaseLocation := location + "/cache.db"

	database, err = sql.Open("sqlite3", databaseLocation)

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

func GetRecommendedApps() ([]RecommendedAppStruct, error) {
	var RecommendedAppStructArray []RecommendedAppStruct
	rows, recommendedAppsDatabaseQueryError := database.Query("SELECT app_name, app_location, app_visits FROM recommended_apps ORDER BY app_visits DESC LIMIT 3")

	if recommendedAppsDatabaseQueryError != nil {
		pterm.Fatal.WithFatal(true).Println(recommendedAppsDatabaseQueryError)
		return RecommendedAppStructArray, recommendedAppsDatabaseQueryError
	}
	defer rows.Close()

	for rows.Next() {
		var currentRecommendedApp RecommendedAppStruct

		rowsScanError := rows.Scan(&currentRecommendedApp.Name, &currentRecommendedApp.Location, &currentRecommendedApp.Visits)

		if rowsScanError != nil {
			pterm.Fatal.WithFatal(true).Println(rowsScanError)
			return RecommendedAppStructArray, rowsScanError
		}

		RecommendedAppStructArray = append(RecommendedAppStructArray, currentRecommendedApp)
	}

	return RecommendedAppStructArray, nil
}