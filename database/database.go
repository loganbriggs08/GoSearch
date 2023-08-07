package database

import (
	"database/sql"

	"github.com/NotKatsu/GoSearch/backend/appdata"

	"github.com/NotKatsu/GoSearch/backend"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pterm/pterm"
)

var default_database *sql.DB
var cache_database *sql.DB

func SetupDatabase() bool {
	var databaseErrorOne error
	var databaseErrorTwo error

	location, createAppDataFolderError := appdata.CreateAppDataFolder()

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
	_, databaseTableCreationError2 := default_database.Exec("CREATE TABLE IF NOT EXISTS settings(system_cached boolean)")
	_, databaseTableCreationError3 := cache_database.Exec("CREATE TABLE IF NOT EXISTS cache(file_location VARCHAR(255), file_name VARCHAR(255), file_extention)")

	if databaseTableCreationError1 != nil && databaseTableCreationError2 != nil && databaseTableCreationError3 != nil {
		pterm.Fatal.WithFatal(true).Println(databaseTableCreationError1, databaseTableCreationError2, databaseTableCreationError3)
		return false
	} else {
		return true
	}
}

func InsertIntoCache(fileLocation string, fileName string, fileExtention string) bool {
	_, InsertIntoCacheError := cache_database.Exec("INSERT INTO cache(file_location, file_name, file_extention) VALUES (?, ?, ?)", fileLocation, fileName, fileExtention)

	if InsertIntoCacheError != nil {
		pterm.Fatal.WithFatal(true).Println(InsertIntoCacheError)
		return false
	}
	return true
}

func SystemCached() bool {
	var systemCached bool

	err := default_database.QueryRow("SELECT system_cached FROM settings LIMIT 1").Scan(&systemCached)

	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		pterm.Fatal.WithFatal(true).Println(err)
	}
	return systemCached
}

func UpdateCachedSetting(value bool) {
	var systemCached bool

	err := default_database.QueryRow("SELECT system_cached FROM settings LIMIT 1").Scan(&systemCached)

	if err != nil {
		if err == sql.ErrNoRows {
			_, err := default_database.Exec("INSERT INTO settings(system_cached) VALUES(?)", value)

			if err != nil {
				pterm.Fatal.WithFatal(true).Println(err)
			}
		}
		pterm.Fatal.WithFatal(true).Println(err)
	}
}

func ClearDatabaseCache() bool {
	_, databaseClearCacheErrorOne := default_database.Exec("DELETE FROM recommended_apps")
	_, databaseClearCacheErrorTwo := cache_database.Exec("DELETE FROM cache")

	if databaseClearCacheErrorOne != nil && databaseClearCacheErrorTwo != nil {
		pterm.Fatal.WithFatal(true).Println(databaseClearCacheErrorOne, databaseClearCacheErrorTwo)
		return false
	} else {
		return true
	}
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
