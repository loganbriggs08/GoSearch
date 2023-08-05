package database

import (
	"database/sql"

	"github.com/NotKatsu/GoSearch/backend"

	"github.com/NotKatsu/GoSearch/backend/os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pterm/pterm"
)

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

	_, databaseTableCreationError := database.Exec("CREATE TABLE IF NOT EXISTS recommended_apps(app_name VARCHAR(50), app_location VARCHAR(255), app_icon_location VARCHAR(255), app_favorited BOOLEAN, app_visits BIGINT)")

	if databaseTableCreationError != nil {
		pterm.Fatal.WithFatal(true).Println(err)
		return false
	} else {
		return true
	}
}

func GetRecommendedApps() ([]backend.FileReturnStruct, error) {
	var RecommendedAppStructArray []backend.FileReturnStruct
	rows, recommendedAppsDatabaseQueryError := database.Query("SELECT app_name, app_location, app_visits, app_favorited FROM recommended_apps ORDER BY CASE WHEN app_favorited = 1 THEN 0 ELSE 1 END, app_visits DESC LIMIT 15")

	if recommendedAppsDatabaseQueryError != nil {
		pterm.Fatal.WithFatal(true).Println(recommendedAppsDatabaseQueryError)
		return RecommendedAppStructArray, recommendedAppsDatabaseQueryError
	}
	defer rows.Close()

	for rows.Next() {
		var currentRecommendedApp backend.FileReturnStruct

		rowsScanError := rows.Scan(&currentRecommendedApp.Name, &currentRecommendedApp.Location, &currentRecommendedApp.Visits, &currentRecommendedApp.Favorite)

		if rowsScanError != nil {
			pterm.Fatal.WithFatal(true).Println(rowsScanError)
			return RecommendedAppStructArray, rowsScanError
		}

		RecommendedAppStructArray = append(RecommendedAppStructArray, currentRecommendedApp)
	}

	return RecommendedAppStructArray, nil
}

func UpdateFavorite(name string, location string, favorite bool) {
	var favoriteNumberBool uint8

	if favorite == true {
		favoriteNumberBool = 0
	} else {
		favoriteNumberBool = 1
	}

	_, databaseUpdateError := database.Exec("UPDATE recommended_apps SET app_favorited = ? WHERE app_name = ? AND app_location = ?", favoriteNumberBool, name, location)

	if databaseUpdateError != nil {
		pterm.Fatal.WithFatal(true).Println(databaseUpdateError)
	}
}

func ClearDatabaseCache() bool {
	_, databaseClearCacheError := database.Exec("DELETE FROM recommended_apps")

	if databaseClearCacheError != nil {
		pterm.Fatal.WithFatal(true).Println(databaseClearCacheError)
		return false
	} else {
		return true
	}
}
