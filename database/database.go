package database

import (
	"database/sql"

	"github.com/NotKatsu/GoSearch/backend"

	"github.com/NotKatsu/GoSearch/backend/os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pterm/pterm"
)

var default_database *sql.DB
var cache_database *sql.DB

func SetupDatabase() bool {
	var databaseErrorOne error
	var databaseErrorTwo error

	location, createAppDataFolderError := os.CreateAppDataFolder()

	if createAppDataFolderError != nil {
		pterm.Fatal.WithFatal(true).Println(createAppDataFolderError)
		return false
	}

	defaultDatabaseLocation := location + "/default_database.db"
	cacheDatabaseLocation := location + "/cache.db"

	default_database, databaseErrorOne = sql.Open("sqlite3", defaultDatabaseLocation)
	cache_database, databaseErrorTwo = sql.Open("sqlite3", cacheDatabaseLocation)

	if databaseErrorOne != nil && databaseErrorTwo != nil {
		pterm.Fatal.WithFatal(true).Println(databaseErrorOne, databaseErrorTwo)
		return false
	}

	_, databaseTableCreationError1 := default_database.Exec("CREATE TABLE IF NOT EXISTS recommended_apps(app_name VARCHAR(50), app_location VARCHAR(255), app_icon_location VARCHAR(255), app_favorited BOOLEAN, app_visits BIGINT)")
	_, databaseTableCreationError2 := default_database.Exec("CREATE TABLE IF NOT EXISTS settings(id BIGINT, theme VARCHAR(255))")
	_, databaseTableCreationError3 := cache_database.Exec("CREATE TABLE IF NOT EXISTS cache(file_location VARCHAR(255), file_name VARCHAR(255), file_extention)")

	if databaseTableCreationError1 != nil && databaseTableCreationError2 != nil && databaseTableCreationError3 != nil {
		pterm.Fatal.WithFatal(true).Println(databaseTableCreationError1, databaseTableCreationError2, databaseTableCreationError3)
		return false
	} else {
		return true
	}
}

func CacheSize() int64 {
	var CacheRecordCount int64

	err := cache_database.QueryRow("SELECT COUNT(*) FROM cache").Scan(&CacheRecordCount)

	if err != nil {
		pterm.Fatal.WithFatal(true).Println(err)
		return 0
	}
	return CacheRecordCount
}

func SetTheme(theme string) bool {
	var count int

	err := default_database.QueryRow("SELECT COUNT(*) FROM settings WHERE id = 1").Scan(&count)

	if err != nil {
		pterm.Fatal.WithFatal(true).Println(err)
		return false
	}

	if count == 0 {
		_, err = default_database.Exec("INSERT INTO settings (id, theme) VALUES (1, ?)", theme)

		if err != nil {
			pterm.Fatal.WithFatal(true).Println(err)
			return false
		}
		return true
	} else {
		_, err = default_database.Exec("UPDATE settings SET theme = ? WHERE id = 1", theme)

		if err != nil {
			pterm.Fatal.WithFatal(true).Println(err)
			return false
		}
		return true
	}
}

func GetCurrentTheme() string {
	var theme string

	err := default_database.QueryRow("SELECT theme FROM settings WHERE id = 1").Scan(&theme)

	if err != nil {
		if err == sql.ErrNoRows {
			return "Default Theme"
		}
		pterm.Fatal.WithFatal(true).Println(err)
	}
	return theme
}

func GetRecommendedApps() ([]backend.FileReturnStruct, error) {
	var RecommendedAppStructArray []backend.FileReturnStruct
	rows, recommendedAppsDatabaseQueryError := default_database.Query("SELECT app_name, app_location, app_visits, app_favorited FROM recommended_apps ORDER BY CASE WHEN app_favorited = 1 THEN 0 ELSE 1 END, app_visits DESC LIMIT 15")

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

	_, databaseUpdateError := default_database.Exec("UPDATE recommended_apps SET app_favorited = ? WHERE app_name = ? AND app_location = ?", favoriteNumberBool, name, location)

	if databaseUpdateError != nil {
		pterm.Fatal.WithFatal(true).Println(databaseUpdateError)
	}
}

func ClearDatabaseCache() bool {
	_, databaseClearCacheError := default_database.Exec("DELETE FROM recommended_apps; DELETE FROM settings")

	if databaseClearCacheError != nil {
		pterm.Fatal.WithFatal(true).Println(databaseClearCacheError)
		return false
	} else {
		return true
	}
}
